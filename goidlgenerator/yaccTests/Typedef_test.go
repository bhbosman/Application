package yaccTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/assert"
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
		idlExprLex := yacc.NewIdlExprLex(
			reader, createContext(),
			verbose)
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
		idlExprLex := yacc.NewIdlExprLex(reader, createContext(), verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
	})

}
