package DefinedTypes

import "encoding/json"

type SignedShortType struct {
}

func (*SignedShortType) GetName() string {
	return "IDLSignedShortType"
}

func (self *SignedShortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewSignedShortType() *SignedShortType {
	return &SignedShortType{}
}
