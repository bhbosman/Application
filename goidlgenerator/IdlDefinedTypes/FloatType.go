package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type FloatType struct {

}

func (self *FloatType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *FloatType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *FloatType) DefaultValue() string {
	return "0.0"
}

func (self *FloatType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *FloatType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *FloatType) ClearNext() {

}

func (self *FloatType) GetScopeName() string {
	return self.GetName()
}

func (self *FloatType) Kind() interfaces.Kind {
	return interfaces.Invalid
}
func (self *FloatType) Predefined() bool {
	return true
}

func (*FloatType) GetName() string {
	return "IDLFloatType"
}

func (self *FloatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewFloatType() *FloatType {
	return &FloatType{}
}
