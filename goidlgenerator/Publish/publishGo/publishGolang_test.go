package publishGo

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/MitchDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestPublishOfStructDefinition(t *testing.T) {
	verbose := false
	typeInformation := MitchDefinedTypes.NewMitchTypeInformation()
	//typeValueHelper := Extensions DefaultTypeValueHelper()
	t.Run("", func(t *testing.T) {
		data := `
			typedef MitchBitField
			<
				InverseOrderBook,
				b1,
				b2,
				b3,
				b4,
				b5,
				b6,
				b7
			> SymbolDirectoryFlags;

			struct A
			{
				SymbolDirectoryFlags B;
			};
		`

		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()

		lexParams := yacc.NewIdlExprLexParams{
			IdlExprContext: IdlExprContext,
			Verbose:        verbose,
		}
		idlExprLex, _ := yacc.NewIdlExprLex(reader, typeInformation, lexParams)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}

		params := Publish.ExportParams{
			OutputStream:  os.Stdout,
			PackageName:   "ddd",
			DeclaredTypes: DeclaredTypes,
		}
		err := newPublishGolang().Export(typeInformation, params)
		assert.NoError(t, err)
	})

}
