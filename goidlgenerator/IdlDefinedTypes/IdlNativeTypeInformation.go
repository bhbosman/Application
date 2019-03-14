package IdlDefinedTypes

import (
	"errors"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type idlNativeTypeInformation struct {
	BooleanType          *BooleanType
	CharType             *CharType
	DoubleType           *DoubleType
	FixedType            *FixedType
	FloatType            *FloatType
	LongDoubleType       *LongDoubleType
	OctetType            *OctetType
	SignedLongLongType   *SignedLongLongType
	SignedLongType       *SignedLongType
	SignedShortType      *SignedShortType
	StringType           *StringType
	UnsignedLongLongType *UnsignedLongLongType
	UnsignedLongType     *UnsignedLongType
	UnSignedShortType    *UnSignedShortType
	WideCharType         *WideCharType
	WideStringType       *WideStringType
	createFunc           map[interfaces.Kind]func(data interface{}) interfaces.IDefinedType
}

func (self *idlNativeTypeInformation) CanScope(decl interfaces.IDefinedType) bool {
	return true
}

func (self *idlNativeTypeInformation) CreateType(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
	if result, ok := self.createFunc[kind]; ok {
		return result(data), nil
	}
	return nil, errors.New(fmt.Sprintf("type (%v) not available in %v type information", kind.String(), self.Name()))

}

func (self *idlNativeTypeInformation) DefaultDecls() ([]interfaces.IDefinitionDeclaration, error) {

	return []interfaces.IDefinitionDeclaration{
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.BooleanType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.CharType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.DoubleType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.FixedType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.FloatType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.LongDoubleType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.OctetType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.SignedLongLongType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.SignedLongType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.SignedShortType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.StringType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.UnsignedLongLongType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.UnsignedLongType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.UnSignedShortType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.WideCharType),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.WideStringType),
	}, nil
}

func (self *idlNativeTypeInformation) Name() interfaces.BaseTypeDescription {
	return interfaces.IDlBaseType_Native
}

func NewIdlNativeTypeInformation() *idlNativeTypeInformation {
	result := &idlNativeTypeInformation{
		BooleanType:        &BooleanType{},
		CharType:           &CharType{},
		DoubleType:         &DoubleType{},
		FixedType:          &FixedType{},
		FloatType:          &FloatType{},
		LongDoubleType:     &LongDoubleType{},
		OctetType:          &OctetType{},
		SignedLongLongType: &SignedLongLongType{},
		SignedLongType:     &SignedLongType{},
		SignedShortType:    &SignedShortType{},
		StringType: &StringType{
			length: 0,
		},
		UnsignedLongLongType: &UnsignedLongLongType{},
		UnsignedLongType:     &UnsignedLongType{},
		UnSignedShortType:    &UnSignedShortType{},
		WideCharType:         &WideCharType{},
		WideStringType: &WideStringType{
			length: 0,
		},
		createFunc: make(map[interfaces.Kind]func(data interface{}) interfaces.IDefinedType),
	}

	result.createFunc[interfaces.Bool] = func(data interface{}) interfaces.IDefinedType { return result.BooleanType }
	result.createFunc[interfaces.Char] = func(data interface{}) interfaces.IDefinedType { return result.CharType }
	result.createFunc[interfaces.Double] = func(data interface{}) interfaces.IDefinedType { return result.DoubleType }
	result.createFunc[interfaces.Fixed] = func(data interface{}) interfaces.IDefinedType { return result.FixedType }
	result.createFunc[interfaces.Float] = func(data interface{}) interfaces.IDefinedType { return result.FloatType }
	result.createFunc[interfaces.LongDouble] = func(data interface{}) interfaces.IDefinedType { return result.LongDoubleType }
	result.createFunc[interfaces.Octet] = func(data interface{}) interfaces.IDefinedType { return result.OctetType }
	result.createFunc[interfaces.Int64] = func(data interface{}) interfaces.IDefinedType { return result.SignedLongLongType }
	result.createFunc[interfaces.Int32] = func(data interface{}) interfaces.IDefinedType { return result.SignedLongType }
	result.createFunc[interfaces.Int16] = func(data interface{}) interfaces.IDefinedType { return result.SignedShortType }
	result.createFunc[interfaces.String] = func(data interface{}) interfaces.IDefinedType { return result.StringType }
	result.createFunc[interfaces.Uint64] = func(data interface{}) interfaces.IDefinedType { return result.UnsignedLongLongType }
	result.createFunc[interfaces.Uint32] = func(data interface{}) interfaces.IDefinedType { return result.UnsignedLongType }
	result.createFunc[interfaces.Uint16] = func(data interface{}) interfaces.IDefinedType { return result.UnSignedShortType }
	result.createFunc[interfaces.WChar] = func(data interface{}) interfaces.IDefinedType { return result.WideCharType }
	result.createFunc[interfaces.WideString] = func(data interface{}) interfaces.IDefinedType { return result.WideStringType }

	return result
}
