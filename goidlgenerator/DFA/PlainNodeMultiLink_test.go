package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiLink(t *testing.T) {
	t.Run("Add correctly", func(t *testing.T) {
		nodeA := NewPlainNode("A", false)
		nodeB := NewPlainNode("B", true)

		err := PlainNodeMultiLink('a', 'd', nodeA, nodeB)
		assert.NoError(t, err)

		node, ok := nodeA.nextNode['a']
		assert.True(t, ok)
		assert.Equal(t, nodeB, node)

		node, ok = nodeA.nextNode['b']
		assert.True(t, ok)
		assert.Equal(t, nodeB, node)

		node, ok = nodeA.nextNode['c']
		assert.True(t, ok)
		assert.Equal(t, nodeB, node)

		node, ok = nodeA.nextNode['d']
		assert.True(t, ok)
		assert.Equal(t, nodeB, node)

		node, ok = nodeA.nextNode['z']
		assert.False(t, ok)
	})

	t.Run("Add Incorrect", func(t *testing.T) {
		nodeA := NewPlainNode("A", false)
		nodeB := NewPlainNode("B", true)
		err := PlainNodeMultiLink('a', 'd', nodeA, nodeB)
		assert.NoError(t, err)
		assert.Equal(t, 4, nodeA.NextNodeCount())
		err = PlainNodeMultiLink('d', 'e', nodeA, nodeB)
		assert.Error(t, err)
		assert.Equal(t, 4, nodeA.NextNodeCount())
	})

}
