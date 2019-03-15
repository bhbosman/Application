package DFA

import (
	"errors"
	"github.com/bhbosman/Application/Common"
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
	Common.ErrorListFactory.Replace(func() Common.IErrorList {
		return &errorListWhichWillReturnError{}
	})
	defer func() {
		Common.ErrorListFactory.Reset()
	}()
	t.Run("Error on construction", func(t *testing.T) {

		_, err := NewHexDfa(1234)
		assert.Error(t, err)
	})
}
