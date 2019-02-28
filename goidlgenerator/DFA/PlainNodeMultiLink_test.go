package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestMultiLink(t *testing.T) {

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
}
