package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchTypeInformation struct {
	mitchByte    *mitchByte
	mitchDate    *mitchDate
	mitchTime    *mitchTime
	mitchPrice04 *mitchPrice04
	mitchPrice08 *mitchPrice08
	mitchUInt08  *mitchUInt08
	mitchUInt16  *mitchUInt16
	mitchUInt32  *mitchUInt32
	mitchUInt64  *mitchUInt64
	createFunc   map[interfaces.Kind]func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error)
}

func (self *mitchTypeInformation) CreateType(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
	if result, ok := self.createFunc[kind]; ok {
		return result(kind, data)
	}
	return nil, fmt.Errorf("type (%v) not available in %v type information", kind.String(), self.Name())
}

func (self *mitchTypeInformation) CanScope(decl interfaces.IDefinedType) bool {
	return true
}

func (self *mitchTypeInformation) DefaultDecls() ([]interfaces.IDefinitionDeclaration, error) {
	return []interfaces.IDefinitionDeclaration{}, nil
}

func (self *mitchTypeInformation) Name() interfaces.BaseTypeDescription {
	return interfaces.IDlBaseType_Mitch
}

func NewMitchTypeInformation() *mitchTypeInformation {

	result := &mitchTypeInformation{
		mitchByte:    &mitchByte{},
		mitchDate:    &mitchDate{},
		mitchTime:    &mitchTime{},
		mitchPrice04: &mitchPrice04{},
		mitchPrice08: &mitchPrice08{},
		mitchUInt08:  &mitchUInt08{},
		mitchUInt16:  &mitchUInt16{},
		mitchUInt32:  &mitchUInt32{},
		mitchUInt64:  &mitchUInt64{},
		createFunc:   make(map[interfaces.Kind]func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error)),
	}
	result.createFunc[interfaces.MitchAlpha] =
		func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
			if i, ok := data.(int64); ok {
				return &mitchAlpha{
					length: i,
				}, nil

			}
			return nil, fmt.Errorf("type (%v) not available in %v type information", kind.String(), result.Name())
		}

	result.createFunc[interfaces.MitchMessageNumber] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		if i, ok := data.(int64); ok {
			if 0 <= i && i <= 255 {
				return newMitchMessageNumber(byte(i)), nil
			}
		}
		return nil, fmt.Errorf("type (%v) not available in %v type information", kind.String(), result.Name())
	}

	result.createFunc[interfaces.MitchMessageLength] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		if i, ok := data.(int64); ok {
			if 0 <= i && i <= 65535 {
				return newMitchMessageLength(uint16(i)), nil
			}
		}
		return nil, fmt.Errorf("type (%v) not available in %v type information", kind.String(), result.Name())
	}

	result.createFunc[interfaces.MitchBitField] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		if arrayOfData, ok := data.([]string); ok {
			if len(arrayOfData) == 8 {
			}
			return newMitchBitField(
				arrayOfData[0],
				arrayOfData[1],
				arrayOfData[2],
				arrayOfData[3],
				arrayOfData[4],
				arrayOfData[5],
				arrayOfData[6],
				arrayOfData[7]), nil
		}
		return nil, fmt.Errorf("type (%v) not available in %v type information", kind.String(), result.Name())
	}
	result.createFunc[interfaces.MitchByte] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchByte, nil
	}
	result.createFunc[interfaces.MitchDate] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchDate, nil
	}
	result.createFunc[interfaces.MitchTime] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchTime, nil
	}
	result.createFunc[interfaces.MitchPrice04] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchPrice04, nil
	}
	result.createFunc[interfaces.MitchPrice08] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchPrice08, nil
	}
	result.createFunc[interfaces.MitchUInt08] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchUInt08, nil
	}
	result.createFunc[interfaces.MitchUInt16] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchUInt16, nil
	}
	result.createFunc[interfaces.MitchUInt32] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchUInt32, nil
	}
	result.createFunc[interfaces.MitchUInt64] = func(kind interfaces.Kind, data interface{}) (interfaces.IDefinedType, error) {
		return result.mitchUInt64, nil
	}

	return result
}
