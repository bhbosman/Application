package Streams

type IReaderString interface {
	Read_string(length int) (value string, n int, err error)
}
