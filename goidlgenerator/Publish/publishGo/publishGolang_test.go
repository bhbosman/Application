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
			> SymbolDirectoryFlags, sss, ssss;
		
		`
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
