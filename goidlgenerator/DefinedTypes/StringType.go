package DefinedTypes

import (
	"encoding/json"
	"fmt"
)

type StringType struct {
	length int64
}

func (*StringType) GetName() string {
	return "IDLStringType"
}

func (receiver *StringType) String() string {
	return fmt.Sprintf("StringType")
}

func (self *StringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewStringType(length int64) *StringType {
	return &StringType{
		length: length,
	}
}
