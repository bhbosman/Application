package interfaces

type IDeclarator interface {
	Identifier() string
	Next() IDeclarator
	SetNext(next IDeclarator)
	GetNext() IDeclarator
	ClearNext()
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
	DefinedTyped() IDefinedType
	Declarator() IDeclarator
}
