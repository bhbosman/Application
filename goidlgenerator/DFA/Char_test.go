package DFA

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)



type FailStream struct {

}

func (self *FailStream) Write(p []byte) (n int, err error) {
	return 0, errors.New("Fa")
}

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
