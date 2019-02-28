package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type TypeDeclarator struct {
	definedTyped interfaces.IDefinedType
	declarator   interfaces.IDeclarator
	next         interfaces.IDefinitionDeclaration
}

func (self *TypeDeclarator) GetScopeName() string {
	return self.declarator.Identifier()
}

func (self *TypeDeclarator) GetNext() interfaces.IDefinitionDeclaration {
	return self.next
}

func (self *TypeDeclarator) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.next = typeSpec
}

func (self *TypeDeclarator) ClearNext() {
	self.next = nil
}

func (self *TypeDeclarator) DefinedTyped() interfaces.IDefinedType {
	return self.definedTyped
}

func (self *TypeDeclarator) Declarator() interfaces.IDeclarator {
	return self.declarator
}

func (self *TypeDeclarator) GetName() string {
	return fmt.Sprintf("typedef %v %v: ", self.definedTyped.GetName(), self.declarator.Identifier())
}

func Newtypedef_dcl(definedTyped interfaces.IDefinedType, declarator interfaces.IDeclarator) *TypeDeclarator {
	if definedTyped == nil {
		return nil
	}
	if declarator == nil {
		return nil
	}
	return &TypeDeclarator{
		definedTyped: definedTyped,
		declarator:   declarator,
	}
}
