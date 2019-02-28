package yaccTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/assert"
	"strings"
	"testing"
)

func TestBitField(t *testing.T) {
	verbose := false

	createContext := func() *yacc.IdlExprContext {
		return yacc.NewIdlExprContext()
	}

	t.Run("No Decl", func(t *testing.T) {
		data := `typedef bitfield<a0, a1, a2, a3, a4, a5, a6, a7> newBitType;`

		reader := bufio.NewReader(strings.NewReader(data))
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader, createContext(),
			verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
	})

}
