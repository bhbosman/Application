package Streams

import "time"

type IMitchWriter interface {
	IWriter
	Write_mitch_time(time time.Time) (int, error)
	Write_mitch_date(date time.Time) (int, error)
	Write_mitch_price04(f float64) (int, error)
	Write_mitch_price08(f float64) (int, error)
}
