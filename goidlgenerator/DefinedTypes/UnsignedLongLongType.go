package DefinedTypes

import "encoding/json"

type UnsignedLongLongType struct {
}

func (*UnsignedLongLongType) GetName() string {
	return "IDLUnsignedLongLongType"
}

func (self *UnsignedLongLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewUnsignedLongLongType() *UnsignedLongLongType {
	return &UnsignedLongLongType{}
}
