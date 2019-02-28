package DefinedTypes

import "encoding/json"

type BooleanType struct {
}

func (*BooleanType) GetName() string {
	return "boolean"
}

func (self *BooleanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewBooleanType() *BooleanType {
	return &BooleanType{}
}
