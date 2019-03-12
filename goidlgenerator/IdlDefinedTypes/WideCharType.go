package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type WideCharType struct {
}

func (self *WideCharType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *WideCharType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *WideCharType) DefaultValue() string {
	return "'0'"
}

func (*WideCharType) GetName() string {
	return "IDLWideCharType"
}
func (self *WideCharType) Kind() interfaces.Kind {
	return interfaces.WChar
}
func (self *WideCharType) Predefined() bool {
	return true
}

func (self *WideCharType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

