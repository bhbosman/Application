package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeFactory(t *testing.T) {
	t.Run("Check default node factory", func(t *testing.T) {
		t.Run("Plain NodeLink", func(t *testing.T) {
			nodeA := NewPlainNode("A", false)
			nodeB := NewPlainNode("B", true)
			err := NodeFactory.PlainNodeLink('a', nodeA, nodeB)
			if !assert.NoError(t, err) {
				return
			}
			node, ok := nodeA.nextNode['a']
			if !assert.True(t, ok) {
				return
			}
			if !assert.Equal(t, nodeB, node) {
				return
			}
		})

		t.Run("Plain Multi NodeLink", func(t *testing.T) {
			nodeA := NewPlainNode("A", false)
			nodeB := NewPlainNode("B", true)
			err := NodeFactory.PlainNodeMultiLink('a', 'c', nodeA, nodeB)
			if !assert.NoError(t, err) {
				return
			}
			checkPass := func(c byte) bool {
				node, ok := nodeA.nextNode[c]
				if !assert.True(t, ok) {
					return false
				}
				if !assert.Equal(t, nodeB, node) {
					return false
				}
				return true
			}

			if !checkPass('a') {
				assert.Fail(t, "Path is valid, but failed")
				return
			}
			if !checkPass('b') {
				assert.Fail(t, "Path is valid, but failed")
				return
			}
			if !checkPass('c') {
				assert.Fail(t, "Path is valid, but failed")
				return
			}

			checkFailure := func(c byte) bool {
				nodeFailure, ok := nodeA.nextNode[c]
				if !assert.False(t, ok) {
					return false
				}
				if !assert.Equal(t, nil, nodeFailure) {
					return false
				}
				return true
			}

			if !checkFailure('d') {
				assert.Fail(t, "Path is invalid, but passed")
				return
			}
		})

		t.Run("Exit node", func(t *testing.T) {
			nodeA := NewPlainNode("A", false)
			nodeB := NewPlainNode("B", false)

			err := NodeFactory.PlainExitNodeLink('a', nodeA, nodeB)
			assert.NoError(t, err)
			_, ok := nodeA.NextNode('a')
			assert.False(t, ok)

			nextNode, ok := nodeA.ExitNode('a')
			assert.True(t, ok)
			assert.Equal(t, nextNode, nodeB)
		})
	})

	t.Run("Check default node factory, with mistakes", func(t *testing.T) {
		t.Run("Plain NodeLink", func(t *testing.T) {
			nodeA := NewPlainNode("A", false)
			nodeB := NewPlainNode("B", true)
			err := NodeFactory.PlainNodeLink('a', nodeA, nodeB)
			if !assert.NoError(t, err) {
				return
			}
			err = NodeFactory.PlainNodeLink('a', nodeA, nodeB)
			if !assert.Error(t, err) {
				return
			}
		})

		t.Run("Plain Multi NodeLink", func(t *testing.T) {
			nodeA := NewPlainNode("A", false)
			nodeB := NewPlainNode("B", true)
			err := NodeFactory.PlainNodeMultiLink('a', 'c', nodeA, nodeB)
			if !assert.NoError(t, err) {
				return
			}

			err = NodeFactory.PlainNodeMultiLink('a', 'c', nodeA, nodeB)
			if !assert.Error(t, err) {
				return
			}
			err = NodeFactory.PlainNodeLink('c', nodeA, nodeB)
			if !assert.Error(t, err) {
				return
			}

		})

		t.Run("Exit node", func(t *testing.T) {
			nodeA := NewPlainNode("A", false)
			nodeB := NewPlainNode("B", false)
			err := NodeFactory.PlainExitNodeLink('a', nodeA,nodeB)
			assert.NoError(t, err)
			err = NodeFactory.PlainExitNodeLink('a', nodeA,nodeB)
			assert.Error(t, err)

		})
	})

	t.Run("Check override of other factory", func(t *testing.T) {
		NodeFactory.Replace(&InvalidNodeFactory{})
		defer NodeFactory.Reset()
		t.Run("Plain Node", func(t *testing.T) {
			err := NodeFactory.PlainNodeLink('a', nil, nil)
			assert.Error(t, err)
		})
		t.Run("Multi Plain Node", func(t *testing.T) {
			err := NodeFactory.PlainNodeMultiLink('a', 'b', nil, nil)
			assert.Error(t, err)
		})
		t.Run("Exit Plain Node", func(t *testing.T) {
			err := NodeFactory.PlainExitNodeLink('a', nil, nil)
			assert.Error(t, err)
		})

	})
}
