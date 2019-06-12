package Streams

type IWriterUint16 interface {
	Write_uint16(u uint16) (int, error)
}
