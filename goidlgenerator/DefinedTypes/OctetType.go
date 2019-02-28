package DefinedTypes

import "encoding/json"

type OctetType struct {
}

func (*OctetType) GetName() string {
	return "octet"
}

func (self *OctetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewOctetType() *OctetType {
	return &OctetType{}
}
