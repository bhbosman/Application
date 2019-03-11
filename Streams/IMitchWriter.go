package Streams

type IMitchWriter interface {
	Write_uint32(value uint32) (n int, err error)
	Write_MitchUInt16(value MitchUInt16) (int, error)
	Write_MitchByte(value MitchByte) (int, error)
	Write_MitchUInt32(uInt32 MitchUInt32) (int, error)
	Write_MitchAlpha(alpha MitchAlpha) (int, error)
}
type IMitchReader interface {
	Read_uint32() (value uint32, n int, err error)
}
