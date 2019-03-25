package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/Common"
	"github.com/bhbosman/Application/goidlgenerator/Extensions"
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
				defaultValue.GetExportValue(),
				defaultValue.GetExportValue()))
		}
	}
	_, _ = writer.WriteString(fmt.Sprintf(")\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishEnum) Export(
	writer io.StringWriter,
	TypeInformation interfaces.IBaseTypeInformation) {
	self.ExportDefinition(writer)

	typeNamePrefix := self.data.Identifier
	typeCode := CalculateCrc(typeNamePrefix)
	_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))

	self.ExportDefaultConstructor(writer, TypeInformation)

	returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
	_ = GenerateMessageWriteFunction(
		writer,
		returnType,
		GenerateMessageWriteFunctionParams{
			TypeInformation: TypeInformation,
			kind:            interfaces.Enum,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
		})
	_ = GenerateMessageReadFunction(
		writer,
		returnType,
		GenerateMessageReadFunctionParams{
			TypeInformation: TypeInformation,
			kind:            interfaces.Enum,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
			defaultValue:    "0",
		})

	_ = self.GenerateWriteFunction(writer, TypeInformation, typeNamePrefix, typeCode)
	_ = self.GenerateReadFunction(writer, TypeInformation, typeNamePrefix, typeCode)
}

func (self *publishEnum) ExportDefaultConstructor(writer io.StringWriter, TypeInformation interfaces.IBaseTypeInformation) {
	_, _ = writer.WriteString(fmt.Sprintf("func New%v()%v {\n", self.data.Identifier, self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn %v\n", self.data.DefaultValue()))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishEnum) GenerateReadFunction(
	writer io.StringWriter,
	TypeInformation interfaces.IBaseTypeInformation,
	typeNamePrefix string, typeCode uint32) error {

	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
		_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Read_%v(stream Streams.I%vReader) (value %v, byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
		_, _ = writer.WriteString(fmt.Sprintf("\tv, b, e :=  stream.Read_byte()\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn %v(v), b, e\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self publishEnum) GenerateWriteFunction(
	writer io.StringWriter,
	TypeInformation interfaces.IBaseTypeInformation,
	typeNamePrefix string, typeCode uint32) error {

	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
		_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (int, error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn stream.Write_byte(byte(value))\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func NewPublishEnum(data *yacc.EnumDecl, typeInformation interfaces.IBaseTypeInformation) *publishEnum {
	return &publishEnum{
		data:            data,
		typeInformation: typeInformation,
	}
}
