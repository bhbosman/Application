package DefinedTypes

import "encoding/json"

type FloatType struct {
}

func (*FloatType) GetName() string {
	return ""
}

func (self *FloatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewFloatType() *FloatType {
	return &FloatType{}
}
