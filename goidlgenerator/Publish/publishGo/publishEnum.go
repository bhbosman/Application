package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishEnum struct {
	data            *yacc.EnumDecl
	typeInformation interfaces.IBaseTypeInformation
}

func (self publishEnum) ExportDefinition(writer io.StringWriter) {
	_, _ = writer.WriteString(fmt.Sprintf("type %s byte \n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("//noinspection ALL\n"))
	_, _ = writer.WriteString(fmt.Sprintf("const (\n"))
	for index, decl := range self.data.Decls {
		defaultValue := decl.DefaultValue()

		if defaultValue == nil {
			if index == 0 {
				_, _ = writer.WriteString(fmt.Sprintf("\t%v_%v %v = iota\n", self.data.Identifier, decl.Identifier(), self.data.Identifier))
			} else {
				_, _ = writer.WriteString(fmt.Sprintf("\t%v_%v\n", self.data.Identifier, decl.Identifier()))
			}
		} else {

			_, _ = writer.WriteString(fmt.Sprintf(
				"\t%v_%v  = %v // default value: byte(%v)\n ",
				self.data.Identifier,
				decl.Identifier(),
				GetExportValue(defaultValue),
				GetExportValue(defaultValue)))
		}
	}
	_, _ = writer.WriteString(fmt.Sprintf(")\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishEnum) Export(writer io.StringWriter, TypeInformation interfaces.IBaseTypeInformation) {
	self.ExportDefinition(writer)

	typeNamePrefix := self.data.Identifier
	typeCode := CalculateCrc(typeNamePrefix)
	_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))

	GenerateMessageWriteFunction(
		writer,
		self.data,
		GenerateMessageWriteFunctionParams{
			TypeInformation: TypeInformation,
			kind:            interfaces.Enum,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
		})
	GenerateMessageReadFunction(
		writer,
		self.data,
		GenerateMessageReadFunctionParams{
			TypeInformation: TypeInformation,
			kind:            interfaces.Enum,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
			defaultValue:    "0",
		})
}

func NewPublishEnum(data *yacc.EnumDecl, typeInformation interfaces.IBaseTypeInformation) *publishEnum {
	return &publishEnum{
		data:            data,
		typeInformation: typeInformation,
	}
}
