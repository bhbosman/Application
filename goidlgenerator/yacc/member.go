package yacc

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type Member struct {
	DefinedType interfaces.IDefinedType
	Declarator  interfaces.IDeclarator
	Next        *Member
}

func (self *Member) MarshalJSON() ([]byte, error) {
	self.DefinedType.GetName()
	return json.Marshal(&struct {
		Type        string                  `json:"Type"`
		Declarator  interfaces.IDeclarator  `json:"Declarator"`
		DefinedType interfaces.IDefinedType `json:"DefinedType"`
	}{
		Type:        "Member",
		Declarator:  self.Declarator,
		DefinedType: self.DefinedType,
	})
}

func NewMember(type_spec interfaces.IDefinedType, declarator interfaces.IDeclarator, next *Member) *Member {
	return &Member{
		DefinedType: type_spec,
		Declarator:  declarator,
		Next:        next,
	}
}
