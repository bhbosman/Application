package publishGo

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
	"reflect"
)


func GetExportValue(value interfaces.IConstantValue) string {
	if value.GetType() == reflect.TypeOf("stringType"){
		return fmt.Sprintf("\"%v\"", value.GetValue())
	}
	return fmt.Sprintf("%v", value.GetValue())

}

type PublishStruct struct {
	data *yacc.StructDefinition
}


func (self PublishStruct) Export(writer io.Writer)  {
	sb := bufio.NewWriter(writer)
	_, _ = sb.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
	for _, member := range self.data.Members {
		defaultValue := member.Declarator.GetDefaultValue()

		if defaultValue == nil{
			_, _ = sb.WriteString(fmt.Sprintf(
				"\t%v %v\n",
				member.Declarator.Identifier(),
				member.DefinedType.GetName()))
		}else{
			ss := defaultValue.GetType().String()
			_, _ = sb.WriteString(fmt.Sprintf(
				"\t%v %v // default value: %v(%v)\n ",
				member.Declarator.Identifier(),
				member.DefinedType.GetName(),
				ss,
				GetExportValue(defaultValue)))
		}

	}
	_, _ = sb.WriteString(fmt.Sprintf("}\n"))
	_, _ = sb.WriteString(fmt.Sprintf("\n"))
	_ = sb.Flush()
}


func NewPublishStruct(data *yacc.StructDefinition) *PublishStruct {
	return &PublishStruct{data: data}
}

type publishGolang struct {
}


func (self *publishGolang) Export(outputStream io.Writer, packageName string, declaredTypes []interfaces.IDefinitionDeclaration) error {


	for _, declaredType := range declaredTypes {
		if definition, ok := declaredType.(*yacc.StructDefinition); ok {
			publish := NewPublishStruct(definition)
			publish.Export(outputStream)
			continue
		}

		//if definition, ok := declaredType.(*yacc.EnumDecl); ok {
		//	continue
		//}




	}
	return nil
}

func newPublishGolang() *publishGolang {
	return &publishGolang{}
}

func init() {
	Publish.Register(Publish.Go, newPublishGolang())
}
