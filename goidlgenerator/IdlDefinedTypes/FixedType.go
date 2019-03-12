package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type FixedType struct {
}

func (self *FixedType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *FixedType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *FixedType) DefaultValue() string {
	return "0.0"
}

func (self *FixedType) Predefined() bool {
	return true
}

func (self *FixedType) Kind() interfaces.Kind {
	return interfaces.Fixed
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
