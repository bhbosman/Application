package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import "github.com/bhbosman/Application/goidlgenerator/interfaces"

type EnumDecl struct {
	Type       string                            `json:"Type"`
	Identifier string                            `json:"Identifier"`
	Next       interfaces.IDefinitionDeclaration `json:"-"`
	Decls      []interfaces.IDeclarator
}

func (self *EnumDecl) GetName() string {
	return self.Identifier
}

func (self *EnumDecl) GetNext() interfaces.IDefinitionDeclaration {
	return self.Next
}

func (self *EnumDecl) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.Next = typeSpec
}

func (self *EnumDecl) ClearNext() {
	self.Next = nil
}

func (self *EnumDecl) GetScopeName() string {
	return self.Identifier
}

func (self *EnumDecl) AddMember(decl interfaces.IDeclarator) {
	self.Decls = append(self.Decls, decl)
}

func NewEnumDcl(identifier string) *EnumDecl {
	return &EnumDecl{
		Type:       "EnumDecl",
		Identifier: identifier,
		Next:       nil,
		Decls:      make([]interfaces.IDeclarator, 0, 8),
	}
}
