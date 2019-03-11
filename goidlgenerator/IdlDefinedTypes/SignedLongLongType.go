package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type SignedLongLongType struct {

}

func (self *SignedLongLongType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *SignedLongLongType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *SignedLongLongType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *SignedLongLongType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *SignedLongLongType) DefaultValue() string {
	return "0"
}

func (self *SignedLongLongType) ClearNext() {

}

func (self *SignedLongLongType) GetScopeName() string {
	return self.GetName()
}

func (*SignedLongLongType) GetName() string {
	return "IDLSignedLongLongType"
}

func (self *SignedLongLongType) Kind() interfaces.Kind {
	return interfaces.Int64
}
func (self *SignedLongLongType) Predefined() bool {
	return true
}

func (self *SignedLongLongType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewSignedLongLongType() *SignedLongLongType {
	return &SignedLongLongType{}
}
