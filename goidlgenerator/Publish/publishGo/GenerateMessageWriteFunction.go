package publishGo

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Extansions"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"io"
)

type GenerateMessageWriteFunctionParams struct {
	TypeInformation interfaces.IBaseTypeInformation
	kind            interfaces.Kind
	typeNamePrefix  string
	typeCode        uint32
}

func GenerateMessageWriteFunction(writer io.StringWriter, definedType interfaces.IDefinedType, params GenerateMessageWriteFunctionParams) {
	returnType := Extansions.TypeValueForDefinedType(definedType)
	_, _ = writer.WriteString(fmt.Sprintf("// %v WriteMessage \n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func WriteMessage_%v(stream Streams.I%vWriter, value %v) (int, error) {\n", params.typeNamePrefix, params.TypeInformation.Name(), returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount := 0\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tn, err := stream.Write_uint32(0x%08x)\n", params.typeCode))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))

	_, _ = writer.WriteString(fmt.Sprintf("\tn, err = Write_%v(stream, value)\n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))

	_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}
