package Streams

type IReaderUint32 interface {
	Read_uint32() (value uint32, n int, err error)
}

type IReaderReadBytes interface {
	Read_ReadBytes(buffer []byte, size int) (value []byte, n int, err error)
}
