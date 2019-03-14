package Streams

import (
	"io"
)

type MitchWriter struct {
	writer io.Writer
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
	n, err = self.writer.Write(make([]byte, length-n))
	bytesWritten += n
	return bytesWritten, nil
}

//func (self *MitchWriter) Write_byte(i byte) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_uint16(u uint16) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_uint32(i uint32) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_uint64(u uint64) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_mitch_time(time time.Time) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_mitch_date(date time.Time) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_mitch_price04(f float64) (int, error) {
//	panic("implement me")
//}
//
//func (self *MitchWriter) Write_mitch_price08(f float64) (int, error) {
//	panic("implement me")
//}
//
