package IG

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/bhbosman/Application/Common"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"sync"
	"time"
)

// DemoAPIURL - Demo API URL
const DemoAPIURL = "https://demo-api.ig.com"

// LiveAPIURL - Live API URL - Real trading!
const LiveAPIURL = "https://api.ig.com"

type RestApi struct {
	sync.RWMutex
	ApiUrl                string
	APIKey                string
	AccountID             string
	Identifier            string
	Password              string
	LightStreamerEndpoint string
	OAuthToken            OAuthToken
	httpClient            *http.Client
	timerCb               *TimerCallback
}

func (self *RestApi) Open() error {
	err := self.Login()
	if err != nil {
		self.logError(err, "login failure")
	}
	return err
}

func (self *RestApi) Close() error {
	err := self.Logout()
	if err != nil {
		self.logError(err, "logout failure")
	}
	return err
}

func (self *RestApi) Login() error {
	startTime := time.Now()
	defer func() {
		self.logTime("Login", time.Since(startTime))
	}()

	bodyReq := new(bytes.Buffer)

	var authReq = AuthRequest{
		Identifier: self.Identifier,
		Password:   self.Password,
	}

	if err := json.NewEncoder(bodyReq).Encode(authReq); err != nil {
		return fmt.Errorf("unable to encode JSON response: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", self.ApiUrl, "gateway/deal/session"), bodyReq)
	if err != nil {
		return fmt.Errorf("unable to send HTTP request: %v", err)
	}

	igResponseInterface, _, err := self.doRequest(req, 3, AuthResponse{}, nil)
	if err != nil {
		return err
	}
	authResponse, _ := igResponseInterface.(*AuthResponse)

	if authResponse.OAuthToken.AccessToken == "" {
		return fmt.Errorf("got response but access token is empty")
	}

	expiry, err := strconv.ParseInt(authResponse.OAuthToken.ExpiresIn, 10, 32)
	if err != nil {
		return fmt.Errorf("unable to parse OAuthToken expiry field: %v", err)
	}

	// Refresh token before it will expire
	if expiry <= 15 {
		return fmt.Errorf("token expiry is too short for periodically renewals")
	}
	self.Lock()
	self.OAuthToken = authResponse.OAuthToken
	self.AccountID = authResponse.AccountID
	self.LightStreamerEndpoint = authResponse.LightstreamerEndpoint
	self.timerCb = newTimerCallback(time.Duration(expiry-15)*time.Second, self.refreshCallback, nil)
	self.Unlock()
	return nil
}

func (self *RestApi) Logout() error {
	startTime := time.Now()
	defer func() {
		self.logTime("Logout", time.Since(startTime))
	}()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", self.ApiUrl, "gateway/deal/session"), nil)
	if err != nil {
		return fmt.Errorf("unable to send HTTP request: %v", err)
	}

	_, _, err = self.doRequest(req, 1, nil, nil)
	if err != nil {
		return err
	}

	if self.timerCb != nil {
		_ = self.timerCb.Close()
	}

	self.Lock()
	self.OAuthToken = OAuthToken{}
	self.AccountID = ""
	self.Unlock()

	return nil
}

func (self *RestApi) GetSession(fetchSessionTokens bool) (*SessionResponse, error) {
	startTime := time.Now()
	defer func() {
		self.logTime("GetSession", time.Since(startTime))
	}()

	apiUrlBreakdown, _ := url.Parse(self.ApiUrl)
	urlSession := url.URL{
		Scheme:     apiUrlBreakdown.Scheme,
		Opaque:     "",
		User:       nil,
		Host:       apiUrlBreakdown.Host,
		Path:       "gateway/deal/session",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   fmt.Sprintf("fetchSessionTokens=%v", fetchSessionTokens),
		Fragment:   "",
	}
	resource := urlSession.String()
	req, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to send HTTP request: %v", err)
	}

	igResponseInterface, headerValues, err := self.doRequest(
		req,
		1,
		SessionResponse{},
		[]string{"Cst", "X-Security-Token"})
	if err != nil {
		return nil, err
	}
	sessionResponse := igResponseInterface.(*SessionResponse)
	if s, ok := headerValues["Cst"]; ok {
		sessionResponse.Cst = s
	}
	if s, ok := headerValues["X-Security-Token"]; ok {
		sessionResponse.XSECURITYTOKEN = s
	}

	return sessionResponse, nil
}

func (self *RestApi) refreshCallback(t time.Time, cb interface{}) {
	err := self.RefreshToken()
	if err != nil {
		self.logError(err, "refresh token failure")
	}
}

func (self *RestApi) RefreshToken() error {
	startTime := time.Now()
	defer func() {
		self.logTime("RefreshToken", time.Since(startTime))
	}()

	bodyReq := new(bytes.Buffer)

	var authReq = RefreshTokenRequest{
		RefreshToken: self.OAuthToken.RefreshToken,
	}

	if err := json.NewEncoder(bodyReq).Encode(authReq); err != nil {
		return fmt.Errorf("unable to encode JSON response: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", self.ApiUrl, "gateway/deal/session/refresh-token"), bodyReq)
	if err != nil {
		return fmt.Errorf("unable to send HTTP request: %v", err)
	}

	igResponseInterface, _, err := self.doRequest(req, 1, OAuthToken{}, nil)
	if err != nil {
		return err
	}
	oauthToken, _ := igResponseInterface.(*OAuthToken)

	if oauthToken.AccessToken == "" {
		return fmt.Errorf("got response but access token is empty")
	}

	expiry, err := strconv.ParseInt(oauthToken.ExpiresIn, 10, 32)
	if err != nil {
		return fmt.Errorf("unable to parse OAuthToken expiry field: %v", err)
	}

	// Refresh token before it will expire
	if expiry <= 10 {
		return fmt.Errorf("token expiry is too short for periodically renewals")
	}

	self.Lock()
	self.OAuthToken = *oauthToken
	self.Unlock()

	return nil
}

func (self *RestApi) log(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (self *RestApi) logError(err error, format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	fmt.Printf("Error: %v. Message: %v", err, s)
}

func (self *RestApi) logTime(s string, t time.Duration) {
	self.log("%s: %v\n", s, t)
}

func (self *RestApi) Markets(key string, recursive bool) ([]MarketNavigationMarket, error) {
	startTime := time.Now()
	defer func() {
		self.logTime("Markets", time.Since(startTime))
	}()

	internalMarkets := func(key string) (*MarketNavigation, error) {
		startTime := time.Now()
		defer func() {
			self.logTime("internalMarkets", time.Since(startTime))
		}()

		var urlString string
		if key == "" {
			urlString = "gateway/deal/marketnavigation"
		} else {
			urlString = fmt.Sprintf("gateway/deal/marketnavigation/%v", key)
		}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", self.ApiUrl, urlString), nil)
		if err != nil {
			return nil, fmt.Errorf("unable to send HTTP request: %v", err)
		}
		igResponseInterface, _, err := self.doRequest(req, 1, MarketNavigation{}, nil)
		if err != nil {
			return nil, err
		}
		marketNavigation, ok := igResponseInterface.(*MarketNavigation)
		if ok {
			return marketNavigation, nil
		}
		return nil, fmt.Errorf("could not type to *MarketNavigation")
	}
	marketList := list.New()

	err := Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		type channelData struct {
			wg   *sync.WaitGroup
			data interface{}
		}
		requestChannel := make(chan *channelData, 1024)
		responseChannel := make(chan *channelData, 1024)
		canContinue := true
		for i := 0; i < 1; i++ {
			go func(requestChannel <-chan *channelData, responseChannel chan<- *channelData) {
				for entry := range requestChannel {
					func(entry *channelData) {
						defer entry.wg.Done()
						if canContinue {
							key, ok := entry.data.(string)
							if !ok {
								return
							}
							data, err := internalMarkets(key)
							if err != nil {
								errorList.Add(err)
								canContinue = false
								return
							}
							entry.wg.Add(1)
							responseChannel <- &channelData{
								wg:   entry.wg,
								data: data,
							}
						}
						time.Sleep(time.Second)
					}(entry)

				}
			}(requestChannel, responseChannel)
		}

		add := func(key string, wg *sync.WaitGroup) {
			wg.Add(1)
			go func() {
				requestChannel <- &channelData{
					wg:   wg,
					data: key,
				}
			}()
		}

		go func(responseChannel <-chan *channelData) {
			for entry := range responseChannel {
				func(entry *channelData) {
					defer entry.wg.Done()

					data, ok := entry.data.(*MarketNavigation)
					if !ok {
						return
					}
					if data.Markets != nil && len(data.Markets) > 0 {
						for _, market := range data.Markets {
							marketList.PushBack(market)
						}
					}
					if recursive {
						if data.Nodes != nil && len(data.Nodes) > 0 {
							for _, node := range data.Nodes {
								add(node.Id, entry.wg)
							}
						}
					}
				}(entry)
			}
		}(responseChannel)
		//
		wg := &sync.WaitGroup{}
		add(key, wg)
		wg.Wait()
	})
	if err != nil {
		return nil, err
	}

	result := make([]MarketNavigationMarket, marketList.Len(), marketList.Len())
	i := 0
	for e := marketList.Front(); e != nil; e = e.Next() {
		result[i] = *e.Value.(*MarketNavigationMarket)
		i++
	}
	return result, nil
}

func (self *RestApi) doRequest(req *http.Request, endpointVersion int, igResponse interface{}, requiredHeaderValues []string) (interface{}, map[string]string, error) {
	self.RLock()
	req.Header.Set("X-IG-API-KEY", self.APIKey)
	req.Header.Set("Authorization", "Bearer "+self.OAuthToken.AccessToken)
	req.Header.Set("IG-ACCOUNT-ID", self.AccountID)
	self.RUnlock()

	req.Header.Set("Accept", "application/json; charset=UTF-8")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("VERSION", fmt.Sprintf("%d", endpointVersion))

	resp, err := self.httpClient.Do(req)
	if err != nil {
		return igResponse, nil, fmt.Errorf("unable to get markets data: %v", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			self.logError(err, "response body closure failure")
		}
	}()

	var responseHeader map[string]string = nil
	if requiredHeaderValues != nil{
		responseHeader = make(map[string]string)
		for _, item := range requiredHeaderValues {
			responseHeader[item] = resp.Header.Get(item)
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return igResponse, responseHeader, fmt.Errorf("unable to get body of transactions markets data: %v", err)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil, responseHeader, nil
	}

	if resp.StatusCode != http.StatusOK {
		errCode := &ErrorMessage{}
		if err := json.Unmarshal(body, errCode); err != nil {
			return igResponse, responseHeader, fmt.Errorf("unable to unmarshal JSON response: %v", err)
		}

		errCode.httpCode = resp.StatusCode
		errCode.url = req.URL.String()
		return igResponse, responseHeader, errCode
	}

	objType := reflect.TypeOf(igResponse)
	obj := reflect.New(objType).Interface()
	if obj != nil {
		if err := json.Unmarshal(body, &obj); err != nil {
			return obj, responseHeader, fmt.Errorf("unable to unmarshal JSON response: %v", err)
		}

		return obj, responseHeader, nil
	}
	return igResponse, responseHeader, nil
}

func NewIgRestApi(apiURL, apiKey, accountID, identifier, password string, httpTimeout time.Duration) *RestApi {
	if apiURL != DemoAPIURL && apiURL != LiveAPIURL {
		log.Panic("Invalid endpoint URL", apiURL)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:                  nil,
			DialContext:            nil,
			Dial:                   nil,
			DialTLS:                nil,
			TLSClientConfig:        nil,
			TLSHandshakeTimeout:    0,
			DisableKeepAlives:      false,
			DisableCompression:     false,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    5,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        0,
			ResponseHeaderTimeout:  0,
			ExpectContinueTimeout:  0,
			TLSNextProto:           nil,
			ProxyConnectHeader:     nil,
			MaxResponseHeaderBytes: 0,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       httpTimeout,
	}

	api := &RestApi{
		RWMutex:    sync.RWMutex{},
		ApiUrl:     apiURL,
		APIKey:     apiKey,
		AccountID:  accountID,
		Identifier: identifier,
		Password:   password,
		OAuthToken: OAuthToken{},
		httpClient: httpClient,
	}

	return api
}
