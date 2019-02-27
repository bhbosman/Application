package yaccTests

import (
	"bufio"
	"github.com/bhbosman/gofintech/goidlgenerator/yacc"
	"github.com/stretchr/assert"
	"strings"
	"testing"
)

func TestEnums(t *testing.T) {
	verbose := false

	createContext := func() *yacc.IdlExprContext {
		return yacc.NewIdlExprContext()
	}

	t.Run("2", func(t *testing.T) {
		data := `
		enum ABC {
			AA,
			BB};
		`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex := yacc.NewIdlExprLex(reader, ctx, verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)){
			return
		}
		assert.Len(t, ctx.GetSpecification(), 1)
	})


	t.Run("4", func(t *testing.T) {
		data := `
		enum ABC {
			AA=12,
			BB,
			cc=333,
			DD};
		`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := createContext()
		idlExprLex := yacc.NewIdlExprLex(reader, ctx, verbose)
		if !assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex)){
			return
		}
		assert.Len(t, ctx.GetSpecification(), 1)
	})
}
