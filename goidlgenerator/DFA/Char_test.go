package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestCharNode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewCharNode(1234))
		//noinspection SpellCheckingInspection
		assert.True(t, WalkString(nodeWalker, []byte(`'A'`)))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewCharNode(1234))
		//noinspection SpellCheckingInspection
		assert.False(t, WalkString(nodeWalker, []byte(`'AB'`)))
	})

}
