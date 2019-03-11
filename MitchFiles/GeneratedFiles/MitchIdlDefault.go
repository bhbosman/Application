package GeneratedFiles
import "errors"
import "fmt"
import "github.com/bhbosman/Application/Streams"

// MitchAlpha Declaration TypeCode: 0x451c2c75
// MitchAlpha WriteMessage 
func WriteMessage_MitchAlpha(stream Streams.IMitchWriter, value Streams.MitchAlpha) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x451c2c75)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchAlpha(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchAlpha reader
func ReadMessage_MitchAlpha(stream Streams.IMitchReader) (Streams.MitchAlpha, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return "", 0, err
	}
	if typeCode != 0x451c2c75 {
		return "", 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchAlpha. Expected 0x451c2c75, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchAlpha, n, err := Read_MitchAlpha(stream)
	if err != nil {
		return "", 0, err
	}
	byteCount += n
	return value_MitchAlpha, byteCount, nil
}

// MitchAlpha TypeCode
func TypeCode_MitchAlpha() uint32 {
	return 0x451c2c75
}

// MitchAlpha IsTypeCode
func IsTypeCode_MitchAlpha(typeCode uint32) bool {
	return typeCode == 0x451c2c75
}

// MitchAlpha writer 
func Write_MitchAlpha(stream Streams.IMitchWriter, value Streams.MitchAlpha) (int, error) {
	return stream.Write_MitchAlpha(value)
}

// MitchAlpha reader
func Read_MitchAlpha(stream Streams.IStreamReader) (Streams.MitchAlpha, int, error) {
	return stream.Read_MitchAlpha()
}

// MitchBitField Declaration TypeCode: 0x5c92bcfb
// MitchBitField WriteMessage 
func WriteMessage_MitchBitField(stream Streams.IMitchWriter, value Streams.MitchBitField) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x5c92bcfb)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchBitField(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchBitField reader
func ReadMessage_MitchBitField(stream Streams.IMitchReader) (Streams.MitchBitField, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0x5c92bcfb {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchBitField. Expected 0x5c92bcfb, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchBitField, n, err := Read_MitchBitField(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchBitField, byteCount, nil
}

// MitchBitField TypeCode
func TypeCode_MitchBitField() uint32 {
	return 0x5c92bcfb
}

// MitchBitField IsTypeCode
func IsTypeCode_MitchBitField(typeCode uint32) bool {
	return typeCode == 0x5c92bcfb
}

// MitchBitField writer 
func Write_MitchBitField(stream Streams.IMitchWriter, value Streams.MitchBitField) (int, error) {
	return stream.Write_MitchBitField(value)
}

// MitchBitField reader
func Read_MitchBitField(stream Streams.IStreamReader) (Streams.MitchBitField, int, error) {
	return stream.Read_MitchBitField()
}

// MitchByte Declaration TypeCode: 0xcdd7c459
// MitchByte WriteMessage 
func WriteMessage_MitchByte(stream Streams.IMitchWriter, value Streams.MitchByte) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xcdd7c459)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchByte(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchByte reader
func ReadMessage_MitchByte(stream Streams.IMitchReader) (Streams.MitchByte, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xcdd7c459 {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchByte. Expected 0xcdd7c459, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchByte, n, err := Read_MitchByte(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchByte, byteCount, nil
}

// MitchByte TypeCode
func TypeCode_MitchByte() uint32 {
	return 0xcdd7c459
}

// MitchByte IsTypeCode
func IsTypeCode_MitchByte(typeCode uint32) bool {
	return typeCode == 0xcdd7c459
}

// MitchByte writer 
func Write_MitchByte(stream Streams.IMitchWriter, value Streams.MitchByte) (int, error) {
	return stream.Write_MitchByte(value)
}

// MitchByte reader
func Read_MitchByte(stream Streams.IStreamReader) (Streams.MitchByte, int, error) {
	return stream.Read_MitchByte()
}

// MitchDate Declaration TypeCode: 0xdc8fbcd0
// MitchDate WriteMessage 
func WriteMessage_MitchDate(stream Streams.IMitchWriter, value Streams.MitchDate) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xdc8fbcd0)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchDate(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchDate reader
func ReadMessage_MitchDate(stream Streams.IMitchReader) (Streams.MitchDate, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xdc8fbcd0 {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchDate. Expected 0xdc8fbcd0, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchDate, n, err := Read_MitchDate(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchDate, byteCount, nil
}

// MitchDate TypeCode
func TypeCode_MitchDate() uint32 {
	return 0xdc8fbcd0
}

// MitchDate IsTypeCode
func IsTypeCode_MitchDate(typeCode uint32) bool {
	return typeCode == 0xdc8fbcd0
}

// MitchDate writer 
func Write_MitchDate(stream Streams.IMitchWriter, value Streams.MitchDate) (int, error) {
	return stream.Write_MitchDate(value)
}

// MitchDate reader
func Read_MitchDate(stream Streams.IStreamReader) (Streams.MitchDate, int, error) {
	return stream.Read_MitchDate()
}

// MitchTime Declaration TypeCode: 0x7c2ad58e
// MitchTime WriteMessage 
func WriteMessage_MitchTime(stream Streams.IMitchWriter, value Streams.MitchTime) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x7c2ad58e)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchTime(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchTime reader
func ReadMessage_MitchTime(stream Streams.IMitchReader) (Streams.MitchTime, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0x7c2ad58e {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchTime. Expected 0x7c2ad58e, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchTime, n, err := Read_MitchTime(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchTime, byteCount, nil
}

// MitchTime TypeCode
func TypeCode_MitchTime() uint32 {
	return 0x7c2ad58e
}

// MitchTime IsTypeCode
func IsTypeCode_MitchTime(typeCode uint32) bool {
	return typeCode == 0x7c2ad58e
}

// MitchTime writer 
func Write_MitchTime(stream Streams.IMitchWriter, value Streams.MitchTime) (int, error) {
	return stream.Write_MitchTime(value)
}

// MitchTime reader
func Read_MitchTime(stream Streams.IStreamReader) (Streams.MitchTime, int, error) {
	return stream.Read_MitchTime()
}

// MitchPrice04 Declaration TypeCode: 0xd693f5c7
// MitchPrice04 WriteMessage 
func WriteMessage_MitchPrice04(stream Streams.IMitchWriter, value Streams.MitchPrice04) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xd693f5c7)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchPrice04(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchPrice04 reader
func ReadMessage_MitchPrice04(stream Streams.IMitchReader) (Streams.MitchPrice04, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xd693f5c7 {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchPrice04. Expected 0xd693f5c7, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchPrice04, n, err := Read_MitchPrice04(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchPrice04, byteCount, nil
}

// MitchPrice04 TypeCode
func TypeCode_MitchPrice04() uint32 {
	return 0xd693f5c7
}

// MitchPrice04 IsTypeCode
func IsTypeCode_MitchPrice04(typeCode uint32) bool {
	return typeCode == 0xd693f5c7
}

// MitchPrice04 writer 
func Write_MitchPrice04(stream Streams.IMitchWriter, value Streams.MitchPrice04) (int, error) {
	return stream.Write_MitchPrice04(value)
}

// MitchPrice04 reader
func Read_MitchPrice04(stream Streams.IStreamReader) (Streams.MitchPrice04, int, error) {
	return stream.Read_MitchPrice04()
}

// MitchPrice08 Declaration TypeCode: 0xd03fe1d3
// MitchPrice08 WriteMessage 
func WriteMessage_MitchPrice08(stream Streams.IMitchWriter, value Streams.MitchPrice08) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xd03fe1d3)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchPrice08(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchPrice08 reader
func ReadMessage_MitchPrice08(stream Streams.IMitchReader) (Streams.MitchPrice08, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xd03fe1d3 {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchPrice08. Expected 0xd03fe1d3, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchPrice08, n, err := Read_MitchPrice08(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchPrice08, byteCount, nil
}

// MitchPrice08 TypeCode
func TypeCode_MitchPrice08() uint32 {
	return 0xd03fe1d3
}

// MitchPrice08 IsTypeCode
func IsTypeCode_MitchPrice08(typeCode uint32) bool {
	return typeCode == 0xd03fe1d3
}

// MitchPrice08 writer 
func Write_MitchPrice08(stream Streams.IMitchWriter, value Streams.MitchPrice08) (int, error) {
	return stream.Write_MitchPrice08(value)
}

// MitchPrice08 reader
func Read_MitchPrice08(stream Streams.IStreamReader) (Streams.MitchPrice08, int, error) {
	return stream.Read_MitchPrice08()
}

// MitchUInt08 Declaration TypeCode: 0xae0e9249
// MitchUInt08 WriteMessage 
func WriteMessage_MitchUInt08(stream Streams.IMitchWriter, value Streams.MitchUInt08) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xae0e9249)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchUInt08(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchUInt08 reader
func ReadMessage_MitchUInt08(stream Streams.IMitchReader) (Streams.MitchUInt08, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xae0e9249 {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchUInt08. Expected 0xae0e9249, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchUInt08, n, err := Read_MitchUInt08(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchUInt08, byteCount, nil
}

// MitchUInt08 TypeCode
func TypeCode_MitchUInt08() uint32 {
	return 0xae0e9249
}

// MitchUInt08 IsTypeCode
func IsTypeCode_MitchUInt08(typeCode uint32) bool {
	return typeCode == 0xae0e9249
}

// MitchUInt08 writer 
func Write_MitchUInt08(stream Streams.IMitchWriter, value Streams.MitchUInt08) (int, error) {
	return stream.Write_MitchUInt08(value)
}

// MitchUInt08 reader
func Read_MitchUInt08(stream Streams.IStreamReader) (Streams.MitchUInt08, int, error) {
	return stream.Read_MitchUInt08()
}

// MitchUInt16 Declaration TypeCode: 0xa83b1c5d
// MitchUInt16 WriteMessage 
func WriteMessage_MitchUInt16(stream Streams.IMitchWriter, value Streams.MitchUInt16) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xa83b1c5d)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchUInt16(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchUInt16 reader
func ReadMessage_MitchUInt16(stream Streams.IMitchReader) (Streams.MitchUInt16, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xa83b1c5d {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchUInt16. Expected 0xa83b1c5d, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchUInt16, n, err := Read_MitchUInt16(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchUInt16, byteCount, nil
}

// MitchUInt16 TypeCode
func TypeCode_MitchUInt16() uint32 {
	return 0xa83b1c5d
}

// MitchUInt16 IsTypeCode
func IsTypeCode_MitchUInt16(typeCode uint32) bool {
	return typeCode == 0xa83b1c5d
}

// MitchUInt16 writer 
func Write_MitchUInt16(stream Streams.IMitchWriter, value Streams.MitchUInt16) (int, error) {
	return stream.Write_MitchUInt16(value)
}

// MitchUInt16 reader
func Read_MitchUInt16(stream Streams.IStreamReader) (Streams.MitchUInt16, int, error) {
	return stream.Read_MitchUInt16()
}

// MitchUInt32 Declaration TypeCode: 0xa908285d
// MitchUInt32 WriteMessage 
func WriteMessage_MitchUInt32(stream Streams.IMitchWriter, value Streams.MitchUInt32) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xa908285d)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchUInt32(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchUInt32 reader
func ReadMessage_MitchUInt32(stream Streams.IMitchReader) (Streams.MitchUInt32, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xa908285d {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchUInt32. Expected 0xa908285d, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchUInt32, n, err := Read_MitchUInt32(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchUInt32, byteCount, nil
}

// MitchUInt32 TypeCode
func TypeCode_MitchUInt32() uint32 {
	return 0xa908285d
}

// MitchUInt32 IsTypeCode
func IsTypeCode_MitchUInt32(typeCode uint32) bool {
	return typeCode == 0xa908285d
}

// MitchUInt32 writer 
func Write_MitchUInt32(stream Streams.IMitchWriter, value Streams.MitchUInt32) (int, error) {
	return stream.Write_MitchUInt32(value)
}

// MitchUInt32 reader
func Read_MitchUInt32(stream Streams.IStreamReader) (Streams.MitchUInt32, int, error) {
	return stream.Read_MitchUInt32()
}

// MitchUInt64 Declaration TypeCode: 0xad5bce49
// MitchUInt64 WriteMessage 
func WriteMessage_MitchUInt64(stream Streams.IMitchWriter, value Streams.MitchUInt64) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xad5bce49)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchUInt64(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchUInt64 reader
func ReadMessage_MitchUInt64(stream Streams.IMitchReader) (Streams.MitchUInt64, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return false, 0, err
	}
	if typeCode != 0xad5bce49 {
		return false, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchUInt64. Expected 0xad5bce49, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchUInt64, n, err := Read_MitchUInt64(stream)
	if err != nil {
		return false, 0, err
	}
	byteCount += n
	return value_MitchUInt64, byteCount, nil
}

// MitchUInt64 TypeCode
func TypeCode_MitchUInt64() uint32 {
	return 0xad5bce49
}

// MitchUInt64 IsTypeCode
func IsTypeCode_MitchUInt64(typeCode uint32) bool {
	return typeCode == 0xad5bce49
}

// MitchUInt64 writer 
func Write_MitchUInt64(stream Streams.IMitchWriter, value Streams.MitchUInt64) (int, error) {
	return stream.Write_MitchUInt64(value)
}

// MitchUInt64 reader
func Read_MitchUInt64(stream Streams.IStreamReader) (Streams.MitchUInt64, int, error) {
	return stream.Read_MitchUInt64()
}

