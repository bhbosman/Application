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


const decl string  =
`
//
// {{.GetScopeName}}
//
struct {{.GetScopeName}} 
{
	{{range .Members}}{{.Declarator.Identifier}} {{.DefinedType.GetName}};
	{{end}}
};

`




func (self *publishGolang) Export(outputStream io.Writer, packageName string, declaredTypes []interfaces.IDefinitionDeclaration) error{

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
