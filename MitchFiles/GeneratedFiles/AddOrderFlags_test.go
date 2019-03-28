package GeneratedFiles

import (
	"bytes"
	"github.com/bhbosman/Application/Streams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAddOrderFlags(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		msg := NewAddOrderFlags()
		assert.NotNil(t, msg)
	})

	t.Run("WriteMessage", func(t *testing.T) {
		msg := NewAddOrderFlags()
		if !assert.NotNil(t, msg) {
			return
		}
		buffer := bytes.NewBuffer(nil)
		mitchWriter := Streams.NewMitchWriter(buffer)
		if !assert.NotNil(t, mitchWriter) {
			return
		}
		n, err := Write_AddOrderFlags(mitchWriter, msg)
		assert.NoError(t, err)
		assert.Equal(t, 1, n)

	})

}
