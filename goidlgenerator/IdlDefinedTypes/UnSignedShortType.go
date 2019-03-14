package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type UnSignedShortType struct {
}

func (self *UnSignedShortType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *UnSignedShortType) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *UnSignedShortType) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *UnSignedShortType) DefaultValue() string {
	return "0"
}

func (*UnSignedShortType) GetName() string {
	return "IDLUnSignedShortType"
}

func (self *UnSignedShortType) Kind() interfaces.Kind {
	return interfaces.Uint16
}
func (self *UnSignedShortType) Predefined() bool {
	return true
}

func (self *UnSignedShortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}
