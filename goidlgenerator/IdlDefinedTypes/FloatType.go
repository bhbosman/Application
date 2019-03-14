package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type FloatType struct {
}

func (self *FloatType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *FloatType) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *FloatType) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *FloatType) DefaultValue() string {
	return "0.0"
}

func (self *FloatType) Kind() interfaces.Kind {
	return interfaces.Float
}
func (self *FloatType) Predefined() bool {
	return true
}

func (*FloatType) GetName() string {
	return "IDLFloatType"
}

func (self *FloatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}
