package Streams

//noinspection ALL
type IMitchWriter interface {
	Write_uint32(value uint32) (n int, err error)
	Write_MitchUInt16(value MitchUInt16) (int, error)
	Write_MitchByte(value MitchByte) (int, error)
	Write_MitchUInt32(uInt32 MitchUInt32) (int, error)
	Write_MitchAlpha(alpha MitchAlpha) (int, error)
	Write_MitchPrice08(price08 MitchPrice08) (int, error)
	Write_MitchDate(date MitchDate) (int, error)
	Write_MitchTime(time MitchTime) (int, error)
	Write_MitchBitField(field MitchBitField) (int, error)
	Write_MitchUInt64(uInt64 MitchUInt64) (int, error)
}


