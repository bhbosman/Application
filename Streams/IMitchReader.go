package Streams


//noinspection ALL
type IMitchReader interface {
	Read_uint32() (value uint32, n int, err error)
	Read_MitchUInt16() (MitchUInt16, int, error)
	Read_MitchByte() (MitchByte, int, error)
	Read_MitchUInt32() (MitchUInt32, int, error)
	Read_MitchDate() (MitchDate, int, error)
	Read_MitchPrice08() (MitchPrice08, int, error)
	Read_MitchAlpha() (MitchAlpha, int, error)
	Read_MitchBitField() (MitchBitField, int, error)
	Read_MitchTime() (MitchTime, int, error)
	Read_MitchUInt64() (MitchUInt64, int, error)
}

