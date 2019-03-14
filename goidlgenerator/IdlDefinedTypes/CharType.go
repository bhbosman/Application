package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type CharType struct {
}

func (self *CharType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *CharType) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *CharType) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *CharType) DefaultValue() string {
	return "'0'"
}

func (self *CharType) Kind() interfaces.Kind {
	return interfaces.Char
}
func (self *CharType) Predefined() bool {
	return true
}

func (*CharType) GetName() string {
	return "IDLCharType"
}

func (self *CharType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}
