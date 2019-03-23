package IG

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

type Configuration struct {
	FileName string
	ApiUrl     string `json:"ApiUrl"`
	ApiKey     string `json:"ApiKey"`
	Identifier string `json:"Identifier"`
	Password   string `json:"Password"`
	Markets		[]MarketNavigationMarket `json:"Markets"`
}



func (configuration *Configuration) Save() error{

	f, e := os.Create(configuration.FileName)
	if e != nil {
		return e
	}
	defer func() {
		f.Close()
	}()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	return encoder.Encode(configuration)
}

func NewConfiguration() (*Configuration, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	fileName := usr.HomeDir + "/.IG/IGConnect.json"
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	result := &Configuration{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}
	result.FileName = fileName
	return result, nil
}

