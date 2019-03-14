package Streams

import "time"

//noinspection ALL
type IMitchWriter interface {
	Write_string(s string, i int) (int, error)
	Write_byte(i byte) (int, error)
	Write_uint16(u uint16) (int, error)
	Write_uint32(i uint32) (int, error)
	Write_uint64(u uint64) (int, error)
	Write_mitch_time(time time.Time) (int, error)
	Write_mitch_date(date time.Time) (int, error)
	Write_mitch_price04(f float64) (int, error)
	Write_mitch_price08(f float64) (int, error)
}
