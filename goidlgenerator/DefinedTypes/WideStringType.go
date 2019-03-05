package DefinedTypes

import "encoding/json"

type WideStringType struct {
	length int64
}

func (*WideStringType) GetName() string {
	return "IDLWideStringType"
}

func (self *WideStringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewWideStringType(length int64) *WideStringType {
	return &WideStringType{
		length: length,
	}
}
