package Streams

type IStreamWriter interface {
	WriteBool(value bool) (int, error)
	WriteInt8(value int8) (int, error)
	WriteInt16(value int16) (int, error)
	WriteInt32(value int32) (int, error)
	WriteInt64(value int64) (int, error)
	WriteUint8(value uint8) (int, error)
	WriteUint16(value uint16) (int, error)
	WriteUint32(value uint32) (int, error)
	WriteUint64(value uint64) (int, error)
	WriteFloat32(value float32) (int, error)
	WriteFloat64(value float64) (int, error)
	WriteComplex64(value complex64) (int, error)
	WriteComplex128(value complex128) (int, error)
	WriteString(value string) (int, error)
}
