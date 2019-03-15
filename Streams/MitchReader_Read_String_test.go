package Streams

import (
	"bytes"
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMitchReader_Read_string(t *testing.T) {
	t.Run("Test Read Mitch Alpha", func(t *testing.T) {
		reader := &MitchReader{
			reader: bytes.NewReader([]byte("20190202")),
		}
		s, n, e := reader.Read_string(8)
		assert.NoError(t, e)
		assert.Equal(t, 8, n)
		assert.Equal(t, "20190202", s)
	})

	t.Run("Test Read Mitch Alpha", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_string(8)
		assert.Error(t, e)
	})

}

func TestMitchReader_Read_byte(t *testing.T) {
	t.Run("Test Read Mitch Byte", func(t *testing.T) {
		initValue := byte(0x12)
		b := make([]byte, 1)
		b[0] = initValue
		reader := &MitchReader{
			reader: bytes.NewReader(b),
		}
		value, n, e := reader.Read_byte()
		assert.NoError(t, e)
		assert.Equal(t, 1, n)
		assert.Equal(t, initValue, value)
	})
	t.Run("Test Read Mitch Byte", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_byte()
		assert.Error(t, e)
	})
}

func TestMitchReader_Read_mitch_date(t *testing.T) {
	t.Run("Test Read Mitch Date", func(t *testing.T) {
		reader := &MitchReader{
			reader: bytes.NewReader([]byte("20190202")),
		}
		d, n, e := reader.Read_mitch_date()
		assert.NoError(t, e)
		assert.Equal(t, 8, n)
		assert.Equal(t, 2019, d.Year())
		assert.Equal(t, time.Month(2), d.Month())
		assert.Equal(t, 2, d.Day())
	})

	t.Run("Test Read Mitch Date Invalid format", func(t *testing.T) {
		reader := &MitchReader{
			reader: bytes.NewReader([]byte("21212019")),
		}
		_, _, e := reader.Read_mitch_date()
		assert.Error(t, e)
	})
	t.Run("Test Read Mitch Date", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_mitch_date()
		assert.Error(t, e)
	})

}

func TestMitchReader_Read_mitch_price04(t *testing.T) {
	t.Run("Test Read Mitch Price04", func(t *testing.T) {
		intValue := uint64(12345678)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, intValue)
		reader := &MitchReader{
			reader: bytes.NewReader(b),
		}
		value, n, e := reader.Read_mitch_price04()
		assert.NoError(t, e)
		assert.Equal(t, 8, n)
		assert.Equal(t, float64(intValue/10000), value)
	})
	t.Run("Test Read Mitch Price04", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_mitch_price04()
		assert.Error(t, e)
	})
}

func TestMitchReader_Read_mitch_price08(t *testing.T) {
	t.Run("Test Read Mitch Price08", func(t *testing.T) {
		intValue := uint64(1234567890123)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, intValue)
		reader := &MitchReader{
			reader: bytes.NewReader(b),
		}
		value, n, e := reader.Read_mitch_price08()
		assert.NoError(t, e)
		assert.Equal(t, 8, n)
		assert.Equal(t, float64(intValue/100000000), value)
	})
	t.Run("Test Read Mitch Price08", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_mitch_price08()
		assert.Error(t, e)
	})
}

func TestMitchReader_Read_mitch_time(t *testing.T) {
	t.Run("Test Read Mitch Time", func(t *testing.T) {
		reader := &MitchReader{
			reader: bytes.NewReader([]byte("15:02:04")),
		}
		time, n, e := reader.Read_mitch_time()
		assert.NoError(t, e)
		assert.Equal(t, 8, n)
		assert.Equal(t, 15, time.Hour())
		assert.Equal(t, 2, time.Minute())
		assert.Equal(t, 4, time.Second())
	})

	t.Run("Test Read Mitch Time Invalid format", func(t *testing.T) {
		reader := &MitchReader{
			reader: bytes.NewReader([]byte("15302304")),
		}
		_, _, e := reader.Read_mitch_time()
		assert.Error(t, e)
	})

	t.Run("Test Read Mitch Time", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_mitch_time()
		assert.Error(t, e)
	})

}

func TestMitchReader_Read_uint16(t *testing.T) {
	t.Run("Test Read Mitch Uint16", func(t *testing.T) {
		initValue := uint16(0x1234)
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, initValue)
		reader := &MitchReader{
			reader: bytes.NewReader(b),
		}
		value, n, e := reader.Read_uint16()
		assert.NoError(t, e)
		assert.Equal(t, 2, n)
		assert.Equal(t, initValue, value)
	})
	t.Run("Test Read Mitch Uint16", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_uint16()
		assert.Error(t, e)
	})
}

func TestMitchReader_Read_uint32(t *testing.T) {
	t.Run("Test Read Mitch Uint32", func(t *testing.T) {
		initValue := uint32(0x12345678)
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, initValue)
		reader := &MitchReader{
			reader: bytes.NewReader(b),
		}
		value, n, e := reader.Read_uint32()
		assert.NoError(t, e)
		assert.Equal(t, 4, n)
		assert.Equal(t, initValue, value)
	})
	t.Run("Test Read Mitch Uint32", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_uint32()
		assert.Error(t, e)
	})
}

func TestMitchReader_Read_uint64(t *testing.T) {
	t.Run("Test Read Mitch Uint64", func(t *testing.T) {
		initValue := uint64(0x1234567890ABCDEF)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, initValue)
		reader := &MitchReader{
			reader: bytes.NewReader(b),
		}
		value, n, e := reader.Read_uint64()
		assert.NoError(t, e)
		assert.Equal(t, 8, n)
		assert.Equal(t, initValue, value)
	})
	t.Run("Test Read Mitch Uint64", func(t *testing.T) {
		reader := &MitchReader{reader: bytes.NewReader([]byte(""))}
		_, _, e := reader.Read_uint64()
		assert.Error(t, e)
	})
}