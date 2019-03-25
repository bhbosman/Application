package publishGo

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/Application/Common"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
)

type publishGolang struct {
}

func (self *publishGolang) Export(TypeInformation interfaces.IBaseTypeInformation, params Publish.ExportParams) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		sb := bufio.NewWriter(params.OutputStream)

		//_, _ = sb.WriteString(fmt.Sprintf("// Code generated by goyacc -o idl.go -p IdlExpr idl.y. DO NOT EDIT.\n"))
		_, _ = sb.WriteString(fmt.Sprintf("\n"))
		_, _ = sb.WriteString(fmt.Sprintf("package %v\n", params.PackageName))

		_, _ = sb.WriteString(fmt.Sprintf("import \"errors\"\n"))
		_, _ = sb.WriteString(fmt.Sprintf("import \"fmt\"\n"))
		_, _ = sb.WriteString(fmt.Sprintf("import \"github.com/bhbosman/Application/Streams\"\n"))
		_, _ = sb.WriteString(fmt.Sprintf("\n"))

		_, _ = sb.WriteString(fmt.Sprintf("// Declared typed\n"))
		for _, declaredType := range params.DeclaredTypes {
			_, _ = sb.WriteString(fmt.Sprintf("// %v\n", declaredType.GetName()))
		}
		_, _ = sb.WriteString(fmt.Sprintf("// \n"))

		for _, declaredType := range params.DeclaredTypes {
			declaredType.GetName()
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
	})
}

func newPublishGolang() *publishGolang {
	return &publishGolang{}
}

func init() {
	Publish.Register(Publish.Go, newPublishGolang())
}
