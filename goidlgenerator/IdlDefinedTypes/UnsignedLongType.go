package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type UnsignedLongType struct {

}

func (self *UnsignedLongType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *UnsignedLongType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *UnsignedLongType) DefaultValue() string {
	return "0"
}

func (self *UnsignedLongType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *UnsignedLongType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *UnsignedLongType) ClearNext() {

}

func (self *UnsignedLongType) GetScopeName() string {
	return self.GetName()
}

func (*UnsignedLongType) GetName() string {
	return "IDLUnsignedLongType"
}

func (self *UnsignedLongType) Kind() interfaces.Kind {
	return interfaces.Uint32
}
func (self *UnsignedLongType) Predefined() bool {
	return true
}

func (self *UnsignedLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewUnsignedLongType() *UnsignedLongType {
	return &UnsignedLongType{}
}
