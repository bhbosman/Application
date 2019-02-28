package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestCharNode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)
		//noinspection SpellCheckingInspection
		assert.True(t, WalkString(nodeWalker, []byte(`'A'`)))
	})
	t.Run("", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)
		//noinspection SpellCheckingInspection
		assert.False(t, WalkString(nodeWalker, []byte(`'AB'`)))
	})

}
