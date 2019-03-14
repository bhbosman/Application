package DFA

import (
	"errors"
	"github.com/bhbosman/Application/Generic"
	"github.com/stretchr/testify/assert"
	"testing"
)

type errorListWhichWillReturnError struct {
	
}

func (self *errorListWhichWillReturnError) Add(err error) {
	// do nothing
}

func (self *errorListWhichWillReturnError) Error() error {
	return errors.New("an error")
}

func TestNegativeHex(t *testing.T) {
	Generic.ErrorListFactory.Replace(func() Generic.IErrorList {
		return &errorListWhichWillReturnError{}
	})
	defer func() {
		Generic.ErrorListFactory.Reset()
	}()
	t.Run("Error on construction", func(t *testing.T) {

		_, err := NewHexDfa(1234)
		assert.Error(t, err)
	})
}
