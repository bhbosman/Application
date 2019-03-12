package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Extansions"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"io"
)

type GenerateMessageReadFunctionParams struct {
	TypeInformation interfaces.IBaseTypeInformation
	kind            interfaces.Kind
	typeNamePrefix  string
	typeCode        uint32
	defaultValue    string
}

func GenerateMessageReadFunction(writer io.StringWriter, definedType interfaces.IDefinedType, params GenerateMessageReadFunctionParams) {
	returnType := Extansions.TypeValueForDefinedType(definedType)
	_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func ReadMessage_%v(stream Streams.I%vReader) (%v, int, error) {\n", params.typeNamePrefix, params.TypeInformation.Name(), returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount := 0\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\ttypeCode, n, err := stream.Read_uint32()\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn %v, 0, err\n", params.defaultValue))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))

	_, _ = writer.WriteString(fmt.Sprintf("\tif typeCode != 0x%08x {\n", params.typeCode))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn %v, 0, errors.New(fmt.Sprintf(\"typecode mismatch, while reading %v. Expected 0x%08x, got 0x%%08x\", typeCode))\n", params.defaultValue, params.typeNamePrefix, params.typeCode))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))

	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))

	_, _ = writer.WriteString(fmt.Sprintf("\tvalue_%v, n, err := Read_%v(stream)\n", params.typeNamePrefix, params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn %v, 0, err\n", params.defaultValue))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))

	_, _ = writer.WriteString(fmt.Sprintf("\treturn value_%v, byteCount, nil\n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}
