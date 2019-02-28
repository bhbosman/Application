package Go

import (
	"encoding/json"
	"github.com/bhbosman/gofintech/goidlgenerator/Publish"
	"github.com/bhbosman/gofintech/goidlgenerator/yacc"
	"io"
)

type publishGolang struct {
}

func (self *publishGolang) Export(outputStream io.Writer, declaredTypes []yacc.ITypeSpec) {
	for _, declaredType := range declaredTypes {
		bytes, err := json.MarshalIndent(declaredType, "", "\t")
		if err != nil {

		}
		_, _ = outputStream.Write(bytes)

	}
}

func newPublishGolang() *publishGolang {
	return &publishGolang{}
}

func init() {
	Publish.Register(Publish.Go, newPublishGolang())
}
