package DefinedTypes

import "encoding/json"

type SignedLongLongType struct {
}

func (*SignedLongLongType) GetName() string {
	return "IDLSignedLongLongType"
}

func (self *SignedLongLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewSignedLongLongType() *SignedLongLongType {
	return &SignedLongLongType{}
}
