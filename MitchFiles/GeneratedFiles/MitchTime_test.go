package GeneratedFiles

import (
	"bytes"
	"github.com/bhbosman/Application/Streams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeMessage(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		msg, _ := TimeMessageFactory.New()
		assert.NotNil(t, msg)
	})

	t.Run("WriteMessage", func(t *testing.T) {
		msg, _ := TimeMessageFactory.New()
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
		v, _ := TimeMessageFactory.New()
		n, err := TimeMessageFactory.ReadMessageInFull(v, mitchReader)
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
		v, err := TimeMessageFactory.New()
		assert.NoError(t, err)
		if !assert.NotNil(t, v) {
			return
		}

		n, err := TimeMessageFactory.ReadMessageInFull(v, mitchReader)
		if !assert.Error(t, err) {
			return
		}
		if !assert.Equal(t, 0, n) {
			return
		}
	})
}
