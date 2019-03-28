package IG

import "fmt"

type AuthRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// AuthResponse - IG auth response
type AuthResponse struct {
	AccountID             string     `json:"accountId"`
	ClientID              string     `json:"clientId"`
	LightstreamerEndpoint string     `json:"lightstreamerEndpoint"`
	OAuthToken            OAuthToken `json:"oauthToken"`
	TimezoneOffset        int        `json:"timezoneOffset"` // In seconds
}

type SessionResponse struct {
	AccountID             string `json:"accountId"`
	ClientID              string `json:"clientId"`
	Currency              string `json:"currency"`
	LightstreamerEndpoint string `json:"lightstreamerEndpoint"`
	TimezoneOffset        int    `json:"timezoneOffset"` // In seconds
	Locale                string `json:"locale"`
	Cst                   string `json:"CST"`
	XSECURITYTOKEN        string `json:"X-SECURITY-TOKEN"`
}

type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type MarketNavigationNode struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type MarketNavigationMarket struct {
	Bid                      float64 `json:"bid"`                      //	Bid price
	DelayTime                float64 `json:"delayTime"`                //	Price delay time in minutes
	Epic                     string  `json:"epic"`                     //	Instrument epic identifier
	Expiry                   string  `json:"expiry"`                   //	Instrument expiry period
	High                     float64 `json:"high"`                     //	Highest price of the day
	InstrumentName           string  `json:"instrumentName"`           //	Instrument name
	InstrumentType           string  `json:"instrumentType"`           //
	LotSize                  float64 `json:"lotSize"`                  //	Instrument lot size
	Low                      float64 `json:"low"`                      //	Lowest price of the day
	MarketStatus             string  `json:"marketStatus"`             // 	ddd
	NetChange                float64 `json:"netChange"`                //	Price net change
	Offer                    float64 `json:"offer"`                    //	Offer price
	OtcTradeable             bool    `json:"otcTradeable"`             //	True if OTC tradeable
	PercentageChange         float64 `json:"percentageChange"`         //	Percentage price change on the day
	ScalingFactor            float64 `json:"scalingFactor"`            //	multiplying factor to determine actual pip value for the levels used by the instrument
	StreamingPricesAvailable bool    `json:"streamingPricesAvailable"` //	True if streaming prices are available, i.e. the market is tradeable and the client holds the necessary access permissions
	UpdateTime               string  `json:"updateTime"`               //	Local time of last price update (milliseconds since epoch)
	UpdateTimeUTC            string  `json:"updateTimeUtc"`            //	Time of last price update

}

type MarketNavigation struct {
	Markets []*MarketNavigationMarket `json:"markets"`
	Nodes   []*MarketNavigationNode   `json:"nodes"`
}

type ErrorMessage struct {
	ErrorCode string `json:"errorCode"`
	httpCode  int
	url       string
}

func (self *ErrorMessage) Error() string {
	return fmt.Sprintf("Error: %v. Code %v. Busy with %v", self.ErrorCode, self.httpCode, self.url)
}
