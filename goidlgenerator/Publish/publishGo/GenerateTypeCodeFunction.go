package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/Common"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"io"
)

func GenerateTypeCodeFunction(writer io.StringWriter, kind interfaces.Kind, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("// %v TypeCode\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func TypeCode_%v() uint32 {\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn 0x%08x\n", typeCode))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}
