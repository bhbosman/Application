package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeWalker(t *testing.T) {
	createIntegerDfa := func() IDfa {
		dfa, err := NewDfaInteger(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)
		return dfa

	}

	t.Run("reset", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		walker.Reset()
		assert.Equal(t, dfa.StartNode(), walker.currentNode)
	})
	t.Run("TokenValue", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		v1, l1 := dfa.Token("123")
		v2, l2 := walker.Token("123")
		assert.Equal(t, v1, v2)
		assert.Equal(t, l1, l2)
	})
	t.Run("Terminal", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		walker.Reset()
		assert.Equal(t, false, dfa.StartNode().Terminal())
		assert.Equal(t, dfa.StartNode().Terminal(), walker.Terminal())
	})

	t.Run("Walk 1", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		walker.Reset()
		b, e := walker.WalkNode('1')
		assert.True(t, b)
		assert.NoError(t, e)

		integerDfa, ok := dfa.(IIntegerDfa)
		if !assert.True(t, ok) {
			return
		}
		assert.Equal(t, integerDfa.TerminalNode(), walker.currentNode)
	})

	t.Run("Walk +1", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		walker.Reset()
		b, e := walker.WalkNode('+')
		assert.True(t, b)
		assert.NoError(t, e)

		integerDfa, ok := dfa.(IIntegerDfa)
		if !assert.True(t, ok) {
			return
		}
		assert.Equal(t, integerDfa.SignNode(), walker.currentNode)

		assert.Equal(t, integerDfa.SignNode(), walker.currentNode)
		b, e = walker.WalkNode('1')
		assert.True(t, b)
		assert.NoError(t, e)
		assert.Equal(t, integerDfa.TerminalNode(), walker.currentNode)
	})

	t.Run("Walk -1", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		walker.Reset()
		b, e := walker.WalkNode('-')
		assert.True(t, b)
		assert.NoError(t, e)

		integerDfa, ok := dfa.(IIntegerDfa)
		if !assert.True(t, ok) {
			return
		}

		assert.Equal(t, integerDfa.SignNode(), walker.currentNode)
		b, e = walker.WalkNode('1')
		assert.True(t, b)
		assert.NoError(t, e)
		assert.Equal(t, integerDfa.TerminalNode(), walker.currentNode)
	})

	t.Run("Walk -A", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		walker.Reset()
		b, e := walker.WalkNode('-')
		assert.True(t, b)
		assert.NoError(t, e)

		integerDfa, ok := dfa.(IIntegerDfa)
		if !assert.True(t, ok) {
			return
		}
		assert.Equal(t, integerDfa.SignNode(), walker.currentNode)
		b, e = walker.WalkNode('A')
		assert.False(t, b)
		assert.NoError(t, e)
		assert.True(t, walker.invalid)
	})

	t.Run("Walk past valid", func(t *testing.T) {
		dfa := createIntegerDfa()
		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		assert.False(t, walker.invalid)

		walker.Reset()
		b, e := walker.WalkNode('A')
		assert.False(t, b)
		assert.NoError(t, e)

		assert.True(t, walker.invalid)

		b, e = walker.WalkNode('A')
		assert.False(t, b)
		assert.NoError(t, e)
	})

	t.Run("Check exit node", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)

		walker := NewNodeWalker(dfa)
		assert.NotNil(t, walker)

		b, e := walker.WalkNode('/')
		assert.True(t, b)
		assert.NoError(t, e)
		assert.Equal(t, dfa.SecondNode(), walker.currentNode)

		b, e = walker.WalkNode('/')
		assert.True(t, b)
		assert.NoError(t, e)
		assert.Equal(t, dfa.OtherNode(), walker.currentNode)

		b, e = walker.WalkNode('f') // random char, not '\n'
		assert.True(t, b)
		assert.NoError(t, e)
		assert.Equal(t, dfa.OtherNode(), walker.currentNode)

		// trigger exit node
		b, e = walker.WalkNode('\n')
		assert.True(t, b)
		assert.NoError(t, e)
	})
}
