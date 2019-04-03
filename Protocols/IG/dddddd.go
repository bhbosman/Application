package IG

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)




type ISubscription interface {
	GetItems() []string
	GetFields() []string
	GetMode() string
	GetDataAdapter() string
	GetLocalRef() string
	OnNewData(data string)
	GetId() string
}
type QuoteAdapterSubscription struct {
	subscriptionId string
	Items          []string
	Fields         []string
	Mode           string
	DataAdapter    string
	LocalRef       string
}

func (qas *QuoteAdapterSubscription) GetId() string {
	return qas.subscriptionId
}

func (qas *QuoteAdapterSubscription) OnNewData(data string) {

}

func (qas *QuoteAdapterSubscription) GetItems() []string {
	return qas.Items
}

func (qas *QuoteAdapterSubscription) GetFields() []string {
	return qas.Fields
}

func (qas *QuoteAdapterSubscription) GetMode() string {
	return qas.Mode
}

func (qas *QuoteAdapterSubscription) GetDataAdapter() string {
	return qas.DataAdapter
}

func (qas *QuoteAdapterSubscription) GetLocalRef() string {
	return qas.LocalRef
}




type ConnectAnswer struct {
	Answer    string
	Success   bool
	ErrCode   int64
	ErrReason string
	Status    int
}







type LightStreamConnection struct {
	client           http.Client
	InitialApiUrl    string
	connectionStream io.Closer
	log              *log.Logger
	dataDict         map[string]string
	connectionCount  int
	subscriptions    map[string]ISubscription
	SubscriptionId   int
	connectionOpen   bool
}

func CreateLightStreamConnection(logOut io.Writer, apiUrl string) (*LightStreamConnection, error) {
	return &LightStreamConnection{
		InitialApiUrl:    apiUrl,
		client:           http.Client{},
		connectionStream: nil,
		log:              log.New(logOut, "LightStreamer: ", log.LstdFlags),
		dataDict:         make(map[string]string),
		connectionCount:  0,
		subscriptions:    make(map[string]ISubscription),
	}, nil
}

func (connection *LightStreamConnection) IsConnected() bool {
	return connection.connectionOpen
}

func (connection *LightStreamConnection) readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		connection.log.Println("Error reading line. Error", err.Error())
		return "", err
	}
	return string(bytes.Runes(bytes.TrimSpace(line))), nil
}

func (connection *LightStreamConnection) Subscribe(subscription ISubscription) (ans ConnectAnswer, subscriptionId string, err error) {
	data := url.Values{}
	connection.SubscriptionId += 1
	subscriptionId = strconv.Itoa(connection.SubscriptionId)
	data.Set("LS_table", subscriptionId)
	data.Set("LS_data_adapter", subscription.GetDataAdapter())
	data.Set("LS_mode", subscription.GetMode())
	data.Set("LS_schema", strings.Join(subscription.GetFields(), " "))
	data.Set("LS_id", strings.Join(subscription.GetItems(), " "))

	response, err := connection.Control("add", data)
	if err != nil {
		return ConnectAnswer{}, "", err
	}
	defer func() {
		response.Body.Close()
	}()
	streamReader := bufio.NewReader(response.Body)
	answer, _ := connection.readLine(streamReader)

	success := answer == "OK"
	if success {
		connection.subscriptions[subscriptionId] = subscription
		return ConnectAnswer{
			Answer:    answer,
			Success:   true,
			ErrCode:   0,
			ErrReason: "",
			Status:    response.StatusCode,
		},
			subscriptionId,
			nil
	}
	errorCodeStr, _ := connection.readLine(streamReader)
	errorCode, _ := strconv.ParseInt(errorCodeStr, 10, 32)
	errorMessage, _ := connection.readLine(streamReader)

	return ConnectAnswer{
		Answer:    answer,
		Success:   false,
		ErrCode:   errorCode,
		ErrReason: errorMessage,
		Status:    response.StatusCode,
	},
		"",
		nil
}

func (connection *LightStreamConnection) Unsubscribe(subscription ISubscription) (ans ConnectAnswer, err error) {
	data := url.Values{}

	data.Add("LS_table", subscription.GetId())

	response, err := connection.Control("delete", data)
	if err != nil {
		return ConnectAnswer{}, err
	}
	defer func() {
		response.Body.Close()
	}()
	streamReader := bufio.NewReader(response.Body)
	answer, _ := connection.readLine(streamReader)

	success := answer == "OK"
	if success {
		delete(connection.subscriptions, subscription.GetId())
		return ConnectAnswer{
			Answer:    answer,
			Success:   true,
			ErrCode:   0,
			ErrReason: "",
			Status:    response.StatusCode,
		},
			nil
	}
	errorCodeStr, _ := connection.readLine(streamReader)
	errorCode, _ := strconv.ParseInt(errorCodeStr, 10, 32)
	errorMessage, _ := connection.readLine(streamReader)

	return ConnectAnswer{
		Answer:    answer,
		Success:   false,
		ErrCode:   errorCode,
		ErrReason: errorMessage,
		Status:    response.StatusCode,
	},
		nil
}

func (connection *LightStreamConnection) findUrl() string {
	result := connection.InitialApiUrl

	redirectUrl, ok := connection.dataDict["ControlAddress"]
	if ok {
		result = "https://" + redirectUrl
	}

	return result

}

func (connection *LightStreamConnection) Connect(dataFunc func(data url.Values)) (*ConnectAnswer, error) {

	connection.connectionCount += 1
	var resource string
	_, ok := connection.dataDict["SessionId"]
	if !ok {
		resource = "/lightstreamer/create_session.txt"
	} else {
		resource = "/lightstreamer/bind_session.txt"
	}

	urlToUse := connection.findUrl()

	u, err := url.ParseRequestURI(urlToUse)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := u.String()
	response, err := connection.Send(
		urlStr,
		dataFunc,
		func(header http.Header) {
			header.Set("Content-Type", "application/x-www-form-urlencoded")
			header.Set("Cache-Control", "no-cache")
		})
	if err != nil {
		connection.log.Println("Error connecting to api. Error:", err.Error())
		return nil, err
	}
	streamReader := bufio.NewReader(response.Body)
	ans, err := connection.readLine(streamReader)

	if connection.connectionStream != nil {
		if err := connection.connectionStream.Close(); err != nil {
			connection.log.Println("Error in closing connection")
		}
	}

	connection.connectionStream = response.Body
	success := ans == "OK"
	if success {
		var wg sync.WaitGroup
		wg.Add(1)
		go func(closer io.Closer, reader *bufio.Reader) {
			defer func() {
				closer.Close()
				connection.connectionOpen = false
			}()
			connection.connectionOpen = true

			// create a new dict
			connection.dataDict = make(map[string]string)

			// read settings
			for {
				ans, err := connection.readLine(streamReader)
				if err != nil {
					connection.log.Println("Error closing from stream. Connection #", connection.connectionCount, "Error: ", err.Error())
					break
				}

				if ans == "" {
					break
				}
				data := strings.Split(ans, ":")
				if len(data) == 2 {
					if data[0] == "Preamble" {
						continue
					}
					connection.log.Println(ans)
					if _, ok := connection.dataDict[data[0]]; !ok {
						connection.dataDict[data[0]] = data[1]
					}
				}
			}
			wg.Done()

			rebind := false

			// read data
			for {
				ans, err := connection.readLine(streamReader)
				if err != nil {
					connection.log.Println("Error closing from stream. Connection #", connection.connectionCount, "Error: ", err.Error())
					break
				}
				if strings.Index(ans, "PROBE") == 0 {
					connection.log.Println("Heartbeat...")
					continue
				} else if strings.Index(ans, "ERROR") == 0 {
					connection.log.Println("Error on queue. Message from queue:", ans)
					break
				} else if strings.Index(ans, "END ") == 0 {
					connection.log.Println("Connection closed. Message from queue:", ans)
					break
				} else if strings.Index(ans, "LOOP") == 0 {
					connection.log.Println("Connection closed. Message from queue:", ans)
					rebind = true
					break
				} else if strings.Index(ans, "SYNC ERROR") == 0 {
					connection.log.Println("Connection closed. Message from queue:", ans)
					break
				}
				data := strings.SplitN(ans, ",", 2)
				if len(data) == 2 {
					if subscription, ok := connection.subscriptions[data[0]]; ok {
						subscription.OnNewData(data[1])
					}
				}
				connection.log.Println(ans)
			}
			if rebind {
				connection.log.Println("Connection lost. Connection can rebind. Calling connect")
				connection.Connect(
					func(data url.Values) {
						data.Set("LS_session", connection.dataDict["SessionId"])
					})

			} else {
				connection.log.Println("Connection lost. Resetting data. Currently no auto connect")
				connection.Reset()

			}
		}(response.Body, streamReader)
		wg.Wait()

		return &ConnectAnswer{
			Answer:    ans,
			Success:   true,
			ErrCode:   0,
			ErrReason: "",
			Status:    response.StatusCode,
		}, nil

	} else {
		errorCodeStr, _ := connection.readLine(streamReader)
		errorCode, _ := strconv.ParseInt(errorCodeStr, 10, 32)
		errorMessage, _ := connection.readLine(streamReader)
		response.Body.Close()
		return &ConnectAnswer{
			Answer:    ans,
			Success:   false,
			ErrCode:   errorCode,
			ErrReason: errorMessage,
			Status:    response.StatusCode,
		}, nil
	}
}

func (connection *LightStreamConnection) Disconnect() (*ConnectAnswer, error) {
	response, err := connection.Control("destroy", url.Values{})
	if err != nil {
		return nil, err
	}

	defer func() {
		response.Body.Close()
	}()

	streamReader := bufio.NewReader(response.Body)
	ans, err := connection.readLine(streamReader)
	if ans == "OK" {
		return &ConnectAnswer{
			Answer:    ans,
			Success:   true,
			ErrCode:   0,
			ErrReason: "",
			Status:    response.StatusCode,
		}, nil
	}

	closeCloser := func(closer io.Closer, what string) {
		connection.log.Println("Closing", what, "...")
		if closer != nil {
			if err := closer.Close(); err != nil {
				connection.log.Println("Error while closing", what, "Error: %s", err.Error())
			}
		}
	}
	closeCloser(connection.connectionStream, "connectionStream")

	connection.Reset()

	return &ConnectAnswer{
		Answer:    ans,
		Success:   false,
		ErrCode:   0,
		ErrReason: "",
		Status:    response.StatusCode,
	}, nil
}

func (connection *LightStreamConnection) Send(
	urlStr string,
	dataFunc func(data url.Values),
	addToHeaderFunc func(header http.Header)) (*http.Response, error) {
	data := url.Values{}
	if dataFunc != nil {
		dataFunc(data)
	}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))

	if addToHeaderFunc != nil {
		addToHeaderFunc(r.Header)
	}
	resp, err := connection.client.Do(r)
	if err != nil {
		connection.log.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	return resp, nil
}

func (connection *LightStreamConnection) Control(operation string, otherData url.Values) (*http.Response, error) {
	if sessionId, ok := connection.dataDict["SessionId"]; ok {
		resource := "/lightstreamer/control.txt"
		u, _ := url.ParseRequestURI(connection.findUrl())
		u.Path = resource
		urlStr := u.String()
		response, err := connection.Send(
			urlStr,
			func(data url.Values) {
				data.Set("LS_session", sessionId)
				data.Set("LS_op", operation)
				for k, _ := range otherData {
					data.Set(k, otherData.Get(k))
				}
			},
			func(header http.Header) {
				header.Add("Content-Type", "application/x-www-form-urlencoded")
			})
		if err != nil {
			connection.log.Println("Error connecting to api. Error:", err.Error())
		}
		return response, err
	}
	return nil, errors.New("no sessionid")
}
func (connection *LightStreamConnection) Reset() {
	connection.dataDict = make(map[string]string)
}

