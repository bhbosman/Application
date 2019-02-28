package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestStringNode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewStringNode(1234))
		//noinspection SpellCheckingInspection
		assert.True(t, WalkString(nodeWalker, []byte(`"ABCD"`)))
	})

}
