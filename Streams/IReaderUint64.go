package Streams

type IReaderUint64 interface {
	Read_uint64() (value uint64, n int, err error)
}
