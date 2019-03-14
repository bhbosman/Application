package Streams

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)


type failBuffer struct{
}

func (self *failBuffer) Write(p []byte) (n int, err error) {
	return 0, errors.New("failed")
}

func TestMitchWriter_Write_String(t *testing.T) {

	t.Run("Write String", func(t *testing.T) {
		buffer := bytes.NewBuffer(nil)
		writer := &MitchWriter{
			writer: buffer,
		}
		dataWritten := "ABCDEF"
		dataLength := 6

		n, e := writer.Write_string(dataWritten, dataLength)
		assert.NoError(t, e)
		assert.Equal(t, dataLength, n)

		dataInBuffer := buffer.String()
		assert.Equal(t, dataWritten, dataInBuffer)
	})

	t.Run("Write String", func(t *testing.T) {
		buffer := bytes.NewBuffer(nil)
		writer := &MitchWriter{
			writer: buffer,
		}
		dataWritten := "ABCDEFGHI"
		dataLength := 6

		n, e := writer.Write_string(dataWritten, dataLength)
		assert.NoError(t, e)
		assert.Equal(t, dataLength, n)

		dataInBuffer := buffer.String()
		assert.Equal(t, dataWritten[0:6], dataInBuffer)
	})
	t.Run("Write String", func(t *testing.T) {
		buffer := bytes.NewBuffer(nil)
		writer := &MitchWriter{
			writer: buffer,
		}
		dataWritten := "ABC"
		dataLength := 6

		n, e := writer.Write_string(dataWritten, dataLength)
		assert.NoError(t, e)
		assert.Equal(t, dataLength, n)

		dataInBuffer := buffer.Bytes()
		assert.Equal(t, []byte{'A', 'B', 'C', 0, 0, 0}, dataInBuffer)
	})

	t.Run("Write String", func(t *testing.T) {

		writer := &MitchWriter{
			writer: &failBuffer{},
		}
		dataWritten := "ABC"
		dataLength := 6

		_, e := writer.Write_string(dataWritten, dataLength)
		assert.Error(t, e)

	})



}
