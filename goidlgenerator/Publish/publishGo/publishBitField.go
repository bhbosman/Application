package publishGo

import (
	"fmt"
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

func (self publishBitField) Export(writer io.StringWriter) {
	self.ExportDefinition(writer)
}

func NewpublishBitField(data *yacc.BitField, identifier string) *publishBitField {
	return &publishBitField{
		data:       data,
		Identifier: identifier,
	}
}
