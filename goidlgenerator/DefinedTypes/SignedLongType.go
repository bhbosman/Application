package DefinedTypes

import (
	"encoding/json"
)

type SignedLongType struct {
}

func (*SignedLongType) GetName() string {
	return "IDLSignedLongType"
}

func (self *SignedLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewSignedLongType() *SignedLongType {
	return &SignedLongType{}
}
