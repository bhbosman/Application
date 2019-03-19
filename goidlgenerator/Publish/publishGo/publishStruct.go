package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/Common"
	"github.com/bhbosman/Application/goidlgenerator/Extensions"
	. "github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
	"strconv"
)

type publishStruct struct {
	data            *yacc.StructDefinition
	typeInformation IBaseTypeInformation
}

type LineWriter struct {
	errorList Common.IErrorList
	writer    io.StringWriter
}

func (receiver LineWriter) Write(format string, a ...interface{})  {
	_, e := receiver.writer.WriteString(fmt.Sprintf(format, a))
	receiver.errorList.Add(e)
}

func NewLineWriter(writer io.StringWriter, errorList Common.IErrorList) LineWriter {
	return LineWriter{
		errorList: errorList,
		writer:    writer,
	}
}

func (self *publishStruct) Export(writer io.StringWriter, TypeInformation IBaseTypeInformation) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		lineWriter := NewLineWriter(writer, errorList)
		typeNamePrefix := self.data.Identifier
		typeCode := CalculateCrc(typeNamePrefix)
		lineWriter.Write("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode)
		errorList.Add(self.ExportDefinition(writer, TypeInformation))
		errorList.Add(self.ExportDefaultConstructor(writer, TypeInformation))
		errorList.Add(self.GenerateWriteFunction(writer, TypeInformation, typeNamePrefix, typeCode))
		errorList.Add(self.GenerateReadFunction(writer, TypeInformation, typeNamePrefix, typeCode))

		returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
		errorList.Add(GenerateMessageWriteFunction(
			writer,
			returnType,
			GenerateMessageWriteFunctionParams{
				TypeInformation: TypeInformation,
				kind:            Struct,
				typeNamePrefix:  typeNamePrefix,
				typeCode:        typeCode,
			}))
		errorList.Add(GenerateMessageReadFunction(
			writer,
			returnType,
			GenerateMessageReadFunctionParams{
				TypeInformation: TypeInformation,
				kind:            Struct,
				typeNamePrefix:  typeNamePrefix,
				typeCode:        typeCode,
				defaultValue:    "nil",
			}))
		errorList.Add(GenerateTypeCodeFunction(writer, Struct, typeNamePrefix, typeCode))
		errorList.Add(GenerateIsTypeCodeFunction(writer, typeNamePrefix, typeCode))
	})
}

func (self *publishStruct) ExportDefinition(writer io.StringWriter, TypeInformation IBaseTypeInformation) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
		for _, member := range self.data.Members {
			defaultValue := member.Declarator.DefaultValue()
			returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(member.DefinedType)
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
					defaultValue.GetExportValue()))
			}
		}
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self *publishStruct) ExportDefaultConstructor(writer io.StringWriter, TypeInformation IBaseTypeInformation) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		lineWriter := NewLineWriter(writer, errorList)
		lineWriter.Write("func New%v()*%v {\n", self.data.Identifier, self.data.Identifier)
		lineWriter.Write("\treturn &%v{}\n", self.data.Identifier)
		lineWriter.Write("}\n")
		lineWriter.Write("\n")
	})
}

func (self *publishStruct) GenerateReadFunction(writer io.StringWriter, TypeInformation IBaseTypeInformation, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		returnType := "*" + typeNamePrefix
		_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Read_%v(stream Streams.I%vReader) (value %v, byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
		_, _ = writer.WriteString(fmt.Sprintf("\tvalue = New%v()\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))
		for index, item := range self.data.Members {
			errorCheckFunc := func() {
				_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\t\treturn nil, 0, err\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
			}
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t// Index: %v, Member Name: %v, Type: %v \n", index, item.Declarator.Identifier(), item.DefinedType.GetName()))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			switch item.DefinedType.Kind() {
			case MitchAlpha:
				_, seqCount := item.DefinedType.GetSequenceCount()
				_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v, n, err = stream.Read_%v(%v)\n", item.Declarator.Identifier(), item.DefinedType.GetStreamFunctionName(), seqCount))
				errorCheckFunc()
				break
			case MitchMessageNumber:
				_, _ = writer.WriteString(fmt.Sprintf("\tb, n, err := stream.Read_byte()\n"))
				i, _ := strconv.Atoi(item.DefinedType.DefaultValue())
				_, _ = writer.WriteString(fmt.Sprintf("\tif b != 0x%x {\n", i))
				_, _ = writer.WriteString(fmt.Sprintf("\t\treturn nil, 0, errors.New(fmt.Sprintf(\"Message type numbers does not match up. For Message %v was expected 0x%x, but 0x%%x was found.)\", b))\n", self.data.Identifier, i))
				_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
				errorCheckFunc()
				_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v = b\n", item.Declarator.Identifier()))
				break
			case Enum:
				_, _ = writer.WriteString(fmt.Sprintf("\ttemp_%v, n, err := stream.Read_byte()\n", item.Declarator.Identifier()))
				_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v = %v(temp_%v)\n", item.Declarator.Identifier(), item.DefinedType.GetName(), item.Declarator.Identifier()))
				errorCheckFunc()
				break
			case MitchBitField:
				if typeDecl, ok := item.DefinedType.(ITypeDeclaration); ok {
					_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v, n, err = Read_%v(stream)\n", item.Declarator.Identifier(), typeDecl.GetDeclarator().Identifier()))
					errorCheckFunc()
				} else {
					panic("DDDD")
				}
			default:
				_, _ = writer.WriteString(fmt.Sprintf("\tvalue.%v, n, err = stream.Read_%v()\n", item.Declarator.Identifier(), item.DefinedType.GetStreamFunctionName()))
				errorCheckFunc()
			}

		}
		_, _ = writer.WriteString(fmt.Sprintf("\treturn value, byteCount, nil\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self *publishStruct) GenerateWriteFunction(writer io.StringWriter, TypeInformation IBaseTypeInformation, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		returnType := "*" + typeNamePrefix
		_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
		_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))
		for index, item := range self.data.Members {
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t// Index: %v, Member Name: %v, Type: %v \n", index, item.Declarator.Identifier(), item.DefinedType.GetName()))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			switch item.DefinedType.Kind() {
			case MitchAlpha:
				_, seqCount := item.DefinedType.GetSequenceCount()
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_%v(value.%v, %v)\n", item.DefinedType.GetStreamFunctionName(), item.Declarator.Identifier(), seqCount))
			case MitchMessageNumber:
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_byte(%v)\n", item.DefinedType.DefaultValue()))
				break
			case Enum:
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_byte(byte(value.%v))\n", item.Declarator.Identifier()))
				break
			case MitchBitField:
				if typeDecl, ok := item.DefinedType.(ITypeDeclaration); ok {
					_, _ = writer.WriteString(fmt.Sprintf("\tn, err = Write_%v(stream, value.%v)\n", typeDecl.GetDeclarator().Identifier(), item.Declarator.Identifier()))
				} else {
					panic("DDDD")
				}
			default:
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_%v(value.%v)\n", item.DefinedType.GetStreamFunctionName(), item.Declarator.Identifier()))
			}
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
		}
		_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func NewPublishStruct(data *yacc.StructDefinition, typeInformation IBaseTypeInformation) *publishStruct {
	return &publishStruct{
		data:            data,
		typeInformation: typeInformation,
	}
}
