package DFA

import (
	"github.com/stretchr/assert"
	"testing"
)

func TestSingleLink(t *testing.T) {

	nodeA := NewPlainNode("A", false)
	nodeB := NewPlainNode("B", true)

	err := PlainNodeLink('a', nodeA, nodeB)
	assert.NoError(t, err)
	node, ok := nodeA.nextNode['a']
	assert.True(t, ok)
	assert.Equal(t, nodeB, node)
}
