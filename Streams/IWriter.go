package Streams

type IWriter interface {
	Write_byte(i byte) (int, error)
	Write_uint16(u uint16) (int, error)
	Write_uint32(i uint32) (int, error)
	Write_uint64(u uint64) (int, error)
	Write_string(s string, i int) (int, error)
}
