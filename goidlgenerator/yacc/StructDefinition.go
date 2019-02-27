package yacc

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/gofintech/goidlgenerator/interfaces"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type StructDefinition struct {
	Type       string                            `json:"Type"`
	Identifier string                            `json:"Identifier"`
	Members    []*Member                         `json:"Members"`
	Next       interfaces.IDefinitionDeclaration `json:"-"`
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
	return fmt.Sprintf("struct %v", self.Identifier)
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
