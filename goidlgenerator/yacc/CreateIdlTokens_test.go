package yacc

import (
	"github.com/bhbosman/Application/goidlgenerator/DFA"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIdlTokens(t *testing.T) {
	// Some constant that may change in future
	currentLength := 15

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
		if ok {
			assert.Equal(t, 14, identifierDfa.GetMapCount())
			return
		}
		assert.Equal(t, Rwenum, identifierDfa.GetMapValue("enum"))
		assert.Equal(t, Rwstruct, identifierDfa.GetMapValue("struct"))
		assert.Equal(t, Rwtypedef, identifierDfa.GetMapValue("typedef"))
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
	})
}
