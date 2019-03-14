package interfaces

type IDeclarator interface {
	Identifier() string
	Next() IDeclarator
	SetNext(next IDeclarator)
	GetNext() IDeclarator
	ClearNext()
	DefaultValue() IConstantValue
}

type IDefinedType interface {
	GetPackageName() (bool, string, string)
	GetSequenceCount() (bool, int)
	GetName() string
	Kind() Kind
	DefaultValue() string
	Predefined() bool
	GetStreamFunctionName() string
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
	ValueKind() Kind
	MaxLength() int
}
