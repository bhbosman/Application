package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishEnum struct {
	data *yacc.EnumDecl
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

func (self publishEnum) Export(writer io.StringWriter) {
	self.ExportDefinition(writer)
}

func NewPublishEnum(data *yacc.EnumDecl) *publishEnum {
	return &publishEnum{
		data: data}
}
