package Go

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"io"
)

type publishGolang struct {
}

func (self *publishGolang) Export(outputStream io.Writer, declaredTypes []interfaces.IDefinitionDeclaration) {
	for _, declaredType := range declaredTypes {
		if structDefinition, ok :=  declaredType.(*yacc.StructDefinition); ok{
			structDefinition.Identifier
		}

	}
}

func newPublishGolang() *publishGolang {
	return &publishGolang{}
}

func init() {
	Publish.Register(Publish.Go, newPublishGolang())
}
