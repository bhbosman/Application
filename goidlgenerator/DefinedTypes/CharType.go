package DefinedTypes

import "encoding/json"

type CharType struct {
}

func (*CharType) GetName() string {
	return "char"
}

func (self *CharType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewCharType() *CharType {
	return &CharType{}
}
