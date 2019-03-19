package json

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type publishJson struct {
}

func (self *publishJson) Export(TypeInformation interfaces.IBaseTypeInformation, params Publish.ExportParams) error {
	fmt.Println(len(params.DeclaredTypes))
	bytes, err := json.MarshalIndent(params.DeclaredTypes, "", "\t")
	if err != nil {
		return err
	}
	_, err = params.OutputStream.Write(bytes)
	return err
}

func newPublishJson() *publishJson {
	return &publishJson{}
}

func init() {
	Publish.Register(Publish.Json, newPublishJson())
}
