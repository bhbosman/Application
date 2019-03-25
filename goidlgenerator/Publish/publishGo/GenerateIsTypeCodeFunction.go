package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/Common"
	"io"
)

func GenerateIsTypeCodeFunction(writer io.StringWriter, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("// %v IsTypeCode\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func IsTypeCode_%v(typeCode uint32) bool {\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn typeCode == 0x%08x\n", typeCode))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}
