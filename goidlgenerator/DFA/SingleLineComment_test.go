package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestComments(t *testing.T) {

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewSingleLineComment(1234))
		assert.True(t, WalkString(nodeWalker, []byte("//\n")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewSingleLineComment(1234))
		assert.True(t, WalkString(nodeWalker, []byte("//Comment with \t\n")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewSingleLineComment(1234))
		assert.False(t, WalkString(nodeWalker, []byte("//")))
	})

}
