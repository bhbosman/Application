package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestDfaInteger(t *testing.T) {

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		//noinspection SpellCheckingInspection
		assert.False(t, WalkString(nodeWalker, []byte("ABCD")))
	})

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("_")))
	})

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("+")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("-")))
	})

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("00")))
	})

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("01")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("02")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.False(t, WalkString(nodeWalker, []byte("09")))
	})

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.True(t, WalkString(nodeWalker, []byte("+1")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.True(t, WalkString(nodeWalker, []byte("-1")))
	})
	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.True(t, WalkString(nodeWalker, []byte("1")))
	})

	t.Run("", func(t *testing.T) {
		nodeWalker := NewNodeWalker(NewDfaInteger(1234))
		assert.True(t, WalkString(nodeWalker, []byte("12345")))
	})

}
