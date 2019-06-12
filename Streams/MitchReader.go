package Streams

import (
	"encoding/binary"
	"io"
	"math"
	"strings"
	"time"
)

type MitchReader struct {
	Reader io.Reader
}

func (self *MitchReader) Read_ReadBytes(size int) (value []byte, n int, err error) {
	result := make([]byte, size)
	n, err = self.Reader.Read(result)
	if err != nil {
		return nil, 0, err
	}
	return result, n, nil
}

func NewMitchReader(reader io.Reader) *MitchReader {
	return &MitchReader{Reader: reader}
}

func (self *MitchReader) bufferSize(length int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(length)))))
}

//noinspection ALL
func (self *MitchReader) Read_string(length int) (value string, n int, err error) {
	bs := make([]byte, length)
	n, err = io.ReadAtLeast(self.Reader, bs, length)
	if err != nil {
		return "", 0, err
	}

	return string(bs), n, nil
}

//noinspection ALL
func (self *MitchReader) Read_byte() (value byte, n int, err error) {
	bs := []byte{0}
	n, err = io.ReadFull(self.Reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = bs[0]
	return value, n, nil
}

//noinspection ALL
func (self *MitchReader) Read_uint16() (value uint16, n int, err error) {
	bs := []byte{0, 0}
	n, err = io.ReadFull(self.Reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = binary.LittleEndian.Uint16(bs)
	return value, n, nil
}

//noinspection ALL
func (self *MitchReader) Read_uint32() (value uint32, n int, err error) {
	bs := []byte{0, 0, 0, 0}
	n, err = io.ReadFull(self.Reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = binary.LittleEndian.Uint32(bs)
	return value, n, nil
}

//noinspection ALL
func (self *MitchReader) Read_uint64() (value uint64, n int, err error) {
	bs := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	n, err = io.ReadFull(self.Reader, bs)
	if err != nil {
		return 0, 0, err
	}
	value = binary.LittleEndian.Uint64(bs)
	return value, n, nil
}

//noinspection ALL
func (self *MitchReader) Read_mitch_date() (value time.Time, n int, err error) {
	var s string
	s, n, err = self.Read_string(8)
	if err != nil {
		return time.Time{}, 0, err
	}
	value, err = time.Parse("20060102", s)
	if err != nil {
		if _, ok := err.(*time.ParseError); ok {
			s := strings.Trim(s, " ")
			if len(s) == 0 {
				return time.Time{}, n, nil
			}
		}
		return time.Time{}, 0, err
	}
	return value, n, nil
}

//noinspection ALL
func (self *MitchReader) Read_mitch_time() (value time.Time, n int, err error) {
	var s string
	s, n, err = self.Read_string(8)
	if err != nil {
		return time.Time{}, 0, err
	}
	value, err = time.Parse("15:04:05", s)
	if err != nil {
		if _, ok := err.(*time.ParseError); ok {
			s := strings.Trim(s, " ")
			if len(s) == 0 {
				return time.Time{}, n, nil
			}
		}
		return time.Time{}, 0, err
	}
	return value, n, nil
}

//noinspection ALL
func (self *MitchReader) Read_mitch_price04() (float64, int, error) {
	v, n, e := self.Read_uint64()
	if e != nil {
		return 0, 0, e
	}
	return float64(v / 10000), n, e
}

//noinspection ALL
func (self *MitchReader) Read_mitch_price08() (value float64, n int, err error) {
	v, n, e := self.Read_uint64()
	if e != nil {
		return 0, 0, e
	}
	return float64(v / 100000000), n, e
}
