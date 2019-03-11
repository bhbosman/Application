package interfaces

type IDlBaseType string

const IDlBaseType_Native_Value IDlBaseType = "IdlNative"
const IDlBaseType_Mitch_Value IDlBaseType = "Mitch"

const (
	IDlBaseType_Native IDlBaseType = IDlBaseType_Native_Value
	IDlBaseType_Mitch              = IDlBaseType_Mitch_Value
)

type IBaseTypeInformation interface {
	Name() IDlBaseType
	DefaultDecls() ([]IDefinitionDeclaration, error)
	CanScope(decl IDefinedType) bool
	//GetPackageName() (bool, string)
}

type IDeclarator interface {
	Identifier() string
	Next() IDeclarator
	SetNext(next IDeclarator)
	GetNext() IDeclarator
	ClearNext()
	DefaultValue() IConstantValue
}

type IDefinedType interface {
	GetPackageName() (bool, string)
	GetSequenceCount() (bool, int)
	GetName() string
	Kind() Kind
	DefaultValue() string
	Predefined() bool
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
