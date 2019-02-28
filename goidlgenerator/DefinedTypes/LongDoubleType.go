package DefinedTypes

import "encoding/json"

type LongDoubleType struct {
}

func (*LongDoubleType) GetName() string {
	return "LongDouble"
}

func (self *LongDoubleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewLongDoubleType() *LongDoubleType {
	return &LongDoubleType{}
}
