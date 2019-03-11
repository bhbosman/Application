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

func (self *SignedLongType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *SignedLongType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
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

func (self *SignedLongType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	// do Nothing
}

func (self *SignedLongType) ClearNext() {
	// do Nothing
}

func (self *SignedLongType) GetScopeName() string {
	return self.GetName()
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

func NewSignedLongType() *SignedLongType {
	return &SignedLongType{}
}
