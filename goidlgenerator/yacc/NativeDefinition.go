package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type NativeDefinition struct {
	Identifier string                            `json:"Identifier"`
	Next       interfaces.IDefinitionDeclaration `json:"-"`
}

func (self *NativeDefinition) GetStreamFunctionName() string {
	return self.Identifier
}

func (self *NativeDefinition) GetPackageName() (bool, string, string) {
	return false, "", ""
}

func (self *NativeDefinition) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *NativeDefinition) Kind() interfaces.Kind {
	return interfaces.Native
}

func (self *NativeDefinition) DefaultValue() string {
	return "nil"
}

func (self *NativeDefinition) Predefined() bool {
	return false
}

func (self *NativeDefinition) GetName() string {
	return self.Identifier
}

func (self *NativeDefinition) GetNext() interfaces.IDefinitionDeclaration {
	return self.Next
}

func (self *NativeDefinition) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.Next = typeSpec
}

func (self *NativeDefinition) ClearNext() {
	self.Next = nil
}

func (self *NativeDefinition) GetScopeName() string {
	return self.Identifier
}

func (self *NativeDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type       string `json:"Type"`
		Identifier string `json:"Identifier"`
	}{
		Type:       "native",
		Identifier: self.Identifier,
	})
}

func NewNativeDeclaration(identifier string) *NativeDefinition {
	return &NativeDefinition{
		Identifier: identifier,
		Next:       nil,
	}
}
