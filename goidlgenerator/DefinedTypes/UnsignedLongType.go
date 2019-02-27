package DefinedTypes

import "encoding/json"

type UnsignedLongType struct {
}

func (*UnsignedLongType) GetName() string {
	return "UnsignedLongType"
}

func (self *UnsignedLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewUnsignedLongType() *UnsignedLongType {
	return &UnsignedLongType{}
}
