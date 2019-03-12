package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type OctetType struct {
}

func (self *OctetType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *OctetType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *OctetType) DefaultValue() string {
	return "0"
}

func (self *OctetType) Kind() interfaces.Kind {
	return interfaces.Octet
}
func (self *OctetType) Predefined() bool {
	return true
}

func (*OctetType) GetName() string {
	return "IDLOctetType"
}

func (self *OctetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

