package Streams

import (
	"encoding/binary"
	"io"
	"time"
)

type MitchReader struct {
	reader io.Reader
}

func (self *MitchReader) Read_string(length int) (value string, n int, err error) {
	byteArray := make([]byte, length)
	err = binary.Read(self.reader, binary.LittleEndian, byteArray)
	if err != nil {
		return "", 0, err
	}
	value = string(byteArray)
	return value, length, err
}

func (self *MitchReader) Read_byte() (value byte, n int, err error) {
	err = binary.Read(self.reader, binary.LittleEndian, &value)
	if err != nil {
		return 0, 0, err
	}
	return value, 1, nil
}

func (self *MitchReader) Read_uint16() (value uint16, n int, err error) {
	err = binary.Read(self.reader, binary.LittleEndian, &value)
	if err != nil {
		return 0, 0, err
	}
	return value, 2, nil
}

func (self *MitchReader) Read_uint32() (value uint32, n int, err error) {
	err = binary.Read(self.reader, binary.LittleEndian, &value)
	if err != nil {
		return 0, 0, err
	}
	return value, 4, nil
}

func (self *MitchReader) Read_uint64() (value uint64, n int, err error) {
	err = binary.Read(self.reader, binary.LittleEndian, &value)
	if err != nil {
		return 0, 0, err
	}
	return value, 8, nil
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
