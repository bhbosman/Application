package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type UnsignedLongLongType struct {
}

func (self *UnsignedLongLongType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *UnsignedLongLongType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *UnsignedLongLongType) DefaultValue() string {
	return "0"
}

func (*UnsignedLongLongType) GetName() string {
	return "IDLUnsignedLongLongType"
}

func (self *UnsignedLongLongType) Kind() interfaces.Kind {
	return interfaces.Uint64
}

func (self *UnsignedLongLongType) Predefined() bool {
	return true
}

func (self *UnsignedLongLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

