package yacc

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type StructDefinition struct {
	Type       string                            `json:"Type"`
	Identifier string                            `json:"Identifier"`
	Members    []*Member                         `json:"Members"`
	Next       interfaces.IDefinitionDeclaration `json:"-"`
}

func (self *StructDefinition) GetStreamFunctionName() string {
	return self.Identifier
}

func (self *StructDefinition) GetPackageName() (bool, string, string) {
	return true, "", self.Identifier
}

func (self *StructDefinition) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *StructDefinition) Kind() interfaces.Kind {
	return interfaces.Struct
}

func (self *StructDefinition) DefaultValue() string {
	return "nil"
}

func (self *StructDefinition) Predefined() bool {
	return false
}

func (self *StructDefinition) GetScopeName() string {
	return self.Identifier
}

func (self *StructDefinition) ClearNext() {
	self.Next = nil
}

func (self *StructDefinition) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.Next = typeSpec
}

func (self *StructDefinition) GetName() string {
	return self.Identifier

}

func (self *StructDefinition) AddMember(typeSpec interfaces.IDefinedType, declarator interfaces.IDeclarator) {
	self.Members = append(self.Members, NewMember(typeSpec, declarator, nil))
}

func (self *StructDefinition) String() string {
	bytes, err := json.MarshalIndent(self, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (self *StructDefinition) GetNext() interfaces.IDefinitionDeclaration {
	return self.Next
}

func NewStructDefinition(identifier string) *StructDefinition {
	return &StructDefinition{
		Type:       "struct",
		Identifier: identifier,
		Members:    make([]*Member, 0),
		Next:       nil,
	}
}
