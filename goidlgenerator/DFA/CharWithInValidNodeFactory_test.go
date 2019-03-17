package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharNodeWithInvalidNodeFactory(t *testing.T) {
	// Link an invalid node factory to this test
	NodeFactory.Replace(&InvalidNodeFactory{})
	// will remove it
	defer NodeFactory.Reset()

	t.Run("Valid Char as Input", func(t *testing.T) {
		_, err := NewCharNode(1234)
		assert.Error(t, err)

	})
}
