package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericSingleCharToken(t *testing.T) {
	t.Run("Check lexem", func(t *testing.T) {

		charNode, _ := NewGenericSingleCharToken("A", 'A', 'A')
		nodeWalker := NewNodeWalker(charNode)
		tokenReceived, _ := nodeWalker.Token(`'A'`)
		assert.Equal(t, int(byte('A')), tokenReceived)
	})
}
