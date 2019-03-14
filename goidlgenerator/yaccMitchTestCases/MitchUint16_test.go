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
		data, err := typeInformation.CreateType(interfaces.Sequence, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Native, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Struct, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Uint08, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Uint16, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Uint32, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Uint64, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Int08, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Int16, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Int32, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Int64, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.String, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Bool, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.WChar, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Char, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Double, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Fixed, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Float, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.LongDouble, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Octet, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.WideString, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Enum, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchAlpha, int64(1))
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchBitField, []string{"1", "2", "3", "4", "5", "6", "7", "8"})
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchByte, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchDate, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchTime, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchPrice04, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchPrice08, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt08, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt16, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt32, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt64, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)
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
				Verbose:        false,
			})
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
	})

}
