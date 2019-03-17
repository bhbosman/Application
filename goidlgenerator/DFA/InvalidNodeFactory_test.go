package DFA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidNodeFactory(t *testing.T) {
	invalidNodeFactory := InvalidNodeFactory{}
	t.Run("", func(t *testing.T) {
		assert.Error(t, invalidNodeFactory.PlainNodeLink('a', nil, nil))
	})
	t.Run("", func(t *testing.T) {
		assert.Error(t, invalidNodeFactory.PlainNodeMultiLink('a', 'b', nil, nil))
	})
	t.Run("", func(t *testing.T) {
		assert.Error(t, invalidNodeFactory.PlainExitNodeLink('a', nil, nil))
	})
}
