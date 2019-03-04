package Go

import (

	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
	"text/template"
)

type publishGolang struct {
}

func (self *publishGolang) Export(outputStream io.Writer, declaredTypes []interfaces.IDefinitionDeclaration) error{
	decl := `struct {{.GetName()}} 
			{};`


	templateDeclaration, err := template.New("structDef").Parse(decl)

	if err != nil {
		return err
	}
	for _, declaredType := range declaredTypes {
		if structDefinition, ok :=  declaredType.(*yacc.StructDefinition); ok{
			err = templateDeclaration.Execute(outputStream, structDefinition)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func newPublishGolang() *publishGolang {
	return &publishGolang{}
}

func init() {
	Publish.Register(Publish.Go, newPublishGolang())
}
