package DefinedTypes

import "encoding/json"

type FixedType struct {
}

func (*FixedType) GetName() string {
	return "IDLFixedType"
}

func (self *FixedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewFixedType() *FixedType {
	return &FixedType{}
}
