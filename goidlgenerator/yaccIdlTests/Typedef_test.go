package yaccIdlTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/IdlDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTypeDef(t *testing.T) {
	verbose := false

	createContext := func() *yacc.IdlExprContext {
		return yacc.NewIdlExprContext()
	}

	t.Run("No Decl", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			NewType a;
		};`

		reader := bufio.NewReader(strings.NewReader(data))
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: createContext(),
				Verbose:        verbose})
		assert.Equal(t, yacc.DefNotFound, yacc.IdlExprParse(idlExprLex))

	})

	t.Run("No Decl", func(t *testing.T) {
		data := `
		typedef long NewType, abc;
		struct HelloWorld
		{
			NewType a;
		};`

		reader := bufio.NewReader(strings.NewReader(data))
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			IdlDefinedTypes.NewIdlNativeTypeInformation(),
			yacc.NewIdlExprLexParams{
				IdlExprContext: createContext(),
				Verbose:        verbose,
			})
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
	})
}
