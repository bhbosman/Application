package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishStruct struct {
	data *yacc.StructDefinition
}

func (self publishStruct) ExportDefinition(writer io.StringWriter) {
	_, _ = writer.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
	for _, member := range self.data.Members {
		defaultValue := member.Declarator.DefaultValue()

		if defaultValue == nil {
			_, _ = writer.WriteString(fmt.Sprintf(
				"\t%v %v\n",
				member.Declarator.Identifier(),
				member.DefinedType.GetName()))
		} else {

			ss := defaultValue.ValueKind().String()
			_, _ = writer.WriteString(fmt.Sprintf(
				"\t%v %v // default value: %v(%v)\n ",
				member.Declarator.Identifier(),
				member.DefinedType.GetName(),
				ss,
				GetExportValue(defaultValue)))
		}
	}
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishStruct) Export(writer io.StringWriter) {
	self.ExportDefinition(writer)
}

func NewPublishStruct(data *yacc.StructDefinition) *publishStruct {
	return &publishStruct{data: data}
}
