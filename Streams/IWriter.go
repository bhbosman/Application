package Streams

type IWriterByte interface {
	Write_byte(i byte) (int, error)
}

type IWriterUint16 interface {
	Write_uint16(u uint16) (int, error)
}
type IWriterUint32 interface {
	Write_uint32(u uint32) (int, error)
}
type IWriterUint64 interface {
	Write_uint64(u uint64) (int, error)
}

type IWriterString interface {
	Write_string(s string, i int) (int, error)
}

