package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestHex_Name(t *testing.T) {

	t.Run("CheckForHex", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewHexValue(1234))
		assert.True(t, WalkString(nodeWalker, []byte("0xABCD")))
	})
	t.Run("CheckForHex", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewHexValue(1234))
		assert.True(t, WalkString(nodeWalker, []byte("0xabcdef")))
	})
	t.Run("CheckForHex", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewHexValue(1234))
		assert.True(t, WalkString(nodeWalker, []byte("0x01234567890ABCDEF")))
	})
	t.Run("CheckForHex", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewHexValue(1234))
		assert.True(t, WalkString(nodeWalker, []byte("0X01234567890ABCDEF")))
	})

}
