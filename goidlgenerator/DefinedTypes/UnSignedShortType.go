package DefinedTypes

import "encoding/json"

type UnSignedShortType struct {
}

func (*UnSignedShortType) GetName() string {
	return "IDLUnSignedShortType"
}

func (self *UnSignedShortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewUnSignedShortType() *UnSignedShortType {
	return &UnSignedShortType{}
}
