package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharNodeWithNodeFactory(t *testing.T) {

	t.Run("Check lexem", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)
		tokenReceived, lexem := nodeWalker.Token(`'A'`)
		assert.Equal(t, 1234, tokenReceived)
		assert.Equal(t, "A", lexem)

	})

	t.Run("Valid Char as Input", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)

		assert.True(t, WalkString(nodeWalker, []byte(`'A'`)))
	})
	t.Run("Valid Char as Input", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)

		assert.True(t, WalkString(nodeWalker, []byte(`'\\'`)))
	})
	t.Run("Valid Char as Input", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)

		assert.True(t, WalkString(nodeWalker, []byte(`'\t'`)))
	})
	t.Run("Valid Char as Input", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)

		assert.True(t, WalkString(nodeWalker, []byte(`'\n'`)))
	})
	t.Run("Valid Char as Input", func(t *testing.T) {
		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)

		assert.True(t, WalkString(nodeWalker, []byte(`'\r'`)))
	})

	t.Run("InValid char as input", func(t *testing.T) {

		charNode, _ := NewCharNode(1234)
		nodeWalker := NewNodeWalker(charNode)

		assert.False(t, WalkString(nodeWalker, []byte(`'AB'`)))
	})

}
