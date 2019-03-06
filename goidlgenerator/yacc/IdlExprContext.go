package yacc

//publishGo:generate goyacc -o idl.publishGo -p "IdlExpr" idl.y

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type IdlExprContext struct {
	declaredTypes []interfaces.IDefinitionDeclaration
	specification []interfaces.IDefinitionDeclaration
}

func NewIdlExprContext() *IdlExprContext {
	return &IdlExprContext{
		declaredTypes: make([]interfaces.IDefinitionDeclaration, 0, 16),
	}
}

func (self *IdlExprContext) AddDefinition(definition interfaces.IDefinitionDeclaration) {
	self.declaredTypes = append(self.declaredTypes, definition)
}

func (self *IdlExprContext) FindScopeName(identifier string) interfaces.IDefinitionDeclaration {
	for _, decl := range self.declaredTypes {
		if decl.GetScopeName() == identifier {
			return decl
		}
	}
	return nil
}

func (self *IdlExprContext) GetSpecification() []interfaces.IDefinitionDeclaration {
	return self.specification
}
