package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Extansions"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishBitField struct {
	data       *yacc.BitField
	Identifier string
}

func (self *publishBitField) ExportDefinition(writer io.StringWriter) {
	_, _ = writer.WriteString(fmt.Sprintf("type %s struct{\n", self.Identifier))

	if self.data.BitsUsed&0x01 == 0x01 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 00\n", self.data.BitField00))
	}
	if self.data.BitsUsed&0x02 == 0x02 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 01\n", self.data.BitField01))
	}
	if self.data.BitsUsed&0x04 == 0x04 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 02\n", self.data.BitField02))
	}
	if self.data.BitsUsed&0x08 == 0x08 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 03\n", self.data.BitField03))
	}
	if self.data.BitsUsed&0x10 == 0x10 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 04\n", self.data.BitField04))
	}
	if self.data.BitsUsed&0x20 == 0x20 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 05\n", self.data.BitField05))
	}
	if self.data.BitsUsed&0x40 == 0x40 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 06\n", self.data.BitField06))
	}
	if self.data.BitsUsed&0x80 == 0x80 {
		_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 07\n", self.data.BitField07))
	}
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishBitField) Export(writer io.StringWriter, TypeInformation interfaces.IBaseTypeInformation) {
	self.ExportDefinition(writer)

	typeNamePrefix := self.data.Identifier
	typeCode := CalculateCrc(typeNamePrefix)
	_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))


	GenerateMessageWriteFunction(
		writer,
		self.data,
		GenerateMessageWriteFunctionParams{
			TypeInformation: TypeInformation,
			kind:            interfaces.BitField,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
		})
	GenerateMessageReadFunction(
		writer,
		self.data,
		GenerateMessageReadFunctionParams{
			TypeInformation: TypeInformation,
			kind:            interfaces.BitField,
			typeNamePrefix:  typeNamePrefix,
			typeCode:        typeCode,
			defaultValue:    "0",
		})

	self.GenerateWriteFunction(writer, TypeInformation, typeNamePrefix, typeCode)
	self.GenerateReadFunction(writer, TypeInformation, typeNamePrefix, typeCode)
}


func (self publishBitField) GenerateReadFunction(writer io.StringWriter, TypeInformation interfaces.IBaseTypeInformation, typeNamePrefix string, typeCode uint32) {
	returnType := Extansions.TypeValueForDefinedType(self.data)
	_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func Read_%v(stream Streams.I%vReader) (value %v, byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))

	_, _ = writer.WriteString(fmt.Sprintf("\tvalue = New%v()\n", typeNamePrefix))

	_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))

	_, _ = writer.WriteString(fmt.Sprintf("\treturn value, byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishBitField) GenerateWriteFunction(writer io.StringWriter, TypeInformation interfaces.IBaseTypeInformation, typeNamePrefix string, typeCode uint32) {
	returnType := Extansions.TypeValueForDefinedType(self.data)
	_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (byteCount int, err error) {\n", typeNamePrefix, TypeInformation.Name(), returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))

	_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}


func NewpublishBitField(data *yacc.BitField, identifier string) *publishBitField {
	return &publishBitField{
		data:       data,
		Identifier: identifier,
	}
}
