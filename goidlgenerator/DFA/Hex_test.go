package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexDfa(t *testing.T) {

	t.Run("CheckForHex", func(t *testing.T) {
		hex, err := NewHexDfa(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(hex)
		assert.True(t, WalkString(nodeWalker, []byte("0xABCD")))
	})
	t.Run("CheckForHex", func(t *testing.T) {
		hex, err := NewHexDfa(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(hex)
		assert.True(t, WalkString(nodeWalker, []byte("0xabcdef")))
	})
	t.Run("CheckForHex", func(t *testing.T) {
		hex, err := NewHexDfa(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(hex)
		assert.True(t, WalkString(nodeWalker, []byte("0x01234567890ABCDEF")))
	})
	t.Run("CheckForHex", func(t *testing.T) {
		hex, err := NewHexDfa(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(hex)
		assert.True(t, WalkString(nodeWalker, []byte("0X01234567890ABCDEF")))
	})

	t.Run("CheckForHex Token", func(t *testing.T) {
		hex, err := NewHexDfa(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(hex)
		tokenValue, _ := nodeWalker.Token("")
		assert.Equal(t, 1234, tokenValue)
	})

}
