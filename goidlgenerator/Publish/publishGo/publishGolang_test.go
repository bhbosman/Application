package publishGo

import (
	"bufio"
	_ "github.com/bhbosman/Application/goidlgenerator/Publish/json"
	"os"

	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPublishOfStructDefinition(t *testing.T) {
	verbose := false

	t.Run("WithBooleanType", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			long a = "ssss";
			long a = 123;
			long a = 123;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(reader, IdlExprContext, verbose)
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
		DeclaredTypes := IdlExprContext.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}

		err := newPublishGolang().Export(os.Stdout, "dddd", DeclaredTypes)
		assert.NoError(t, err)
	})


}