package Streams

type IReaderUint16 interface {
	Read_uint16() (value uint16, n int, err error)
}
