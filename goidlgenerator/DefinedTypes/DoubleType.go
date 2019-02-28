package DefinedTypes

import "encoding/json"

type DoubleType struct {
}

func (*DoubleType) GetName() string {
	return "double"
}

func (self *DoubleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewDoubleType() *DoubleType {
	return &DoubleType{}
}
