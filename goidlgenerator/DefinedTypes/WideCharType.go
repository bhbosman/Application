package DefinedTypes

import "encoding/json"

type WideCharType struct {
}

func (*WideCharType) GetName() string {
	return "IDLWideCharType"
}

func (self *WideCharType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewWideCharType() *WideCharType {
	return &WideCharType{}
}
