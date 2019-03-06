package yacc

//publishGo:generate goyacc -o idl.publishGo -p "IdlExpr" idl.y

import (
	"errors"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

const NoCtx int = 99998
const NoLex int = 99999

const DefNotFound int = 10005
const ErrorOnAddTypedefDcl = 10006
const ErrorTypeisNull = 10007
const ErrorMustbeAnInt = 10008


func GetIdlExprLex(item IdlExprLexer) (*IdlExprLex, error) {
	if item == nil {
		return nil, errors.New("could not find *IdlExprLex")
	}
	if result, ok := item.(*IdlExprLex); ok {
		return result, nil
	}
	return nil, errors.New("could not find *IdlExprLex")
}

func GetIdlExprContext(item IdlExprLexer) (*IdlExprContext, error) {
	lex, err := GetIdlExprLex(item)
	if err != nil {
		return nil, errors.New("could not find *IdlExprLex")
	}
	return lex.idlExprContext, nil
}

func AddDefinitions(definitions interfaces.IDefinitionDeclaration) []interfaces.IDefinitionDeclaration {
	result := make([]interfaces.IDefinitionDeclaration, 0, 16)
	definition01 := definitions
	for definition01 != nil {
		inner01 := definition01
		if typeDeclaration, ok := definition01.(interfaces.ITypeDeclaration); ok {
			declarator02 := typeDeclaration.GetDeclarator()
			for declarator02 != nil {
				inner02 := declarator02
				newType := Newtypedef_dcl(typeDeclaration.GetDefinedTyped(), inner02)
				result = append(result, newType)

				declarator02 = declarator02.GetNext()
				inner02.ClearNext()
			}
		} else {
			result = append(result, inner01)
		}
		definition01 = definition01.GetNext()
		inner01.ClearNext()
	}
	return result
}

func AddTypedefDcl(idlExprlex IdlExprLexer, typeDecl interfaces.ITypeDeclaration) error {
	context, err := GetIdlExprContext(idlExprlex)
	if err != nil {
		return err
	}

	declarator := typeDecl.GetDeclarator()
	for declarator != nil {
		newType := Newtypedef_dcl(typeDecl.GetDefinedTyped(), declarator)
		context.AddDefinition(newType)
		declarator = declarator.GetNext()
	}

	return nil
}

func AddTypeDclToContext(idlExprlex IdlExprLexer, typeDecl interfaces.IDefinitionDeclaration) error {
	context, err := GetIdlExprContext(idlExprlex)
	if err != nil {
		return err
	}
	context.AddDefinition(typeDecl)
	return nil
}

func SendError(item IdlExprLexer, msg string) {
	lex, err := GetIdlExprLex(item)
	if err != nil {
		return
	}
	item.Error(fmt.Sprintf("%v at %v, %v", msg, lex.Row, lex.Col))
}

func GetLast(item interfaces.IDefinitionDeclaration) interfaces.IDefinitionDeclaration {
	result := item
	for result.GetNext() != nil {
		result = result.GetNext()
	}
	return result
}
