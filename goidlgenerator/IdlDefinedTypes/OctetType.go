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

func (self *OctetType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *OctetType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *OctetType) ClearNext() {

}

func (self *OctetType) GetScopeName() string {
	return self.GetName()
}

func (self *OctetType) Kind() interfaces.Kind {
	return interfaces.Invalid
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

func NewOctetType() *OctetType {
	return &OctetType{}
}
