package Streams

type IWriterUint64 interface {
	Write_uint64(u uint64) (int, error)
}
