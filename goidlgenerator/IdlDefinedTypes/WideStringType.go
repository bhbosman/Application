package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type WideStringType struct {
	length int64
}

func (self *WideStringType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *WideStringType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *WideStringType) DefaultValue() string {
	return "\"\""
}

func (*WideStringType) GetName() string {
	return "IDLWideStringType"
}
func (self *WideStringType) Kind() interfaces.Kind {
	return interfaces.String
}
func (self *WideStringType) Predefined() bool {
	return true
}

func (self *WideStringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

