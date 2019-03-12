package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Extansions"
	. "github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishStruct struct {
	data            *yacc.StructDefinition
	typeInformation IBaseTypeInformation
}

func (self *publishStruct) Export(writer io.StringWriter, TypeInformation IBaseTypeInformation) {
	typeNamePrefix := self.data.Identifier
	typeCode := CalculateCrc(typeNamePrefix)
	_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))

	self.ExportDefinition(writer, TypeInformation)
	self.ExportDefaultConstructor(writer, TypeInformation)
	self.GenerateWriteFunction(writer, TypeInformation, typeNamePrefix, typeCode)
	self.GenerateReadFunction(writer, TypeInformation, typeNamePrefix, typeCode)

	GenerateMessageWriteFunction(
		writer,
		self.data,
		GenerateMessageWriteFunctionParams{
			TypeInformation: TypeInformation,
			kind:            Struct,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
		})
	GenerateMessageReadFunction(
		writer,
		self.data,
		GenerateMessageReadFunctionParams{
			TypeInformation: TypeInformation,
			kind:            Struct,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
			defaultValue:    "nil",
		})
	GenerateTypeCodeFunction(writer, Struct, typeNamePrefix, typeCode)
	GenerateIsTypeCodeFunction(writer, typeNamePrefix, typeCode)

}

func (self *publishStruct) ExportDefinition(writer io.StringWriter, TypeInformation IBaseTypeInformation) {
	_, _ = writer.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
	for _, member := range self.data.Members {
		defaultValue := member.Declarator.DefaultValue()
		returnType := Extansions.TypeValueForDefinedType(member.DefinedType)
		if defaultValue == nil {
			_, _ = writer.WriteString(fmt.Sprintf(
				"\t%v %v\n",
				member.Declarator.Identifier(),
				returnType))
		} else {

			ss := defaultValue.ValueKind().String()
			_, _ = writer.WriteString(fmt.Sprintf(
				"\t%v %v // default value: %v(%v)\n ",
				member.Declarator.Identifier(),
				returnType,
				ss,
				GetExportValue(defaultValue)))
		}
	}
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self *publishStruct) ExportDefaultConstructor(writer io.StringWriter, TypeInformation IBaseTypeInformation) {
	_, _ = writer.WriteString(fmt.Sprintf("func New%v()*%v {\n", self.data.Identifier, self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn &%v{}\n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self *publishStruct) GenerateReadFunction(writer io.StringWriter, TypeInformation IBaseTypeInformation, typeNamePrefix string, typeCode uint32) {
	returnType := "*" + typeNamePrefix
	_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func Read_%v(stream Streams.I%vReader) (value %v, byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\tvalue = New%v()\n", typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))
	for _, item := range self.data.Members {
		switch item.DefinedType.Kind() {
		case Enum:
			_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v, n, err = Read_%v(stream)\n", item.Declarator.Identifier(), item.DefinedType.GetName()))
			break
		default:
			_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v, n, err = stream.Read_%v()\n", item.Declarator.Identifier(), item.DefinedType.Kind().String()))
		}
		_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t\treturn nil, 0, err\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
	}
	_, _ = writer.WriteString(fmt.Sprintf("\treturn value, byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self *publishStruct) GenerateWriteFunction(writer io.StringWriter, TypeInformation IBaseTypeInformation, typeNamePrefix string, typeCode uint32) {
	returnType := "*" + typeNamePrefix
	_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))
	for _, item := range self.data.Members {
		switch item.DefinedType.Kind() {
		case Enum:
			_, _ = writer.WriteString(fmt.Sprintf("\tn, err = Write_%v(stream, value.%v)\n", item.DefinedType.GetName(), item.Declarator.Identifier()))
			break
		default:
			_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_%v(value.%v)\n", item.DefinedType.Kind().String(), item.Declarator.Identifier()))
		}
		_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
	}
	_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func NewPublishStruct(data *yacc.StructDefinition, typeInformation IBaseTypeInformation) *publishStruct {
	return &publishStruct{
		data:            data,
		typeInformation: typeInformation,
	}
}
