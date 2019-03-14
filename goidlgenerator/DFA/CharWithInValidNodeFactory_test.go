package DFA

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type InvalidNodeFactory struct {
}

func (Self *InvalidNodeFactory) PlainNodeLink(a byte, start, end IPlainNode) error {
	return fmt.Errorf("error for testing")
}

func (Self *InvalidNodeFactory) PlainExitNodeLink(a byte, start, end IPlainNode) error {
	return fmt.Errorf("error for testing")
}

func (Self *InvalidNodeFactory) PlainNodeMultiLink(a, b byte, start, end IPlainNode) error {
	return fmt.Errorf("error for testing")
}

func TestCharNodeWithInvalidNodeFactory(t *testing.T) {
	t.Run("Valid Char as Input", func(t *testing.T) {
		_, err := NewCharNode(1234)
		assert.Error(t, err)

	})
}

