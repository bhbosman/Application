package Streams

import "time"

type IMitchReader interface {
	IReaderByte
	IReaderString
	IReaderUint16
	IReaderUint32
	IReaderUint64
	IReaderReadBytes
	Read_mitch_date() (value time.Time, n int, err error)
	Read_mitch_time() (value time.Time, n int, err error)
	Read_mitch_price08() (value float64, n int, err error)
	Read_mitch_price04() (float64, int, error)
}
