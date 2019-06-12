package Streams

type IWriterByte interface {
	Write_byte(i byte) (int, error)
}
