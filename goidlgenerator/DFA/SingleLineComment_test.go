package DFA

import (
	"github.com/bhbosman/Application/Common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComments(t *testing.T) {
	t.Run("Constructor error", func(t *testing.T) {
		Common.ErrorListFactory.Replace(func() Common.IErrorList {
			return &errorListWhichWillReturnError{}
		})
		defer func() {
			Common.ErrorListFactory.Reset()
		}()
		_, err := NewSingleLineComment(1234)
		assert.Error(t, err)
	})

	t.Run("Check IDfa interface", func(t *testing.T) {
		var obj interface{}
		var err error
		obj, err = NewSingleLineComment(1234)
		assert.NoError(t, err)
		assert.NotNil(t, obj)
		dfa, ok := obj.(IDfa)
		assert.True(t, ok)
		assert.NotNil(t, dfa)
	})

	t.Run("Check iSingleLineComment interface", func(t *testing.T) {
		var obj interface{}
		var err error
		obj, err = NewSingleLineComment(1234)
		assert.NoError(t, err)
		assert.NotNil(t, obj)
		dfa, ok := obj.(iSingleLineComment)
		assert.True(t, ok)
		assert.NotNil(t, dfa)
	})

	t.Run("", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("//\n")))
	})

	t.Run("", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.True(t, WalkString(nodeWalker, []byte("//Comment with \t\n")))
	})

	t.Run("", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		assert.False(t, WalkString(nodeWalker, []byte("//")))
	})

	t.Run("Check Token Value", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		nodeWalker := NewNodeWalker(dfa)
		tokenValue, _ := nodeWalker.Token("(any token)")
		assert.Equal(t, 1234, tokenValue)
	})

	t.Run("Check SecondNode Method", func(t *testing.T) {
		dfa, err := NewSingleLineComment(1234)
		assert.NoError(t, err)
		assert.NotNil(t, dfa)

		assert.Equal(t, dfa.SecondNode(), dfa.secondNode)
		assert.Equal(t, dfa.OtherNode(), dfa.otherNode)

	})
}
