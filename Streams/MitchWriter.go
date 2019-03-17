package Streams

import (
	"encoding/binary"
	"io"
	"time"
)

type MitchWriter struct {
	writer io.Writer
}

func NewMitchWriter(writer io.Writer) *MitchWriter {
	return &MitchWriter{
		writer: writer,
	}
}

func (self *MitchWriter) Write_string(s string, length int) (int, error) {
	if len(s) >= length {
		return self.writer.Write([]byte(s[0:length]))
	}
	bytesWritten := 0
	n, err := self.writer.Write([]byte(s))
	if err != nil {
		return n, err
	}
	bytesWritten += n
	for bytesWritten < length {
		n, err := self.Write_byte(' ')
		if err != nil {
			return 0, err
		}
		bytesWritten += n
	}
	return bytesWritten, nil
}

func (self *MitchWriter) Write_byte(i byte) (int, error) {
	return self.writer.Write([]byte{i})
}

func (self *MitchWriter) Write_uint16(u uint16) (int, error) {
	err := binary.Write(self.writer, binary.LittleEndian, &u)
	if err != nil {
		return 0, err
	}
	return 2, nil
}

func (self *MitchWriter) Write_uint32(i uint32) (int, error) {
	err := binary.Write(self.writer, binary.LittleEndian, &i)
	if err != nil {
		return 0, err
	}
	return 4, nil
}

func (self *MitchWriter) Write_uint64(u uint64) (int, error) {
	err := binary.Write(self.writer, binary.LittleEndian, &u)
	if err != nil {
		return 0, err
	}
	return 8, nil
}

func (self *MitchWriter) Write_mitch_time(time time.Time) (int, error) {
	return self.Write_string(time.Format("15:04:05"), 8)
}

func (self *MitchWriter) Write_mitch_date(date time.Time) (int, error) {
	return self.Write_string(date.Format("20060102"), 8)
}

func (self *MitchWriter) Write_mitch_price04(f float64) (int, error) {
	return self.Write_uint64(uint64(f * 10000))
}

func (self *MitchWriter) Write_mitch_price08(f float64) (int, error) {
	return self.Write_uint64(uint64(f * 100000000))
}

//
