package Streams

type IWriterUint32 interface {
	Write_uint32(u uint32) (int, error)
}
