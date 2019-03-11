package Streams

type IStreamWriter interface {
	Write_MitchUInt16(value MitchUInt16) (int, error)
	Write_MitchByte(value MitchByte) (int, error)
	Write_MitchUInt32(value MitchUInt32) (int, error)
}
