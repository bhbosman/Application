package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type SignedShortType struct {
}

func (self *SignedShortType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *SignedShortType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *SignedShortType) DefaultValue() string {
	return "0"
}

func (*SignedShortType) GetName() string {
	return "IDLSignedShortType"
}
func (self *SignedShortType) Predefined() bool {
	return true
}

func (self *SignedShortType) Kind() interfaces.Kind {
	return interfaces.Int16
}

func (self *SignedShortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

