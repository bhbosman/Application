package Streams

type IStreamReader interface {
	Read_MitchUInt16() (MitchUInt16, int, error)
	Read_MitchByte() (MitchByte, int, error)
	Read_MitchUInt32() (MitchUInt32, int, error)

}
