package Streams

type IStreamReader interface {
	ReadBool() (bool, error)
	ReadInt8() (int8, error)
	ReadInt16() (int16, error)
	ReadInt32() (int32, error)
	ReadInt64() (int64, error)
	ReadUint8() (uint8, error)
	ReadUint16() (int16, error)
	ReadUint32() (uint32, error)
	ReadUint64() (int64, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
	ReadComplex64() (complex64, error)
	ReadComplex128() (complex128, error)
	ReadString() (string, error)
}
