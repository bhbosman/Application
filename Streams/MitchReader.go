package Streams

import (
	"encoding/binary"
	"io"
	"math"
	"sync"
	"time"
)

type MitchReader struct {
	reader io.Reader
}

func NewMitchReader(reader io.Reader) *MitchReader {
	return &MitchReader{reader: reader}
}

func (self *MitchReader) bufferSize(length int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(length)))))
}

func (self *MitchReader) Read_string(length int) (value string, n int, err error) {
	bufferLength := self.bufferSize(length)
	pool, ok := byteArrayPools[bufferLength]
	var bs []byte
	if ok {
		bs = pool.Get().([]byte)
		defer pool.Put(bs)
	} else {
		bs = make([]byte, length)
	}
	n, err = io.ReadAtLeast(self.reader, bs, length)
	if err != nil {
		return "", 0, err
	}
	if ok {
		value = string(bs[0:length])
	} else {
		value = string(bs)
	}
	return value, n, nil
}

func (self *MitchReader) Read_byte() (value byte, n int, err error) {
	pool := byteArrayPools[1]
	bs := pool.Get().([]byte)
	defer pool.Put(bs)

	n, err = io.ReadFull(self.reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = bs[0]
	return value, n, nil
}

func (self *MitchReader) Read_uint16() (value uint16, n int, err error) {
	pool := byteArrayPools[2]
	bs := pool.Get().([]byte)
	defer pool.Put(bs)

	n, err = io.ReadFull(self.reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = binary.LittleEndian.Uint16(bs)
	return value, n, nil
}

func (self *MitchReader) Read_uint32() (value uint32, n int, err error) {
	pool := byteArrayPools[4]
	bs := pool.Get().([]byte)
	defer pool.Put(bs)

	n, err = io.ReadFull(self.reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = binary.LittleEndian.Uint32(bs)
	return value, n, nil
}

func (self *MitchReader) Read_uint64() (value uint64, n int, err error) {
	pool := byteArrayPools[8]
	bs := pool.Get().([]byte)
	defer pool.Put(bs)

	n, err = io.ReadFull(self.reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = binary.LittleEndian.Uint64(bs)
	return value, n, nil
}

func (self *MitchReader) Read_mitch_date() (value time.Time, n int, err error) {
	var s string
	s, n, err = self.Read_string(8)
	if err != nil {
		return time.Time{}, 0, err
	}
	value, err = time.Parse("20060102", s)
	if err != nil {
		return time.Time{}, 0, err
	}
	return value, n, nil
}

func (self *MitchReader) Read_mitch_time() (value time.Time, n int, err error) {
	var s string
	s, n, err = self.Read_string(8)
	if err != nil {
		return time.Time{}, 0, err
	}
	value, err = time.Parse("15:04:05", s)
	if err != nil {
		return time.Time{}, 0, err
	}
	return value, n, nil
}

func (self *MitchReader) Read_mitch_price04() (float64, int, error) {
	v, n, e := self.Read_uint64()
	if e != nil {
		return 0, 0, e
	}
	return float64(v / 10000), n, e
}

func (self *MitchReader) Read_mitch_price08() (value float64, n int, err error) {
	v, n, e := self.Read_uint64()
	if e != nil {
		return 0, 0, e
	}
	return float64(v / 100000000), n, e
}

var byteArrayPools map[int]*sync.Pool

func init() {
	byteArrayPools = make(map[int]*sync.Pool)
	byteArrayPools[1] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1)
		},
	}
	byteArrayPools[2] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 2)
		},
	}
	byteArrayPools[4] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 4)
		},
	}
	byteArrayPools[8] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 8)
		},
	}
	byteArrayPools[16] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 16)
		},
	}
	byteArrayPools[32] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 32)
		},
	}
	byteArrayPools[64] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 64)
		},
	}
}
