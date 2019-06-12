package Streams

type IReaderByte interface {
	Read_byte() (value byte, n int, err error)
}
