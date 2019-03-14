package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type SignedLongType struct {
}

func (self *SignedLongType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *SignedLongType) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *SignedLongType) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *SignedLongType) DefaultValue() string {
	return "0"
}

func (self *SignedLongType) Kind() interfaces.Kind {
	return interfaces.Int32
}

func (self *SignedLongType) Predefined() bool {
	return true
}

func (*SignedLongType) GetName() string {
	return "IDLSignedLongType"
}

func (self *SignedLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}
