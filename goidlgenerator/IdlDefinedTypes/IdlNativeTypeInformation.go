package IdlDefinedTypes

import "github.com/bhbosman/Application/goidlgenerator/interfaces"

type IdlNativeTypeInformation struct {
}



func (self *IdlNativeTypeInformation) CanScope(decl interfaces.IDefinedType) bool {
	return true
}

func (self *IdlNativeTypeInformation) DefaultDecls() ([]interfaces.IDefinitionDeclaration, error) {

	return []interfaces.IDefinitionDeclaration{
		&BooleanType{

		},
		&CharType{

		},
		&DoubleType{

		},
		&FixedType{

		},
		&FloatType{

		},
		&LongDoubleType{

		},
		&OctetType{

		},
		&SignedLongLongType{

		},
		&SignedLongType{

		},
		&SignedShortType{

		},
		&StringType{

		},
		&UnsignedLongLongType{

		},
		&UnsignedLongType{

		},
		&UnSignedShortType{

		},
		&WideCharType{

		},
		&WideStringType{

		},
	}, nil
}

func (self *IdlNativeTypeInformation) Name() interfaces.IDlBaseType {
	return interfaces.IDlBaseType_Native
}
