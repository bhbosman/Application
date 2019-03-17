package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDfaInteger(t *testing.T) {
	if !t.Run("Check Interface IDfa", func(t *testing.T) {
		var dfa interface{}
		var err error
		dfa, err = NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		_, ok := dfa.(IDfa)
		assert.True(t, ok)
	}) {
		return
	}
	if !t.Run("Check Interface IIntegerDfa", func(t *testing.T) {
		var dfa IDfa
		var err error
		dfa, err = NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		_, ok := dfa.(IIntegerDfa)
		assert.True(t, ok)
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		//noinspection SpellCheckingInspection
		assert.False(t, WalkString(nodeWalker, []byte("ABCD")))
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		assert.False(t, WalkString(nodeWalker, []byte("_")))
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		assert.False(t, WalkString(nodeWalker, []byte("+")))
	}) {
		return
	}
	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		assert.False(t, WalkString(nodeWalker, []byte("-")))
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		assert.False(t, WalkString(nodeWalker, []byte("00")))
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		assert.False(t, WalkString(nodeWalker, []byte("01")))
	}) {
		return
	}
	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)
		assert.False(t, WalkString(nodeWalker, []byte("02")))
	}) {
		return
	}
	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)
		assert.False(t, WalkString(nodeWalker, []byte("09")))
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("+1")))
	}) {
		return
	}
	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("-1")))
	}) {
		return
	}
	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("1")))
	}) {
		return
	}

	if !t.Run("", func(t *testing.T) {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("12345")))
	}) {
		return
	}

	if !t.Run("Check Token Value", func(t *testing.T) {
		inputTokenValue := 1234
		dfa, err := NewDfaInteger(inputTokenValue)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		nodeWalker := NewNodeWalker(dfa)

		inputLexemValue := "(any lexem)"
		assert.True(t, WalkString(nodeWalker, []byte("12345")))
		token, lexem := nodeWalker.Token(inputLexemValue)
		assert.Equal(t, inputLexemValue, lexem)
		assert.Equal(t, inputTokenValue, token)

	}) {
		return
	}

	if !t.Run("Constructor error", func(t *testing.T) {
		NodeFactory.Replace(&InvalidNodeFactory{})
		defer NodeFactory.Reset()
		inputTokenValue := 1234
		_, err := NewDfaInteger(inputTokenValue)
		assert.Error(t, err)
	}) {
		return
	}

	t.Run("Check SignNode Method", func(t *testing.T) {
		inputTokenValue := 1234
		dfa, err := NewDfaInteger(inputTokenValue)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		assert.Equal(t, dfa.signNode, dfa.SignNode())
	})

	t.Run("Check Terminal Method", func(t *testing.T) {
		inputTokenValue := 1234
		dfa, err := NewDfaInteger(inputTokenValue)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		assert.Equal(t, dfa.terminalNode, dfa.TerminalNode())
	})

}
