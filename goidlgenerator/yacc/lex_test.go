package yacc

import (
	"github.com/bhbosman/Application/goidlgenerator/DFA"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIdlTokens(t *testing.T) {
	// Some constant that may change in future
	currentLength := 16

	dfas, err := CreateIdlTokens()
	assert.NoError(t, err)
	assert.Equal(t, currentLength, len(dfas))

	dfamap := make(map[int]DFA.IDfa)
	for _, dfa := range dfas {
		tokenValue, _ := dfa.Token("(some arb value)")
		dfamap[tokenValue] = dfa
	}

	assert.Equal(t, currentLength, len(dfamap))
	t.Run("Check Identifier Token and reserved words", func(t *testing.T) {
		dfa, ok := dfamap[Identifier]
		if !ok {
			assert.Fail(t, "")
			return
		}
		identifierDfa, ok := dfa.(DFA.IIdentifierDfa)
		if !ok {
			assert.Equal(t, 38, identifierDfa.GetMapCount())
			return
		}
		assert.Equal(t, Rwlong, identifierDfa.GetMapValue("long"))
		assert.Equal(t, Rwboolean, identifierDfa.GetMapValue("boolean"))
		assert.Equal(t, Rwcase, identifierDfa.GetMapValue("case"))
		assert.Equal(t, Rwchar, identifierDfa.GetMapValue("char"))
		assert.Equal(t, Rwconst, identifierDfa.GetMapValue("const"))
		assert.Equal(t, Rwdefault, identifierDfa.GetMapValue("default"))
		assert.Equal(t, Rwdouble, identifierDfa.GetMapValue("double"))
		assert.Equal(t, Rwenum, identifierDfa.GetMapValue("enum"))
		assert.Equal(t, RwFALSE, identifierDfa.GetMapValue("FALSE"))
		assert.Equal(t, Rwfixed, identifierDfa.GetMapValue("fixed"))
		assert.Equal(t, Rwfloat, identifierDfa.GetMapValue("float"))
		assert.Equal(t, Rwlong, identifierDfa.GetMapValue("long"))
		assert.Equal(t, Rwmodule, identifierDfa.GetMapValue("module"))
		assert.Equal(t, Rwnative, identifierDfa.GetMapValue("native"))
		assert.Equal(t, Rwoctet, identifierDfa.GetMapValue("octet"))
		assert.Equal(t, Rwsequence, identifierDfa.GetMapValue("sequence"))
		assert.Equal(t, Rwshort, identifierDfa.GetMapValue("short"))
		assert.Equal(t, Rwstring, identifierDfa.GetMapValue("string"))
		assert.Equal(t, Rwstruct, identifierDfa.GetMapValue("struct"))
		assert.Equal(t, Rwswitch, identifierDfa.GetMapValue("switch"))
		assert.Equal(t, RwTRUE, identifierDfa.GetMapValue("TRUE"))
		assert.Equal(t, Rwtypedef, identifierDfa.GetMapValue("typedef"))
		assert.Equal(t, Rwunsigned, identifierDfa.GetMapValue("unsigned"))
		assert.Equal(t, Rwunion, identifierDfa.GetMapValue("union"))
		assert.Equal(t, Rwvoid, identifierDfa.GetMapValue("void"))
		assert.Equal(t, Rwwchar, identifierDfa.GetMapValue("wchar"))
		assert.Equal(t, Rwwstring, identifierDfa.GetMapValue("wstring"))
		assert.Equal(t, RwMitchAlpha, identifierDfa.GetMapValue("MitchAlpha"))
		assert.Equal(t, RwMitchBitField, identifierDfa.GetMapValue("MitchBitField"))
		assert.Equal(t, RwMitchByte, identifierDfa.GetMapValue("MitchByte"))
		assert.Equal(t, RwMitchDate, identifierDfa.GetMapValue("MitchDate"))
		assert.Equal(t, RwMitchTime, identifierDfa.GetMapValue("MitchTime"))
		assert.Equal(t, RwMitchPrice04, identifierDfa.GetMapValue("MitchPrice04"))
		assert.Equal(t, RwMitchPrice08, identifierDfa.GetMapValue("MitchPrice08"))
		assert.Equal(t, RwMitchUInt08, identifierDfa.GetMapValue("MitchUInt08"))
		assert.Equal(t, RwMitchUInt16, identifierDfa.GetMapValue("MitchUInt16"))
		assert.Equal(t, RwMitchUInt32, identifierDfa.GetMapValue("MitchUInt32"))
		assert.Equal(t, RwMitchUInt64, identifierDfa.GetMapValue("MitchUInt64"))
		assert.Equal(t, RwMitchMessageNumberType, identifierDfa.GetMapValue("MitchMessageNumberType"))
	})

	t.Run("Check Integer Token", func(t *testing.T) {
		dfa, ok := dfamap[Integer_literal]
		if !ok {
			assert.Fail(t, "Could not find integer literal")
			return
		}
	})
}
