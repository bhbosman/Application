package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type DoubleType struct {
}

func (self *DoubleType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *DoubleType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *DoubleType) DefaultValue() string {
	return "0.0"
}

func (*DoubleType) GetName() string {
	return "IDLDoubleType"
}
func (self *DoubleType) Predefined() bool {
	return true
}

func (self *DoubleType) Kind() interfaces.Kind {
	return interfaces.Double
}

func (self *DoubleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}
