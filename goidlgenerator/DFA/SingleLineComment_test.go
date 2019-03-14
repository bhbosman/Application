package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComments(t *testing.T) {

	t.Run("", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("//\n")))
	})
	t.Run("", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("//Comment with \t\n")))
	})
	t.Run("", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.False(t, WalkString(nodeWalker, []byte("//")))
	})

	t.Run("Check Token Value", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		tokenValue, _ := nodeWalker.Token("(any token)")
		assert.Equal(t, 1234, tokenValue)
	})
}
