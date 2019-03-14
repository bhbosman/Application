package publishGo

import (
	"github.com/bhbosman/Application/goidlgenerator/MitchDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishTypeDecl struct {
	data            *yacc.TypeDeclarator
	typeInformation interfaces.IBaseTypeInformation
}

func (self *publishTypeDecl) ExportDefinition(writer io.StringWriter) {
	//_, _ = writer.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
	//for _, member := range self.data.Members {
	//	defaultValue := member.Declarator.DefaultValue()
	//
	//	if defaultValue == nil {
	//		_, _ = writer.WriteString(fmt.Sprintf(
	//			"\t%v %v\n",
	//			member.Declarator.Identifier(),
	//			member.DefinedType.GetName()))
	//	} else {
	//
	//		ss := defaultValue.ValueKind().String()
	//		_, _ = writer.WriteString(fmt.Sprintf(
	//			"\t%v %v // default value: %v(%v)\n ",
	//			member.Declarator.Identifier(),
	//			member.DefinedType.GetName(),
	//			ss,
	//			GetExportValue(defaultValue)))
	//	}
	//}
	//_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	//_, _ = writer.WriteString(fmt.Sprintf("\n"))}

}

func (self *publishTypeDecl) Export(writer io.StringWriter, TypeInformation interfaces.IBaseTypeInformation) {
	if definition, ok := self.data.DefinedTyped.(*MitchDefinedTypes.MitchBitField); ok {
		publish := NewPublishBitField(definition, self.data.Declarator.Identifier())
		publish.Export(writer, TypeInformation)
	}
}

func NewpublishTypeDecl(data *yacc.TypeDeclarator, typeInformation interfaces.IBaseTypeInformation) *publishTypeDecl {
	return &publishTypeDecl{
		data:            data,
		typeInformation: typeInformation,
	}
}
