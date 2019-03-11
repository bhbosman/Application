package IdlDefinedTypes

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type StringType struct {
	length      int64

}

func (self *StringType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *StringType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *StringType) DefaultValue() string {
	return string("\"\"")
}

func (self *StringType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *StringType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *StringType) ClearNext() {

}

func (self *StringType) GetScopeName() string {
	return self.GetName()
}

func (*StringType) GetName() string {
	return "IDLStringType"
}

func (self *StringType) Kind() interfaces.Kind {
	return interfaces.String
}
func (self *StringType) Predefined() bool {
	return true
}

func (receiver *StringType) String() string {
	return fmt.Sprintf("StringType")
}

func (self *StringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewStringType(length int64) *StringType {
	return &StringType{
		length: length,
	}
}
