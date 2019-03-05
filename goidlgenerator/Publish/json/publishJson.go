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

func (self *publishJson) Export(outputStream io.Writer, packageName string, declaredTypes []interfaces.IDefinitionDeclaration) error {
	fmt.Println(len(declaredTypes))
	bytes, err := json.MarshalIndent(declaredTypes, "", "\t")
	if err != nil {
		return err
	}
	_, err = outputStream.Write(bytes)
	return err
}

func newPublishJson() *publishJson {
	return &publishJson{}
}

func init() {
	Publish.Register(Publish.Json, newPublishJson())
}
