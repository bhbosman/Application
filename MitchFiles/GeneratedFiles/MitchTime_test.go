package GeneratedFiles

import (
	"bytes"
	"github.com/bhbosman/Application/Streams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeMessage(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		msg := NewTimeMessage()
		assert.NotNil(t, msg)
	})

	t.Run("WriteMessage", func(t *testing.T) {
		msg := NewTimeMessage()
		if !assert.NotNil(t, msg) {
			return
		}
		buffer := bytes.NewBuffer(nil)
		mitchWriter := Streams.NewMitchWriter(buffer)
		if !assert.NotNil(t, mitchWriter) {
			return
		}
		n, err := Write_TimeMessage(mitchWriter, msg)
		assert.NoError(t, err)
		assert.Equal(t, 7, n)

		streamData := buffer.Bytes()
		assert.Equal(t, byte(0x54), streamData[2])
	})

	t.Run("ReadMessage with correct message code", func(t *testing.T) {
		buffer := bytes.NewBuffer([]byte{0, 0, 84, 0, 0, 0, 0})
		mitchReader := Streams.NewMitchReader(buffer)
		if !assert.NotNil(t, mitchReader) {
			return
		}
		v, n, err := Read_TimeMessage(mitchReader)
		assert.NoError(t, err)
		assert.Equal(t, 7, n)
		assert.NotNil(t, v)
	})

	t.Run("ReadMessage with incorrect message code", func(t *testing.T) {
		buffer := bytes.NewBuffer([]byte{0, 0, 0, 0, 0, 0, 0})
		mitchReader := Streams.NewMitchReader(buffer)
		if !assert.NotNil(t, mitchReader) {
			return
		}
		v, n, err := Read_TimeMessage(mitchReader)
		if !assert.Error(t, err) {
			return
		}
		if !assert.Equal(t, 0, n) {
			return
		}
		if !assert.Nil(t, v) {
			return
		}
	})
}
