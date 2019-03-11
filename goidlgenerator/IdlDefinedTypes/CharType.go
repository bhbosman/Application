package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type CharType struct {

}

func (self *CharType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *CharType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *CharType) DefaultValue() string {
	return "'0'"
}

func (self *CharType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *CharType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *CharType) ClearNext() {

}

func (self *CharType) GetScopeName() string {
	return self.GetName()
}

func (self *CharType) Kind() interfaces.Kind {
	return interfaces.Invalid
}
func (self *CharType) Predefined() bool {
	return true
}

func (*CharType) GetName() string {
	return "IDLCharType"
}

func (self *CharType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewCharType() *CharType {
	return &CharType{}
}
