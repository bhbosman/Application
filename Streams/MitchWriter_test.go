package Streams

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type failBuffer struct {
}

func (self *failBuffer) Write(p []byte) (n int, err error) {
	return 0, errors.New("failed")
}

func TestMitchWriter(t *testing.T) {
	t.Run("Check interface", func(t *testing.T) {
		var obj interface{}
		obj = &MitchWriter{}
		_, ok := obj.(IMitchWriter)
		assert.True(t, ok)
	})
	t.Run("Write String Method", func(t *testing.T) {
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
			assert.Equal(t, []byte{'A', 'B', 'C', 32, 32, 32}, dataInBuffer)
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
	})
	t.Run("Write Byte Method", func(t *testing.T) {
		t.Run("Write one byte on valid buffer", func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_byte('d')
			assert.NoError(t, e)
			assert.Equal(t, 1, n)

			bytesInBuffer := buffer.Bytes()
			assert.Len(t, bytesInBuffer, 1)
			assert.Equal(t, byte('d'), bytesInBuffer[0])
		})

		t.Run("Write one byte on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_byte('d')
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})
	t.Run("Write uint16", func(t *testing.T) {
		t.Run("Write one uint16 on valid buffer. Test number 'd'", func(t *testing.T) {
			value := uint16('d')
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint16(value)
			assert.NoError(t, e)
			assert.Equal(t, 2, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0}
			binary.LittleEndian.PutUint16(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])

		})

		t.Run("Write one uint16 on valid buffer. Test number '0xab'", func(t *testing.T) {
			value := uint16(0xabcd)
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint16(value)
			assert.NoError(t, e)
			assert.Equal(t, 2, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0}
			binary.LittleEndian.PutUint16(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
		})

		t.Run("Write one uint16 on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint16('d')
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})
	t.Run("Write uint32", func(t *testing.T) {
		t.Run("Write one uint32 on valid buffer. Test number 'd'", func(t *testing.T) {
			value := uint32('d')
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint32(value)
			assert.NoError(t, e)
			assert.Equal(t, 4, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0}
			binary.LittleEndian.PutUint32(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
		})

		t.Run("Write one uint32 on valid buffer. Test number '0xabcd'", func(t *testing.T) {
			value := uint32(0xabcd)
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint32(value)
			assert.NoError(t, e)
			assert.Equal(t, 4, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0}
			binary.LittleEndian.PutUint32(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
		})

		t.Run("Write one uint32 on valid buffer. Test number '0xabcdef01'", func(t *testing.T) {
			value := uint32(0xabcdef01)
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint32(value)
			assert.NoError(t, e)
			assert.Equal(t, 4, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0}
			binary.LittleEndian.PutUint32(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
		})

		t.Run("Write one uint32 on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint32('d')
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})
	t.Run("Write uint64", func(t *testing.T) {
		t.Run("Write one uint64 on valid buffer. Test number 'd'", func(t *testing.T) {
			value := uint64('d')
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint64(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0, 0, 0, 0, 0}
			binary.LittleEndian.PutUint64(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
			assert.Equal(t, buffForLittleEndian[4], bytesInBuffer[4])
			assert.Equal(t, buffForLittleEndian[5], bytesInBuffer[5])
			assert.Equal(t, buffForLittleEndian[6], bytesInBuffer[6])
			assert.Equal(t, buffForLittleEndian[7], bytesInBuffer[7])
		})

		t.Run("Write one uint64 on valid buffer. Test number '0xabcd'", func(t *testing.T) {
			value := uint64(0xabcd)
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint64(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0, 0, 0, 0, 0}
			binary.LittleEndian.PutUint64(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
			assert.Equal(t, buffForLittleEndian[4], bytesInBuffer[4])
			assert.Equal(t, buffForLittleEndian[5], bytesInBuffer[5])
			assert.Equal(t, buffForLittleEndian[6], bytesInBuffer[6])
			assert.Equal(t, buffForLittleEndian[7], bytesInBuffer[7])
		})

		t.Run("Write one uint64 on valid buffer. Test number '0xabcdef01'", func(t *testing.T) {
			value := uint64(0xabcdef01)
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint64(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0, 0, 0, 0, 0}
			binary.LittleEndian.PutUint64(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
			assert.Equal(t, buffForLittleEndian[4], bytesInBuffer[4])
			assert.Equal(t, buffForLittleEndian[5], bytesInBuffer[5])
			assert.Equal(t, buffForLittleEndian[6], bytesInBuffer[6])
			assert.Equal(t, buffForLittleEndian[7], bytesInBuffer[7])
		})

		t.Run("Write one uint64 on valid buffer. Test number '0xabcdef01234567'", func(t *testing.T) {
			value := uint64(0xabcdef01234567)
			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint64(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0, 0, 0, 0, 0}
			binary.LittleEndian.PutUint64(buffForLittleEndian, value)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
			assert.Equal(t, buffForLittleEndian[4], bytesInBuffer[4])
			assert.Equal(t, buffForLittleEndian[5], bytesInBuffer[5])
			assert.Equal(t, buffForLittleEndian[6], bytesInBuffer[6])
			assert.Equal(t, buffForLittleEndian[7], bytesInBuffer[7])
		})

		t.Run("Write one uint64 on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint32('d')
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})
	t.Run("Write time", func(t *testing.T) {
		t.Run("Write one time on valid buffer.", func(t *testing.T) {
			value, err := time.Parse("15:04:05", "01:02:03")
			assert.NoError(t, err)

			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_time(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			dataInBuffer := buffer.String()
			assert.Equal(t, "01:02:03", dataInBuffer)
		})

		t.Run("Write one time on valid buffer.", func(t *testing.T) {
			value, err := time.Parse("15:04:05", "23:34:01")
			assert.NoError(t, err)

			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_time(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			dataInBuffer := buffer.String()
			assert.Equal(t, "23:34:01", dataInBuffer)
		})

		t.Run("Write one time on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_uint32('d')
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})
	t.Run("Write date", func(t *testing.T) {
		t.Run("Write one date on valid buffer.", func(t *testing.T) {
			value, err := time.Parse("20060102", "20190101")
			assert.NoError(t, err)

			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_date(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			dataInBuffer := buffer.String()
			assert.Equal(t, "20190101", dataInBuffer)
		})

		t.Run("Write one date on valid buffer.", func(t *testing.T) {
			value, err := time.Parse("20060102", "20190401")
			assert.NoError(t, err)

			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_date(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			dataInBuffer := buffer.String()
			assert.Equal(t, "20190401", dataInBuffer)
		})

		t.Run("Write date on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_date(time.Time{})
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})

	t.Run("Write price04", func(t *testing.T) {
		t.Run("Write price04 on valid buffer.", func(t *testing.T) {
			value := 1234.5678
			streamValue := uint64(value * 10000)

			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_price04(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0, 0, 0, 0, 0}
			binary.LittleEndian.PutUint64(buffForLittleEndian, streamValue)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
			assert.Equal(t, buffForLittleEndian[4], bytesInBuffer[4])
			assert.Equal(t, buffForLittleEndian[5], bytesInBuffer[5])
			assert.Equal(t, buffForLittleEndian[6], bytesInBuffer[6])
			assert.Equal(t, buffForLittleEndian[7], bytesInBuffer[7])
		})

		t.Run("Write price04 on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_price04(1234.5678)
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})

	t.Run("Write price08", func(t *testing.T) {
		t.Run("Write price08 on valid buffer.", func(t *testing.T) {
			value := 1234.56789012
			streamValue := uint64(value * 100000000)

			buffer := bytes.NewBuffer(nil)
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_price08(value)
			assert.NoError(t, e)
			assert.Equal(t, 8, n)

			bytesInBuffer := buffer.Bytes()
			buffForLittleEndian := []byte{0, 0, 0, 0, 0, 0, 0, 0}
			binary.LittleEndian.PutUint64(buffForLittleEndian, streamValue)
			assert.Equal(t, buffForLittleEndian[0], bytesInBuffer[0])
			assert.Equal(t, buffForLittleEndian[1], bytesInBuffer[1])
			assert.Equal(t, buffForLittleEndian[2], bytesInBuffer[2])
			assert.Equal(t, buffForLittleEndian[3], bytesInBuffer[3])
			assert.Equal(t, buffForLittleEndian[4], bytesInBuffer[4])
			assert.Equal(t, buffForLittleEndian[5], bytesInBuffer[5])
			assert.Equal(t, buffForLittleEndian[6], bytesInBuffer[6])
			assert.Equal(t, buffForLittleEndian[7], bytesInBuffer[7])
		})

		t.Run("Write price08 on invalid buffer", func(t *testing.T) {
			buffer := &failBuffer{}
			writer := &MitchWriter{
				writer: buffer,
			}
			n, e := writer.Write_mitch_price04(1234.5678)
			assert.Error(t, e)
			assert.Equal(t, 0, n)
		})
	})

}
