package Streams

import "time"

type IMitchReader interface {
	Read_string(length int) (value string, n int, err error)
	Read_byte() (value byte, n int, err error)
	Read_uint16() (value uint16, n int, err error)
	Read_uint32() (value uint32, n int, err error)
	Read_uint64() (value uint64, n int, err error)
	Read_mitch_date() (value time.Time, n int, err error)
	Read_mitch_time() (value time.Time, n int, err error)
	Read_mitch_price08() (value float64, n int, err error)
	Read_mitch_price04() (float64, int, error)
}
