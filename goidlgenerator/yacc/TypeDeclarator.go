package yacc

//publishGo:generate goyacc -o idl.publishGo -p "IdlExpr" idl.y

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type TypeDeclarator struct {
	DefinedTyped interfaces.IDefinedType
	Declarator   interfaces.IDeclarator
	next         interfaces.IDefinitionDeclaration
}

func (self *TypeDeclarator) GetScopeName() string {
	return self.Declarator.Identifier()
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

func (self *TypeDeclarator) GetDefinedTyped() interfaces.IDefinedType {
	return self.DefinedTyped
}

func (self *TypeDeclarator) GetDeclarator() interfaces.IDeclarator {
	return self.Declarator
}

func (self *TypeDeclarator) GetName() string {
	return fmt.Sprintf("typedef %v %v: ", self.DefinedTyped.GetName(), self.Declarator.Identifier())
}

func Newtypedef_dcl(definedTyped interfaces.IDefinedType, declarator interfaces.IDeclarator) *TypeDeclarator {
	if definedTyped == nil {
		return nil
	}
	if declarator == nil {
		return nil
	}
	return &TypeDeclarator{
		DefinedTyped: definedTyped,
		Declarator:   declarator,
	}
}
