package yaccIdlTests

import (
	"bufio"
	"github.com/bhbosman/Application/goidlgenerator/MitchDefinedTypes"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"github.com/bhbosman/Application/goidlgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMitchTypeInformation(t *testing.T) {
	typeInformation := MitchDefinedTypes.NewMitchTypeInformation()
	t.Run("CheckCreateType", func(t *testing.T) {
		assert.Nil(t, typeInformation.CreateType(interfaces.Sequence, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Native, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Struct, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Uint08, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Uint16, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Uint32, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Uint64, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Int08, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Int16, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Int32, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Int64, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.String, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Bool, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.WChar, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Char, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Double, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Fixed, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Float, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.LongDouble, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Octet, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.WideString, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.Enum, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.BitField, 0))
		assert.Nil(t, typeInformation.CreateType(interfaces.TypeDeclarator, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchAlpha, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchBitField, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchByte, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchDate, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchTime, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchPrice04, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchPrice08, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchUInt08, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchUInt16, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchUInt32, 0))
		assert.NotNil(t, typeInformation.CreateType(interfaces.MitchUInt64, 0))
	})
	t.Run("Struct With MitchUInt16", func(t *testing.T) {
		data := `
		struct HelloWorld
		{
			MitchUInt16 a;
		};`
		reader := bufio.NewReader(strings.NewReader(data))
		IdlExprContext := yacc.NewIdlExprContext()
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			typeInformation,
			yacc.NewIdlExprLexParams{
				IdlExprContext: IdlExprContext,
				Verbose:        false})
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
	})








}
