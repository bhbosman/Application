package GeneratedFiles

import (
	"bytes"
	"encoding/binary"
	"github.com/bhbosman/Application/Streams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAuctionTradeMessage(t *testing.T) {
	t.Run("Constructor", func(t *testing.T) {
		msg := NewAuctionTradeMessage()
		assert.NotNil(t, msg)
	})

	t.Run("WriteMessage", func(t *testing.T) {
		msg := NewAuctionTradeMessage()
		if !assert.NotNil(t, msg) {
			return
		}
		buffer := bytes.NewBuffer(nil)
		mitchWriter := Streams.NewMitchWriter(buffer)
		if !assert.NotNil(t, mitchWriter) {
			return
		}
		n, err := Write_AuctionTradeMessage(mitchWriter, msg)
		assert.NoError(t, err)
		assert.Equal(t, 58, n)

		streamData := buffer.Bytes()
		reader := bytes.NewReader(streamData)

		var length uint16
		err = binary.Read(reader, binary.LittleEndian, &length)
		assert.NoError(t, err)
		assert.Equal(t, uint16(58), length)

		var messageByte byte
		err = binary.Read(reader, binary.LittleEndian, &messageByte)
		assert.NoError(t, err)
		assert.Equal(t, byte(0x51), messageByte)
	})

}
