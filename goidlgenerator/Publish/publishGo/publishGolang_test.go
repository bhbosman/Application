package publishGo

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/IdlDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/Publish"
	_ "github.com/bhbosman/Application/goidlgenerator/Publish/json"
	"os"

	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPublishOfStructDefinition(t *testing.T) {
	verbose := false

	t.Run("", func(t *testing.T) {
		data := `
			typedef bitfield
    		<
        		InverseOrderBook,
        		b1,
        		b2,
        		b3,
        		b4,
        		b5,
        		b6,
        		b7
			> SymbolDirectoryFlags, A, B;
		
		`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()

		lexParams := yacc.NewIdlExprLexParams{
			IdlExprContext: IdlExprContext,
			Verbose:        verbose,
		}

		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlDefinedTypes.NewIdlNativeTypeInformation(),lexParams)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}

		params := Publish.ExportParams{
			OutputStream:    os.Stdout,
			PackageName:     "ddd",
			DeclaredTypes:   DeclaredTypes,
		}
		err := newPublishGolang().Export(IdlDefinedTypes.NewIdlNativeTypeInformation(), params)
		assert.NoError(t, err)
	})

}
