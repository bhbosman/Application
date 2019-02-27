package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import "github.com/bhbosman/gofintech/goidlgenerator/interfaces"

func Newconst_dcl() interfaces.IDefinitionDeclaration {
	return nil
}

func Newunion_dcl() interfaces.IDefinitionDeclaration {
	return nil
}

func Newenum_dcl() interfaces.IDefinitionDeclaration {
	return nil
}

func Newstruct_forward_dcl() interfaces.IDefinitionDeclaration {
	return nil
}
