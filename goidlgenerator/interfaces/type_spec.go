package interfaces

import "reflect"

type IDeclarator interface {
	Identifier() string
	Next() IDeclarator
	SetNext(next IDeclarator)
	GetNext() IDeclarator
	ClearNext()
	DefaultValue() IConstantValue
}

type IScopeName interface {
	GetName() string
}

type IDefinedType interface {
	IScopeName
}

type IPreDefinedType interface {
	IDefinedType
}

type IDefinitionDeclaration interface {
	IDefinedType
	GetNext() IDefinitionDeclaration
	SetNext(typeSpec IDefinitionDeclaration)
	ClearNext()
	GetScopeName() string
}

type ITypeDeclaration interface {
	IDefinitionDeclaration
	GetDefinedTyped() IDefinedType
	GetDeclarator() IDeclarator
}

type IConstantValue interface {
	Value() interface{}
	ValueKind() reflect.Kind
	MaxLength() int
}
