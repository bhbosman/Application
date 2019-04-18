package Streams

type IReaderByte interface {
	Read_byte() (value byte, n int, err error)
}


type IReaderUint16 interface {
	Read_uint16() (value uint16, n int, err error)
}
type IReaderUint32 interface {
	Read_uint32() (value uint32, n int, err error)
}
type IReaderUint64 interface {
	Read_uint64() (value uint64, n int, err error)
}
type IReaderString interface {
	Read_string(length int) (value string, n int, err error)
}

