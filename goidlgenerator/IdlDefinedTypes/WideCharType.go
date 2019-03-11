package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type WideCharType struct {

}

func (self *WideCharType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *WideCharType) GetPackageName() (bool, string) {
	return false, "Streams"
}

func (self *WideCharType) DefaultValue() string {
	return "'0'"
}

func (self *WideCharType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *WideCharType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *WideCharType) ClearNext() {

}

func (self *WideCharType) GetScopeName() string {
	return self.GetName()
}

func (*WideCharType) GetName() string {
	return "IDLWideCharType"
}
func (self *WideCharType) Kind() interfaces.Kind {
	return interfaces.Invalid
}
func (self *WideCharType) Predefined() bool {
	return true
}

func (self *WideCharType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}

func NewWideCharType() *WideCharType {
	return &WideCharType{}
}
