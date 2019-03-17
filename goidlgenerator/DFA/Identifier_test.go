package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDfaIdentifier(t *testing.T) {

	t.Run("CheckForValidIdentifiers", func(t *testing.T) {
		dfa, err := NewIdentifier(1234, nil)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("ABCD")))
	})

	t.Run("CheckForValidIdentifiers", func(t *testing.T) {
		dfa, err := NewIdentifier(1234, nil)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("_A_B_CD")))
	})

	t.Run("CheckForValidIdentifiers", func(t *testing.T) {
		dfa, err := NewIdentifier(1234, nil)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("_0123")))
	})

	t.Run("CheckForValidIdentifiers", func(t *testing.T) {
		dfa, err := NewIdentifier(1234, nil)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("aAqdd123")))
	})

	t.Run("CheckForInvalidIdentifiers", func(t *testing.T) {
		dfa, err := NewIdentifier(1234, nil)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		b := WalkString(nodeWalker, []byte(("123a")))
		assert.False(t, b)
	})
	t.Run("CheckForInvalidIdentifiers", func(t *testing.T) {
		rv := make(map[string]int)
		rv["AA"] = 1
		rv["BB"] = 3
		dfa, err := NewIdentifier(1234, rv)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		b := WalkString(nodeWalker, []byte("walk"))
		assert.True(t, b)
		token01, _ := nodeWalker.Token("AA")
		assert.Equal(t, 1, token01)
		token02, _ := nodeWalker.Token("BB")
		assert.Equal(t, 3, token02)
		token03, _ := nodeWalker.Token("walk")
		assert.Equal(t, 1234, token03)

		assert.Equal(t, 2, dfa.GetMapCount())
		assert.Equal(t, 1, dfa.GetMapValue("AA"))
		assert.Equal(t, 3, dfa.GetMapValue("BB"))
	})

	t.Run("Check for invalid constructor", func(t *testing.T) {
		NodeFactory.Replace(&InvalidNodeFactory{})
		defer NodeFactory.Reset()
		_, err := NewIdentifier(1234, nil)
		assert.Error(t, err)
	})
}
