package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchTypeInformation struct {
	mitchAlpha    *mitchAlpha
	mitchBitField *mitchBitField
	mitchByte     *mitchByte
	mitchDate     *mitchDate
	mitchTime     *mitchTime
	mitchPrice04  *mitchPrice04
	mitchPrice08  *mitchPrice08
	mitchUInt08   *mitchUInt08
	mitchUInt16   *mitchUInt16
	mitchUInt32   *mitchUInt32
	mitchUInt64   *mitchUInt64
	createFunc    map[interfaces.Kind]func(data interface{}) interfaces.IDefinedType
}


func (self *mitchTypeInformation) CreateType(kind interfaces.Kind, data interface{}) interfaces.IDefinedType {
	if result, ok := self.createFunc[kind]; ok {
		return result(data)
	}
	return nil
}

func (self *mitchTypeInformation) CanScope(decl interfaces.IDefinedType) bool {
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

func (self *mitchTypeInformation) DefaultDecls() ([]interfaces.IDefinitionDeclaration, error) {
	return []interfaces.IDefinitionDeclaration{
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchAlpha),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchBitField),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchByte),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchDate),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchTime),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchPrice04),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchPrice08),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchUInt08),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchUInt16),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchUInt32),
		interfaces.NewWrapDefinedTypeToIDefinitionDeclaration(self.mitchUInt64),
	}, nil
}

func (self *mitchTypeInformation) Name() interfaces.BaseTypeDescription {
	return interfaces.IDlBaseType_Mitch
}



func NewMitchTypeInformation() *mitchTypeInformation {

	result := &mitchTypeInformation{
		mitchAlpha: &mitchAlpha{
			length: 0,
		},
		mitchBitField: &mitchBitField{},
		mitchByte:     &mitchByte{},
		mitchDate:     &mitchDate{},
		mitchTime:     &mitchTime{},
		mitchPrice04:  &mitchPrice04{},
		mitchPrice08:  &mitchPrice08{},
		mitchUInt08:   &mitchUInt08{},
		mitchUInt16:   &mitchUInt16{},
		mitchUInt32:   &mitchUInt32{},
		mitchUInt64:   &mitchUInt64{},
		createFunc:    make(map[interfaces.Kind]func(data interface{}) interfaces.IDefinedType),
	}
	result.createFunc[interfaces.MitchAlpha] = func(data interface{}) interfaces.IDefinedType { return result.mitchAlpha }
	result.createFunc[interfaces.MitchBitField] = func(data interface{}) interfaces.IDefinedType { return result.mitchBitField }
	result.createFunc[interfaces.MitchByte] = func(data interface{}) interfaces.IDefinedType { return result.mitchByte }
	result.createFunc[interfaces.MitchDate] = func(data interface{}) interfaces.IDefinedType { return result.mitchDate }
	result.createFunc[interfaces.MitchTime] = func(data interface{}) interfaces.IDefinedType { return result.mitchTime }
	result.createFunc[interfaces.MitchPrice04] = func(data interface{}) interfaces.IDefinedType { return result.mitchPrice04 }
	result.createFunc[interfaces.MitchPrice08] = func(data interface{}) interfaces.IDefinedType { return result.mitchPrice08 }
	result.createFunc[interfaces.MitchUInt08] = func(data interface{}) interfaces.IDefinedType { return result.mitchUInt08 }
	result.createFunc[interfaces.MitchUInt16] = func(data interface{}) interfaces.IDefinedType {return result.mitchUInt16}
	result.createFunc[interfaces.MitchUInt32] = func(data interface{}) interfaces.IDefinedType { return result.mitchUInt32 }
	result.createFunc[interfaces.MitchUInt64] = func(data interface{}) interfaces.IDefinedType { return result.mitchUInt64 }

	return result
}
