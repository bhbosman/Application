package json

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"io"
)

type publishJson struct {
}

func (self *publishJson) Export(outputStream io.Writer, declaredTypes []interfaces.IDefinitionDeclaration) {
	fmt.Println(len(declaredTypes))
	bytes, err := json.MarshalIndent(declaredTypes, "", "\t")
	if err != nil {
	}
	_, _ = outputStream.Write(bytes)
}

func newPublishJson() *publishJson {
	return &publishJson{}
}

func init() {
	Publish.Register(Publish.Json, newPublishJson())
}
