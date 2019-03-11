package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type BooleanType struct {

}

func (self *BooleanType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *BooleanType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *BooleanType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *BooleanType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *BooleanType) ClearNext() {

}

func (self *BooleanType) GetScopeName() string {
	return self.GetName()
}

func (self *BooleanType) DefaultValue() string {
	return "false"
}

func (self *BooleanType) Kind() interfaces.Kind {
	return interfaces.Bool
}

func (self *BooleanType) Predefined() bool {
	return true
}

func (*BooleanType) GetName() string {
	return "IDLBooleanType"
}

func (self *BooleanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewBooleanType() *BooleanType {
	return &BooleanType{}
}
