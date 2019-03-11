package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchTypeInformation struct {
}


func (self *MitchTypeInformation) CanScope(decl interfaces.IDefinedType) bool {
	if decl.Kind() == interfaces.Enum {
		return true
	}

	if decl.Kind() == interfaces.TypeDeclarator {
		if typeDecl, ok := decl.(interfaces.ITypeDeclaration); ok {
			if typeDecl.GetDefinedTyped().Kind() == interfaces.BitField {
				return true
			}
		}
	}

	return false
}

func (self *MitchTypeInformation) DefaultDecls() ([]interfaces.IDefinitionDeclaration, error) {
	return []interfaces.IDefinitionDeclaration{
		&MitchAlpha{

		},
		&MitchBitField{

		},
		&MitchByte{

		},
		&MitchDate{

		},
		&MitchTime{

		},
		&MitchPrice04{
		},
		&MitchPrice08{

		},
		&MitchUInt08{

		},
		&MitchUInt16{

		},
		&MitchUInt32{

		},
		&MitchUInt64{

		},
	}, nil
}

func (self *MitchTypeInformation) Name() interfaces.IDlBaseType {
	return interfaces.IDlBaseType_Mitch
}
