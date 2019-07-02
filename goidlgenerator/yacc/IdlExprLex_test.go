package yacc

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIdlExprLex(t *testing.T) {
	verbose := false

	t.Run("One Enum Value assigned int", func(t *testing.T) {
		data := `enum ABC {AA='a'};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := NewIdlExprContext()
		idlExprLex, _ := NewIdlExprLex(
			reader,
			NewIdlExprLexParams{
				IdlExprContext: ctx,
				Verbose:        verbose})
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})

	t.Run("", func(t *testing.T) {
		a := byte(65)

		fmt.Println(string(a))
	})

}
