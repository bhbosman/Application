package publishGo

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
)

func GetExportValue(value interfaces.IConstantValue) string {
	if value.ValueKind() == interfaces.String {
		if value.MaxLength() == 1 {
			return fmt.Sprintf("'%v'", value.Value())
		}
		return fmt.Sprintf("\"%v\"", value.Value())
	}

	if value.ValueKind() == interfaces.Char {
		return fmt.Sprintf("'%v'", value.Value())
	}

	return fmt.Sprintf("%v", value.Value())
}

type publishGolang struct {
}

func (self *publishGolang) Export(TypeInformation interfaces.IBaseTypeInformation, params Publish.ExportParams) error {
	sb := bufio.NewWriter(params.OutputStream)

	//_, _ = sb.WriteString(fmt.Sprintf("// Code generated by goyacc -o idl.go -p IdlExpr idl.y. DO NOT EDIT.\n"))
	_, _ = sb.WriteString(fmt.Sprintf("\n"))
	_, _ = sb.WriteString(fmt.Sprintf("package %v\n", params.PackageName))

	_, _ = sb.WriteString(fmt.Sprintf("import \"errors\"\n"))
	_, _ = sb.WriteString(fmt.Sprintf("import \"fmt\"\n"))
	_, _ = sb.WriteString(fmt.Sprintf("import \"github.com/bhbosman/Application/Streams\"\n"))

	_, _ = sb.WriteString(fmt.Sprintf("\n"))

	for _, declaredType := range params.DeclaredTypes {

		if declaredType.Predefined() {
			publish := NewPublishPredefined(declaredType)
			publish.Export(sb, TypeInformation)
			continue
		}

		if definition, ok := declaredType.(*yacc.StructDefinition); ok {
			publish := NewPublishStruct(definition, TypeInformation)
			publish.Export(sb, TypeInformation)
			continue
		}

		if definition, ok := declaredType.(*yacc.EnumDecl); ok {
			publish := NewPublishEnum(definition, TypeInformation)
			publish.Export(sb, TypeInformation)
			continue
		}
		if typeDeclarator, ok := declaredType.(*yacc.TypeDeclarator); ok {
			publish := NewpublishTypeDecl(typeDeclarator, TypeInformation)
			publish.Export(sb, TypeInformation)
			continue
		}
	}
	_ = sb.Flush()
	return nil
}

func newPublishGolang() *publishGolang {
	return &publishGolang{}
}

func init() {
	Publish.Register(Publish.Go, newPublishGolang())
}
