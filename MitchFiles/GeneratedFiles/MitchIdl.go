

package GeneratedFiles
import "errors"
import "fmt"
import "github.com/bhbosman/Application/Streams"

// MitchMessage_Time Declaration TypeCode: 0x721c69d4
type MitchMessage_Time struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(84)
 	Seconds Streams.MitchUInt32
}

func NewMitchMessage_Time()*MitchMessage_Time {
	return &MitchMessage_Time{}
}

// MitchMessage_Time writer 
func Write_MitchMessage_Time(stream Streams.IMitchWriter, value *MitchMessage_Time) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Seconds)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_Time reader
func Read_MitchMessage_Time(stream Streams.IStreamReader) (value *MitchMessage_Time, byteCount int, err error) {
	value = NewMitchMessage_Time()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Seconds, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_Time WriteMessage 
func WriteMessage_MitchMessage_Time(stream Streams.IMitchWriter, value *MitchMessage_Time) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x721c69d4)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_Time(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_Time reader
func ReadMessage_MitchMessage_Time(stream Streams.IMitchReader) (*MitchMessage_Time, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x721c69d4 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_Time. Expected 0x721c69d4, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_Time, n, err := Read_MitchMessage_Time(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_Time, byteCount, nil
}

// MitchMessage_Time TypeCode
func TypeCode_MitchMessage_Time() uint32 {
	return 0x721c69d4
}

// MitchMessage_Time IsTypeCode
func IsTypeCode_MitchMessage_Time(typeCode uint32) bool {
	return typeCode == 0x721c69d4
}

type TEventCode byte 
//noinspection ALL
const (
	TEventCode_EndOfDay  = 'C' // default value: byte('C')
 	TEventCode_StartOfDay  = 'O' // default value: byte('O')
 )

// TEventCode Declaration TypeCode: 0x2ff833d1
// TEventCode WriteMessage 
func WriteMessage_TEventCode(stream Streams.IMitchWriter, value TEventCode) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x2ff833d1)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TEventCode(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TEventCode reader
func ReadMessage_TEventCode(stream Streams.IMitchReader) (TEventCode, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x2ff833d1 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading TEventCode. Expected 0x2ff833d1, got 0x%08x", typeCode))
	}
	byteCount += n
	value_TEventCode, n, err := Read_TEventCode(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_TEventCode, byteCount, nil
}

// MitchMessage_SystemEvent Declaration TypeCode: 0xf41fb8e9
type MitchMessage_SystemEvent struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(83)
 	Nanosecond Streams.MitchUInt32
	EventCode TEventCode
}

func NewMitchMessage_SystemEvent()*MitchMessage_SystemEvent {
	return &MitchMessage_SystemEvent{}
}

// MitchMessage_SystemEvent writer 
func Write_MitchMessage_SystemEvent(stream Streams.IMitchWriter, value *MitchMessage_SystemEvent) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.EventCode)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_SystemEvent reader
func Read_MitchMessage_SystemEvent(stream Streams.IStreamReader) (value *MitchMessage_SystemEvent, byteCount int, err error) {
	value = NewMitchMessage_SystemEvent()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.EventCode, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_SystemEvent WriteMessage 
func WriteMessage_MitchMessage_SystemEvent(stream Streams.IMitchWriter, value *MitchMessage_SystemEvent) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xf41fb8e9)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_SystemEvent(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_SystemEvent reader
func ReadMessage_MitchMessage_SystemEvent(stream Streams.IMitchReader) (*MitchMessage_SystemEvent, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xf41fb8e9 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_SystemEvent. Expected 0xf41fb8e9, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_SystemEvent, n, err := Read_MitchMessage_SystemEvent(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_SystemEvent, byteCount, nil
}

// MitchMessage_SystemEvent TypeCode
func TypeCode_MitchMessage_SystemEvent() uint32 {
	return 0xf41fb8e9
}

// MitchMessage_SystemEvent IsTypeCode
func IsTypeCode_MitchMessage_SystemEvent(typeCode uint32) bool {
	return typeCode == 0xf41fb8e9
}

type SymbolDirectorySymbolStatus byte 
//noinspection ALL
const (
	SymbolDirectorySymbolStatus_Halted  = 'H' // default value: byte('H')
 	SymbolDirectorySymbolStatus_Suspended  = 'S' // default value: byte('S')
 	SymbolDirectorySymbolStatus_Inactive  = 'a' // default value: byte('a')
 )

// SymbolDirectorySymbolStatus Declaration TypeCode: 0x810d8fde
// SymbolDirectorySymbolStatus WriteMessage 
func WriteMessage_SymbolDirectorySymbolStatus(stream Streams.IMitchWriter, value SymbolDirectorySymbolStatus) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x810d8fde)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolDirectorySymbolStatus(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectorySymbolStatus reader
func ReadMessage_SymbolDirectorySymbolStatus(stream Streams.IMitchReader) (SymbolDirectorySymbolStatus, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x810d8fde {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading SymbolDirectorySymbolStatus. Expected 0x810d8fde, got 0x%08x", typeCode))
	}
	byteCount += n
	value_SymbolDirectorySymbolStatus, n, err := Read_SymbolDirectorySymbolStatus(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolDirectorySymbolStatus, byteCount, nil
}

type SymbolDirectoryOptionType byte 
//noinspection ALL
const (
	SymbolDirectoryOptionType_CallOption  = 'C' // default value: byte('C')
 	SymbolDirectoryOptionType_PutOption  = 'P' // default value: byte('P')
 )

// SymbolDirectoryOptionType Declaration TypeCode: 0x8ca314f2
// SymbolDirectoryOptionType WriteMessage 
func WriteMessage_SymbolDirectoryOptionType(stream Streams.IMitchWriter, value SymbolDirectoryOptionType) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x8ca314f2)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolDirectoryOptionType(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectoryOptionType reader
func ReadMessage_SymbolDirectoryOptionType(stream Streams.IMitchReader) (SymbolDirectoryOptionType, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x8ca314f2 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading SymbolDirectoryOptionType. Expected 0x8ca314f2, got 0x%08x", typeCode))
	}
	byteCount += n
	value_SymbolDirectoryOptionType, n, err := Read_SymbolDirectoryOptionType(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolDirectoryOptionType, byteCount, nil
}

type SymbolDirectoryFlags struct{
	InverseOrderBook bool // Bit 00
}

type SymbolDirectorySubBook struct{
	Regular bool // Bit 00
	OffBook bool // Bit 01
	BulletinBoard bool // Bit 05
	NegotiatedTrades bool // Bit 06
}

type SymbolDirectorySettlementMethod byte 
//noinspection ALL
const (
	SymbolDirectorySettlementMethod_Cash  = 'C' // default value: byte('C')
 	SymbolDirectorySettlementMethod_Physical  = 'P' // default value: byte('P')
 )

// SymbolDirectorySettlementMethod Declaration TypeCode: 0xc1ac87fd
// SymbolDirectorySettlementMethod WriteMessage 
func WriteMessage_SymbolDirectorySettlementMethod(stream Streams.IMitchWriter, value SymbolDirectorySettlementMethod) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xc1ac87fd)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolDirectorySettlementMethod(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectorySettlementMethod reader
func ReadMessage_SymbolDirectorySettlementMethod(stream Streams.IMitchReader) (SymbolDirectorySettlementMethod, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0xc1ac87fd {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading SymbolDirectorySettlementMethod. Expected 0xc1ac87fd, got 0x%08x", typeCode))
	}
	byteCount += n
	value_SymbolDirectorySettlementMethod, n, err := Read_SymbolDirectorySettlementMethod(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolDirectorySettlementMethod, byteCount, nil
}

// MitchMessage_SymbolDirectory Declaration TypeCode: 0x31da5e63
type MitchMessage_SymbolDirectory struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(82)
 	Nanosecond Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	SymbolStatus SymbolDirectorySymbolStatus
	ISIN Streams.MitchAlpha
	Symbol Streams.MitchAlpha
	TIDM Streams.MitchAlpha
	Segment Streams.MitchAlpha
	PreviousClosePrice Streams.MitchPrice08
	ExpirationDate Streams.MitchDate
	Underlying Streams.MitchAlpha
	StrikePrice Streams.MitchPrice08
	OptionType SymbolDirectoryOptionType
	Issuer Streams.MitchAlpha
	IssueDate Streams.MitchDate
	Coupon Streams.MitchPrice08
	Flags SymbolDirectoryFlags
	SubBook SymbolDirectorySubBook
	CorporateAction Streams.MitchAlpha
	Leg1Symbol Streams.MitchAlpha
	Leg2Symbol Streams.MitchAlpha
	ContractMultiplier Streams.MitchPrice08
	SettlementMethod SymbolDirectorySettlementMethod
	InstrumentSubCategory Streams.MitchAlpha
}

func NewMitchMessage_SymbolDirectory()*MitchMessage_SymbolDirectory {
	return &MitchMessage_SymbolDirectory{}
}

// MitchMessage_SymbolDirectory writer 
func Write_MitchMessage_SymbolDirectory(stream Streams.IMitchWriter, value *MitchMessage_SymbolDirectory) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.SymbolStatus)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.ISIN)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Symbol)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.TIDM)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Segment)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.PreviousClosePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchDate(value.ExpirationDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Underlying)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.StrikePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.OptionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Issuer)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchDate(value.IssueDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Coupon)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_TypeDeclarator(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_TypeDeclarator(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.CorporateAction)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Leg1Symbol)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Leg2Symbol)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.ContractMultiplier)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.SettlementMethod)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.InstrumentSubCategory)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_SymbolDirectory reader
func Read_MitchMessage_SymbolDirectory(stream Streams.IStreamReader) (value *MitchMessage_SymbolDirectory, byteCount int, err error) {
	value = NewMitchMessage_SymbolDirectory()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SymbolStatus, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ISIN, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Symbol, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TIDM, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Segment, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.PreviousClosePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ExpirationDate, n, err = stream.Read_MitchDate()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Underlying, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.StrikePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OptionType, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Issuer, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.IssueDate, n, err = stream.Read_MitchDate()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Coupon, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_TypeDeclarator()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_TypeDeclarator()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.CorporateAction, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Leg1Symbol, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Leg2Symbol, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ContractMultiplier, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SettlementMethod, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentSubCategory, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_SymbolDirectory WriteMessage 
func WriteMessage_MitchMessage_SymbolDirectory(stream Streams.IMitchWriter, value *MitchMessage_SymbolDirectory) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x31da5e63)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_SymbolDirectory(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_SymbolDirectory reader
func ReadMessage_MitchMessage_SymbolDirectory(stream Streams.IMitchReader) (*MitchMessage_SymbolDirectory, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x31da5e63 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_SymbolDirectory. Expected 0x31da5e63, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_SymbolDirectory, n, err := Read_MitchMessage_SymbolDirectory(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_SymbolDirectory, byteCount, nil
}

// MitchMessage_SymbolDirectory TypeCode
func TypeCode_MitchMessage_SymbolDirectory() uint32 {
	return 0x31da5e63
}

// MitchMessage_SymbolDirectory IsTypeCode
func IsTypeCode_MitchMessage_SymbolDirectory(typeCode uint32) bool {
	return typeCode == 0x31da5e63
}

type SymbolStatusTradingStatus byte 
//noinspection ALL
const (
	SymbolStatusTradingStatus_Halt  = 'H' // default value: byte('H')
 	SymbolStatusTradingStatus_RegularTrading  = 'T' // default value: byte('T')
 	SymbolStatusTradingStatus_OpeningAuctionCall  = 'a' // default value: byte('a')
 	SymbolStatusTradingStatus_PostClose  = 'b' // default value: byte('b')
 	SymbolStatusTradingStatus_MarketClose  = 'c' // default value: byte('c')
 	SymbolStatusTradingStatus_ClosingAuctionCall  = 'd' // default value: byte('d')
 	SymbolStatusTradingStatus_VolatilityAuctionCall  = 'e' // default value: byte('e')
 	SymbolStatusTradingStatus_EODVolumeAuctionCall  = 'E' // default value: byte('E')
 	SymbolStatusTradingStatus_ReOpeningAuctionCall  = 'f' // default value: byte('f')
 	SymbolStatusTradingStatus_Pause  = 'l' // default value: byte('l')
 	SymbolStatusTradingStatus_FuturesCloseOut  = 'p' // default value: byte('p')
 	SymbolStatusTradingStatus_ClosingPriceCross  = 's' // default value: byte('s')
 	SymbolStatusTradingStatus_IntraDayAuctionCall  = 'u' // default value: byte('u')
 	SymbolStatusTradingStatus_EndTradeReporting  = 'v' // default value: byte('v')
 	SymbolStatusTradingStatus_NoActiveSession  = 'w' // default value: byte('w')
 	SymbolStatusTradingStatus_EndOfPostClose  = 'x' // default value: byte('x')
 	SymbolStatusTradingStatus_StarOofTrading  = 'y' // default value: byte('y')
 	SymbolStatusTradingStatus_ClosingPricePublication  = 'z' // default value: byte('z')
 )

// SymbolStatusTradingStatus Declaration TypeCode: 0xdeb92876
// SymbolStatusTradingStatus WriteMessage 
func WriteMessage_SymbolStatusTradingStatus(stream Streams.IMitchWriter, value SymbolStatusTradingStatus) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xdeb92876)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolStatusTradingStatus(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolStatusTradingStatus reader
func ReadMessage_SymbolStatusTradingStatus(stream Streams.IMitchReader) (SymbolStatusTradingStatus, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0xdeb92876 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading SymbolStatusTradingStatus. Expected 0xdeb92876, got 0x%08x", typeCode))
	}
	byteCount += n
	value_SymbolStatusTradingStatus, n, err := Read_SymbolStatusTradingStatus(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolStatusTradingStatus, byteCount, nil
}

type SymbolStatusSessionChangeReason byte 
//noinspection ALL
const (
	SymbolStatusSessionChangeReason_ScheduledTransition  = 0 // default value: byte(0)
 	SymbolStatusSessionChangeReason_ExtendedByMarketOps  = 1 // default value: byte(1)
 	SymbolStatusSessionChangeReason_ShortenedByMarketOps  = 2 // default value: byte(2)
 	SymbolStatusSessionChangeReason_MarketOrderImbalance  = 3 // default value: byte(3)
 	SymbolStatusSessionChangeReason_PriceOutsideRange  = 4 // default value: byte(4)
 	SymbolStatusSessionChangeReason_CircuitBreakerTripped  = 5 // default value: byte(5)
 	SymbolStatusSessionChangeReason_Unavailable  = 9 // default value: byte(9)
 )

// SymbolStatusSessionChangeReason Declaration TypeCode: 0xfee8af4f
// SymbolStatusSessionChangeReason WriteMessage 
func WriteMessage_SymbolStatusSessionChangeReason(stream Streams.IMitchWriter, value SymbolStatusSessionChangeReason) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xfee8af4f)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolStatusSessionChangeReason(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolStatusSessionChangeReason reader
func ReadMessage_SymbolStatusSessionChangeReason(stream Streams.IMitchReader) (SymbolStatusSessionChangeReason, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0xfee8af4f {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading SymbolStatusSessionChangeReason. Expected 0xfee8af4f, got 0x%08x", typeCode))
	}
	byteCount += n
	value_SymbolStatusSessionChangeReason, n, err := Read_SymbolStatusSessionChangeReason(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolStatusSessionChangeReason, byteCount, nil
}

type SymbolStatusBookType byte 
//noinspection ALL
const (
	SymbolStatusBookType_OnBook  = 1 // default value: byte(1)
 	SymbolStatusBookType_OffBook  = 2 // default value: byte(2)
 	SymbolStatusBookType_BulletinBoard  = 9 // default value: byte(9)
 	SymbolStatusBookType_NegotiatedTrades  = 11 // default value: byte(11)
 )

// SymbolStatusBookType Declaration TypeCode: 0x0b2355a5
// SymbolStatusBookType WriteMessage 
func WriteMessage_SymbolStatusBookType(stream Streams.IMitchWriter, value SymbolStatusBookType) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x0b2355a5)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolStatusBookType(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolStatusBookType reader
func ReadMessage_SymbolStatusBookType(stream Streams.IMitchReader) (SymbolStatusBookType, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x0b2355a5 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading SymbolStatusBookType. Expected 0x0b2355a5, got 0x%08x", typeCode))
	}
	byteCount += n
	value_SymbolStatusBookType, n, err := Read_SymbolStatusBookType(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolStatusBookType, byteCount, nil
}

// MitchMessage_SymbolStatus Declaration TypeCode: 0xdfd8d6fb
type MitchMessage_SymbolStatus struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(72)
 	Nanosecond Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	TradingStatus SymbolStatusTradingStatus
	Flags Streams.MitchBitField
	Reason Streams.MitchAlpha
	SessionChangeReason SymbolStatusSessionChangeReason
	NewEndTime Streams.MitchTime
	BookType SymbolStatusBookType
}

func NewMitchMessage_SymbolStatus()*MitchMessage_SymbolStatus {
	return &MitchMessage_SymbolStatus{}
}

// MitchMessage_SymbolStatus writer 
func Write_MitchMessage_SymbolStatus(stream Streams.IMitchWriter, value *MitchMessage_SymbolStatus) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.TradingStatus)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchBitField(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Reason)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.SessionChangeReason)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchTime(value.NewEndTime)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.BookType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_SymbolStatus reader
func Read_MitchMessage_SymbolStatus(stream Streams.IStreamReader) (value *MitchMessage_SymbolStatus, byteCount int, err error) {
	value = NewMitchMessage_SymbolStatus()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradingStatus, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_MitchBitField()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reason, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SessionChangeReason, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.NewEndTime, n, err = stream.Read_MitchTime()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.BookType, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_SymbolStatus WriteMessage 
func WriteMessage_MitchMessage_SymbolStatus(stream Streams.IMitchWriter, value *MitchMessage_SymbolStatus) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xdfd8d6fb)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_SymbolStatus(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_SymbolStatus reader
func ReadMessage_MitchMessage_SymbolStatus(stream Streams.IMitchReader) (*MitchMessage_SymbolStatus, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xdfd8d6fb {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_SymbolStatus. Expected 0xdfd8d6fb, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_SymbolStatus, n, err := Read_MitchMessage_SymbolStatus(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_SymbolStatus, byteCount, nil
}

// MitchMessage_SymbolStatus TypeCode
func TypeCode_MitchMessage_SymbolStatus() uint32 {
	return 0xdfd8d6fb
}

// MitchMessage_SymbolStatus IsTypeCode
func IsTypeCode_MitchMessage_SymbolStatus(typeCode uint32) bool {
	return typeCode == 0xdfd8d6fb
}

type OrderSide byte 
//noinspection ALL
const (
	OrderSide_BuyOrder  = 'B' // default value: byte('B')
 	OrderSide_SellOrder  = 'S' // default value: byte('S')
 )

// OrderSide Declaration TypeCode: 0x22a70cb8
// OrderSide WriteMessage 
func WriteMessage_OrderSide(stream Streams.IMitchWriter, value OrderSide) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x22a70cb8)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderSide(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderSide reader
func ReadMessage_OrderSide(stream Streams.IMitchReader) (OrderSide, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x22a70cb8 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading OrderSide. Expected 0x22a70cb8, got 0x%08x", typeCode))
	}
	byteCount += n
	value_OrderSide, n, err := Read_OrderSide(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_OrderSide, byteCount, nil
}

type AddOrderFlags struct{
	MarketOrder bool // Bit 04
	BulletinBoard bool // Bit 05
}

// MitchMessage_AddOrder Declaration TypeCode: 0x760f4121
type MitchMessage_AddOrder struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(65)
 	Nanosecond Streams.MitchUInt32
	OrderId Streams.MitchUInt64
	Side OrderSide
	Quantity Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	Price Streams.MitchPrice08
	Flags AddOrderFlags
}

func NewMitchMessage_AddOrder()*MitchMessage_AddOrder {
	return &MitchMessage_AddOrder{}
}

// MitchMessage_AddOrder writer 
func Write_MitchMessage_AddOrder(stream Streams.IMitchWriter, value *MitchMessage_AddOrder) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.OrderId)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.Side)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_TypeDeclarator(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AddOrder reader
func Read_MitchMessage_AddOrder(stream Streams.IStreamReader) (value *MitchMessage_AddOrder, byteCount int, err error) {
	value = NewMitchMessage_AddOrder()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OrderId, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Side, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Quantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_TypeDeclarator()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_AddOrder WriteMessage 
func WriteMessage_MitchMessage_AddOrder(stream Streams.IMitchWriter, value *MitchMessage_AddOrder) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x760f4121)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_AddOrder(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AddOrder reader
func ReadMessage_MitchMessage_AddOrder(stream Streams.IMitchReader) (*MitchMessage_AddOrder, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x760f4121 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_AddOrder. Expected 0x760f4121, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_AddOrder, n, err := Read_MitchMessage_AddOrder(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_AddOrder, byteCount, nil
}

// MitchMessage_AddOrder TypeCode
func TypeCode_MitchMessage_AddOrder() uint32 {
	return 0x760f4121
}

// MitchMessage_AddOrder IsTypeCode
func IsTypeCode_MitchMessage_AddOrder(typeCode uint32) bool {
	return typeCode == 0x760f4121
}

type AddAttributedOrderFlags struct{
	Regular bool // Bit 00
	BulletinBoard bool // Bit 05
}

// MitchMessage_AddAttributedOrder Declaration TypeCode: 0x6e2d3a04
type MitchMessage_AddAttributedOrder struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(70)
 	Nanosecond Streams.MitchUInt32
	OrderID Streams.MitchUInt64
	Side OrderSide
	Quantity Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Price Streams.MitchPrice08
	Attribution Streams.MitchAlpha
	Flags AddAttributedOrderFlags
}

func NewMitchMessage_AddAttributedOrder()*MitchMessage_AddAttributedOrder {
	return &MitchMessage_AddAttributedOrder{}
}

// MitchMessage_AddAttributedOrder writer 
func Write_MitchMessage_AddAttributedOrder(stream Streams.IMitchWriter, value *MitchMessage_AddAttributedOrder) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.Side)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Attribution)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_TypeDeclarator(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AddAttributedOrder reader
func Read_MitchMessage_AddAttributedOrder(stream Streams.IStreamReader) (value *MitchMessage_AddAttributedOrder, byteCount int, err error) {
	value = NewMitchMessage_AddAttributedOrder()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OrderID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Side, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Quantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Attribution, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_TypeDeclarator()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_AddAttributedOrder WriteMessage 
func WriteMessage_MitchMessage_AddAttributedOrder(stream Streams.IMitchWriter, value *MitchMessage_AddAttributedOrder) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x6e2d3a04)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_AddAttributedOrder(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AddAttributedOrder reader
func ReadMessage_MitchMessage_AddAttributedOrder(stream Streams.IMitchReader) (*MitchMessage_AddAttributedOrder, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x6e2d3a04 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_AddAttributedOrder. Expected 0x6e2d3a04, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_AddAttributedOrder, n, err := Read_MitchMessage_AddAttributedOrder(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_AddAttributedOrder, byteCount, nil
}

// MitchMessage_AddAttributedOrder TypeCode
func TypeCode_MitchMessage_AddAttributedOrder() uint32 {
	return 0x6e2d3a04
}

// MitchMessage_AddAttributedOrder IsTypeCode
func IsTypeCode_MitchMessage_AddAttributedOrder(typeCode uint32) bool {
	return typeCode == 0x6e2d3a04
}

// MitchMessage_OrderDeleted Declaration TypeCode: 0x3985c089
type MitchMessage_OrderDeleted struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(68)
 	Nanosecond Streams.MitchUInt32
	OrderID Streams.MitchUInt64
}

func NewMitchMessage_OrderDeleted()*MitchMessage_OrderDeleted {
	return &MitchMessage_OrderDeleted{}
}

// MitchMessage_OrderDeleted writer 
func Write_MitchMessage_OrderDeleted(stream Streams.IMitchWriter, value *MitchMessage_OrderDeleted) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderDeleted reader
func Read_MitchMessage_OrderDeleted(stream Streams.IStreamReader) (value *MitchMessage_OrderDeleted, byteCount int, err error) {
	value = NewMitchMessage_OrderDeleted()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OrderID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_OrderDeleted WriteMessage 
func WriteMessage_MitchMessage_OrderDeleted(stream Streams.IMitchWriter, value *MitchMessage_OrderDeleted) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x3985c089)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_OrderDeleted(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderDeleted reader
func ReadMessage_MitchMessage_OrderDeleted(stream Streams.IMitchReader) (*MitchMessage_OrderDeleted, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x3985c089 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_OrderDeleted. Expected 0x3985c089, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_OrderDeleted, n, err := Read_MitchMessage_OrderDeleted(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_OrderDeleted, byteCount, nil
}

// MitchMessage_OrderDeleted TypeCode
func TypeCode_MitchMessage_OrderDeleted() uint32 {
	return 0x3985c089
}

// MitchMessage_OrderDeleted IsTypeCode
func IsTypeCode_MitchMessage_OrderDeleted(typeCode uint32) bool {
	return typeCode == 0x3985c089
}

type OrderModifiedFlags struct{
	PriorityFlag bool // Bit 00
}

// MitchMessage_OrderModified Declaration TypeCode: 0xf9c68faa
type MitchMessage_OrderModified struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(85)
 	Nanosecond Streams.MitchUInt32
	OrderID Streams.MitchUInt64
	NewQuantity Streams.MitchUInt32
	NewPrice Streams.MitchPrice08
	Flags OrderModifiedFlags
}

func NewMitchMessage_OrderModified()*MitchMessage_OrderModified {
	return &MitchMessage_OrderModified{}
}

// MitchMessage_OrderModified writer 
func Write_MitchMessage_OrderModified(stream Streams.IMitchWriter, value *MitchMessage_OrderModified) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.NewQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.NewPrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_TypeDeclarator(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderModified reader
func Read_MitchMessage_OrderModified(stream Streams.IStreamReader) (value *MitchMessage_OrderModified, byteCount int, err error) {
	value = NewMitchMessage_OrderModified()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OrderID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.NewQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.NewPrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_TypeDeclarator()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_OrderModified WriteMessage 
func WriteMessage_MitchMessage_OrderModified(stream Streams.IMitchWriter, value *MitchMessage_OrderModified) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xf9c68faa)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_OrderModified(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderModified reader
func ReadMessage_MitchMessage_OrderModified(stream Streams.IMitchReader) (*MitchMessage_OrderModified, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xf9c68faa {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_OrderModified. Expected 0xf9c68faa, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_OrderModified, n, err := Read_MitchMessage_OrderModified(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_OrderModified, byteCount, nil
}

// MitchMessage_OrderModified TypeCode
func TypeCode_MitchMessage_OrderModified() uint32 {
	return 0xf9c68faa
}

// MitchMessage_OrderModified IsTypeCode
func IsTypeCode_MitchMessage_OrderModified(typeCode uint32) bool {
	return typeCode == 0xf9c68faa
}

type OrderBookClearSubBook byte 
//noinspection ALL
const (
	OrderBookClearSubBook_OnBook  = 1 // default value: byte(1)
 	OrderBookClearSubBook_OffBook  = 2 // default value: byte(2)
 	OrderBookClearSubBook_BulletinBoard  = 9 // default value: byte(9)
 	OrderBookClearSubBook_NegotiatedTrades  = 11 // default value: byte(11)
 )

// OrderBookClearSubBook Declaration TypeCode: 0x5ead6996
// OrderBookClearSubBook WriteMessage 
func WriteMessage_OrderBookClearSubBook(stream Streams.IMitchWriter, value OrderBookClearSubBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x5ead6996)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderBookClearSubBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderBookClearSubBook reader
func ReadMessage_OrderBookClearSubBook(stream Streams.IMitchReader) (OrderBookClearSubBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x5ead6996 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading OrderBookClearSubBook. Expected 0x5ead6996, got 0x%08x", typeCode))
	}
	byteCount += n
	value_OrderBookClearSubBook, n, err := Read_OrderBookClearSubBook(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_OrderBookClearSubBook, byteCount, nil
}

type OrderBookClearBookType byte 
//noinspection ALL
const (
	OrderBookClearBookType_MBO  = 0 // default value: byte(0)
 	OrderBookClearBookType_TopOfBook  = 1 // default value: byte(1)
 )

// OrderBookClearBookType Declaration TypeCode: 0x7d3b6454
// OrderBookClearBookType WriteMessage 
func WriteMessage_OrderBookClearBookType(stream Streams.IMitchWriter, value OrderBookClearBookType) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x7d3b6454)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderBookClearBookType(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderBookClearBookType reader
func ReadMessage_OrderBookClearBookType(stream Streams.IMitchReader) (OrderBookClearBookType, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x7d3b6454 {
		return 0, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading OrderBookClearBookType. Expected 0x7d3b6454, got 0x%08x", typeCode))
	}
	byteCount += n
	value_OrderBookClearBookType, n, err := Read_OrderBookClearBookType(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_OrderBookClearBookType, byteCount, nil
}

// MitchMessage_OrderBookClear Declaration TypeCode: 0x04dbefcc
type MitchMessage_OrderBookClear struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte // default value: Int64(121)
 	Nanosecond Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	SubBook OrderBookClearSubBook
	BookType OrderBookClearBookType
}

func NewMitchMessage_OrderBookClear()*MitchMessage_OrderBookClear {
	return &MitchMessage_OrderBookClear{}
}

// MitchMessage_OrderBookClear writer 
func Write_MitchMessage_OrderBookClear(stream Streams.IMitchWriter, value *MitchMessage_OrderBookClear) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_Enum(value.BookType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderBookClear reader
func Read_MitchMessage_OrderBookClear(stream Streams.IStreamReader) (value *MitchMessage_OrderBookClear, byteCount int, err error) {
	value = NewMitchMessage_OrderBookClear()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.BookType, n, err = stream.Read_Enum()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_OrderBookClear WriteMessage 
func WriteMessage_MitchMessage_OrderBookClear(stream Streams.IMitchWriter, value *MitchMessage_OrderBookClear) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x04dbefcc)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_OrderBookClear(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderBookClear reader
func ReadMessage_MitchMessage_OrderBookClear(stream Streams.IMitchReader) (*MitchMessage_OrderBookClear, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x04dbefcc {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_OrderBookClear. Expected 0x04dbefcc, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_OrderBookClear, n, err := Read_MitchMessage_OrderBookClear(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_OrderBookClear, byteCount, nil
}

// MitchMessage_OrderBookClear TypeCode
func TypeCode_MitchMessage_OrderBookClear() uint32 {
	return 0x04dbefcc
}

// MitchMessage_OrderBookClear IsTypeCode
func IsTypeCode_MitchMessage_OrderBookClear(typeCode uint32) bool {
	return typeCode == 0x04dbefcc
}

// MitchMessage_OrderExecuted Declaration TypeCode: 0xb22f1820
type MitchMessage_OrderExecuted struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	OrderID Streams.MitchUInt64
	ExecutedQuantity Streams.MitchUInt32
	TradeID Streams.MitchUInt64
	LastOptPx Streams.MitchPrice08
	Volatility Streams.MitchPrice08
	UnderlyingReferencePrice Streams.MitchPrice08
}

func NewMitchMessage_OrderExecuted()*MitchMessage_OrderExecuted {
	return &MitchMessage_OrderExecuted{}
}

// MitchMessage_OrderExecuted writer 
func Write_MitchMessage_OrderExecuted(stream Streams.IMitchWriter, value *MitchMessage_OrderExecuted) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderExecuted reader
func Read_MitchMessage_OrderExecuted(stream Streams.IStreamReader) (value *MitchMessage_OrderExecuted, byteCount int, err error) {
	value = NewMitchMessage_OrderExecuted()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OrderID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ExecutedQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LastOptPx, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volatility, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.UnderlyingReferencePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_OrderExecuted WriteMessage 
func WriteMessage_MitchMessage_OrderExecuted(stream Streams.IMitchWriter, value *MitchMessage_OrderExecuted) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xb22f1820)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_OrderExecuted(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderExecuted reader
func ReadMessage_MitchMessage_OrderExecuted(stream Streams.IMitchReader) (*MitchMessage_OrderExecuted, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xb22f1820 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_OrderExecuted. Expected 0xb22f1820, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_OrderExecuted, n, err := Read_MitchMessage_OrderExecuted(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_OrderExecuted, byteCount, nil
}

// MitchMessage_OrderExecuted TypeCode
func TypeCode_MitchMessage_OrderExecuted() uint32 {
	return 0xb22f1820
}

// MitchMessage_OrderExecuted IsTypeCode
func IsTypeCode_MitchMessage_OrderExecuted(typeCode uint32) bool {
	return typeCode == 0xb22f1820
}

// MitchMessage_OrderExecutedWithPriceSize Declaration TypeCode: 0x626b5ac2
type MitchMessage_OrderExecutedWithPriceSize struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	OrderID Streams.MitchUInt64
	ExecutedQuantity Streams.MitchUInt32
	DisplayQuantity Streams.MitchUInt32
	TradeID Streams.MitchUInt64
	Printable Streams.MitchByte
	Price Streams.MitchPrice08
	LastOptPx Streams.MitchPrice08
	Volatility Streams.MitchPrice08
	UnderlyingReferencePrice Streams.MitchPrice08
}

func NewMitchMessage_OrderExecutedWithPriceSize()*MitchMessage_OrderExecutedWithPriceSize {
	return &MitchMessage_OrderExecutedWithPriceSize{}
}

// MitchMessage_OrderExecutedWithPriceSize writer 
func Write_MitchMessage_OrderExecutedWithPriceSize(stream Streams.IMitchWriter, value *MitchMessage_OrderExecutedWithPriceSize) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.DisplayQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Printable)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderExecutedWithPriceSize reader
func Read_MitchMessage_OrderExecutedWithPriceSize(stream Streams.IStreamReader) (value *MitchMessage_OrderExecutedWithPriceSize, byteCount int, err error) {
	value = NewMitchMessage_OrderExecutedWithPriceSize()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OrderID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ExecutedQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.DisplayQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Printable, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LastOptPx, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volatility, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.UnderlyingReferencePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_OrderExecutedWithPriceSize WriteMessage 
func WriteMessage_MitchMessage_OrderExecutedWithPriceSize(stream Streams.IMitchWriter, value *MitchMessage_OrderExecutedWithPriceSize) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x626b5ac2)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_OrderExecutedWithPriceSize(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OrderExecutedWithPriceSize reader
func ReadMessage_MitchMessage_OrderExecutedWithPriceSize(stream Streams.IMitchReader) (*MitchMessage_OrderExecutedWithPriceSize, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x626b5ac2 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_OrderExecutedWithPriceSize. Expected 0x626b5ac2, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_OrderExecutedWithPriceSize, n, err := Read_MitchMessage_OrderExecutedWithPriceSize(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_OrderExecutedWithPriceSize, byteCount, nil
}

// MitchMessage_OrderExecutedWithPriceSize TypeCode
func TypeCode_MitchMessage_OrderExecutedWithPriceSize() uint32 {
	return 0x626b5ac2
}

// MitchMessage_OrderExecutedWithPriceSize IsTypeCode
func IsTypeCode_MitchMessage_OrderExecutedWithPriceSize(typeCode uint32) bool {
	return typeCode == 0x626b5ac2
}

// MitchMessage_Trade Declaration TypeCode: 0x61575a1d
type MitchMessage_Trade struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	ExecutedQuantity Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	Price Streams.MitchPrice08
	TradeID Streams.MitchUInt64
	SubBook Streams.MitchUInt08
	Flags Streams.MitchBitField
	TradeSubType Streams.MitchAlpha
	LastOptPx Streams.MitchPrice08
	Volatility Streams.MitchPrice08
	UnderlyingReferencePrice Streams.MitchPrice08
}

func NewMitchMessage_Trade()*MitchMessage_Trade {
	return &MitchMessage_Trade{}
}

// MitchMessage_Trade writer 
func Write_MitchMessage_Trade(stream Streams.IMitchWriter, value *MitchMessage_Trade) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt08(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchBitField(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.TradeSubType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_Trade reader
func Read_MitchMessage_Trade(stream Streams.IStreamReader) (value *MitchMessage_Trade, byteCount int, err error) {
	value = NewMitchMessage_Trade()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ExecutedQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_MitchUInt08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_MitchBitField()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeSubType, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LastOptPx, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volatility, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.UnderlyingReferencePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_Trade WriteMessage 
func WriteMessage_MitchMessage_Trade(stream Streams.IMitchWriter, value *MitchMessage_Trade) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x61575a1d)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_Trade(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_Trade reader
func ReadMessage_MitchMessage_Trade(stream Streams.IMitchReader) (*MitchMessage_Trade, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x61575a1d {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_Trade. Expected 0x61575a1d, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_Trade, n, err := Read_MitchMessage_Trade(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_Trade, byteCount, nil
}

// MitchMessage_Trade TypeCode
func TypeCode_MitchMessage_Trade() uint32 {
	return 0x61575a1d
}

// MitchMessage_Trade IsTypeCode
func IsTypeCode_MitchMessage_Trade(typeCode uint32) bool {
	return typeCode == 0x61575a1d
}

// MitchMessage_AuctionTrade Declaration TypeCode: 0x0989985f
type MitchMessage_AuctionTrade struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	Quantity Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	Price Streams.MitchPrice08
	TradeID Streams.MitchUInt64
	AuctionType Streams.MitchByte
	LastOptPx Streams.MitchPrice08
	Volatility Streams.MitchPrice08
	UnderlyingReferencePrice Streams.MitchPrice08
}

func NewMitchMessage_AuctionTrade()*MitchMessage_AuctionTrade {
	return &MitchMessage_AuctionTrade{}
}

// MitchMessage_AuctionTrade writer 
func Write_MitchMessage_AuctionTrade(stream Streams.IMitchWriter, value *MitchMessage_AuctionTrade) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.AuctionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AuctionTrade reader
func Read_MitchMessage_AuctionTrade(stream Streams.IStreamReader) (value *MitchMessage_AuctionTrade, byteCount int, err error) {
	value = NewMitchMessage_AuctionTrade()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Quantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.AuctionType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LastOptPx, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volatility, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.UnderlyingReferencePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_AuctionTrade WriteMessage 
func WriteMessage_MitchMessage_AuctionTrade(stream Streams.IMitchWriter, value *MitchMessage_AuctionTrade) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x0989985f)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_AuctionTrade(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AuctionTrade reader
func ReadMessage_MitchMessage_AuctionTrade(stream Streams.IMitchReader) (*MitchMessage_AuctionTrade, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x0989985f {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_AuctionTrade. Expected 0x0989985f, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_AuctionTrade, n, err := Read_MitchMessage_AuctionTrade(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_AuctionTrade, byteCount, nil
}

// MitchMessage_AuctionTrade TypeCode
func TypeCode_MitchMessage_AuctionTrade() uint32 {
	return 0x0989985f
}

// MitchMessage_AuctionTrade IsTypeCode
func IsTypeCode_MitchMessage_AuctionTrade(typeCode uint32) bool {
	return typeCode == 0x0989985f
}

// MitchMessage_OffBookTrade Declaration TypeCode: 0xb9417126
type MitchMessage_OffBookTrade struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	ExecutedQuantity Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	Price Streams.MitchPrice08
	TradeID Streams.MitchUInt64
	OffBookTradeType Streams.MitchAlpha
	TradeTime Streams.MitchTime
	TradeDate Streams.MitchDate
	LastOptPx Streams.MitchPrice08
	Volatility Streams.MitchPrice08
	UnderlyingReferencePrice Streams.MitchPrice08
}

func NewMitchMessage_OffBookTrade()*MitchMessage_OffBookTrade {
	return &MitchMessage_OffBookTrade{}
}

// MitchMessage_OffBookTrade writer 
func Write_MitchMessage_OffBookTrade(stream Streams.IMitchWriter, value *MitchMessage_OffBookTrade) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.OffBookTradeType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchTime(value.TradeTime)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchDate(value.TradeDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OffBookTrade reader
func Read_MitchMessage_OffBookTrade(stream Streams.IStreamReader) (value *MitchMessage_OffBookTrade, byteCount int, err error) {
	value = NewMitchMessage_OffBookTrade()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ExecutedQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OffBookTradeType, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeTime, n, err = stream.Read_MitchTime()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeDate, n, err = stream.Read_MitchDate()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LastOptPx, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volatility, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.UnderlyingReferencePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_OffBookTrade WriteMessage 
func WriteMessage_MitchMessage_OffBookTrade(stream Streams.IMitchWriter, value *MitchMessage_OffBookTrade) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xb9417126)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_OffBookTrade(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_OffBookTrade reader
func ReadMessage_MitchMessage_OffBookTrade(stream Streams.IMitchReader) (*MitchMessage_OffBookTrade, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xb9417126 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_OffBookTrade. Expected 0xb9417126, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_OffBookTrade, n, err := Read_MitchMessage_OffBookTrade(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_OffBookTrade, byteCount, nil
}

// MitchMessage_OffBookTrade TypeCode
func TypeCode_MitchMessage_OffBookTrade() uint32 {
	return 0xb9417126
}

// MitchMessage_OffBookTrade IsTypeCode
func IsTypeCode_MitchMessage_OffBookTrade(typeCode uint32) bool {
	return typeCode == 0xb9417126
}

// MitchMessage_TradeBreak Declaration TypeCode: 0xc5e29cc6
type MitchMessage_TradeBreak struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	TradeID Streams.MitchUInt64
	TradeType Streams.MitchByte
}

func NewMitchMessage_TradeBreak()*MitchMessage_TradeBreak {
	return &MitchMessage_TradeBreak{}
}

// MitchMessage_TradeBreak writer 
func Write_MitchMessage_TradeBreak(stream Streams.IMitchWriter, value *MitchMessage_TradeBreak) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.TradeType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_TradeBreak reader
func Read_MitchMessage_TradeBreak(stream Streams.IStreamReader) (value *MitchMessage_TradeBreak, byteCount int, err error) {
	value = NewMitchMessage_TradeBreak()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_TradeBreak WriteMessage 
func WriteMessage_MitchMessage_TradeBreak(stream Streams.IMitchWriter, value *MitchMessage_TradeBreak) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xc5e29cc6)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_TradeBreak(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_TradeBreak reader
func ReadMessage_MitchMessage_TradeBreak(stream Streams.IMitchReader) (*MitchMessage_TradeBreak, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xc5e29cc6 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_TradeBreak. Expected 0xc5e29cc6, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_TradeBreak, n, err := Read_MitchMessage_TradeBreak(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_TradeBreak, byteCount, nil
}

// MitchMessage_TradeBreak TypeCode
func TypeCode_MitchMessage_TradeBreak() uint32 {
	return 0xc5e29cc6
}

// MitchMessage_TradeBreak IsTypeCode
func IsTypeCode_MitchMessage_TradeBreak(typeCode uint32) bool {
	return typeCode == 0xc5e29cc6
}

// MitchMessage_RecoveryTrade Declaration TypeCode: 0xf7238875
type MitchMessage_RecoveryTrade struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	ExecutedQuantity Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	Price Streams.MitchPrice08
	TradeID Streams.MitchUInt64
	AuctionType Streams.MitchByte
	OffBookRFQTradeType Streams.MitchAlpha
	TradeTime Streams.MitchTime
	TradeDate Streams.MitchDate
	ActionType Streams.MitchByte
	SubBook Streams.MitchUInt08
	Flags Streams.MitchBitField
	LastOptPx Streams.MitchPrice08
	Volatility Streams.MitchPrice08
	UnderlyingReferencePrice Streams.MitchPrice08
}

func NewMitchMessage_RecoveryTrade()*MitchMessage_RecoveryTrade {
	return &MitchMessage_RecoveryTrade{}
}

// MitchMessage_RecoveryTrade writer 
func Write_MitchMessage_RecoveryTrade(stream Streams.IMitchWriter, value *MitchMessage_RecoveryTrade) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.AuctionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.OffBookRFQTradeType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchTime(value.TradeTime)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchDate(value.TradeDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.ActionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt08(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchBitField(value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_RecoveryTrade reader
func Read_MitchMessage_RecoveryTrade(stream Streams.IStreamReader) (value *MitchMessage_RecoveryTrade, byteCount int, err error) {
	value = NewMitchMessage_RecoveryTrade()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ExecutedQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeID, n, err = stream.Read_MitchUInt64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.AuctionType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OffBookRFQTradeType, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeTime, n, err = stream.Read_MitchTime()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.TradeDate, n, err = stream.Read_MitchDate()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ActionType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_MitchUInt08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Flags, n, err = stream.Read_MitchBitField()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LastOptPx, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volatility, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.UnderlyingReferencePrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_RecoveryTrade WriteMessage 
func WriteMessage_MitchMessage_RecoveryTrade(stream Streams.IMitchWriter, value *MitchMessage_RecoveryTrade) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xf7238875)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_RecoveryTrade(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_RecoveryTrade reader
func ReadMessage_MitchMessage_RecoveryTrade(stream Streams.IMitchReader) (*MitchMessage_RecoveryTrade, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xf7238875 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_RecoveryTrade. Expected 0xf7238875, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_RecoveryTrade, n, err := Read_MitchMessage_RecoveryTrade(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_RecoveryTrade, byteCount, nil
}

// MitchMessage_RecoveryTrade TypeCode
func TypeCode_MitchMessage_RecoveryTrade() uint32 {
	return 0xf7238875
}

// MitchMessage_RecoveryTrade IsTypeCode
func IsTypeCode_MitchMessage_RecoveryTrade(typeCode uint32) bool {
	return typeCode == 0xf7238875
}

// MitchMessage_AuctionInfo Declaration TypeCode: 0x9c43b6c7
type MitchMessage_AuctionInfo struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	PairedQuantity Streams.MitchUInt32
	Reserved01 Streams.MitchUInt32
	ImbalanceDirection Streams.MitchByte
	InstrumentID Streams.MitchUInt32
	Reserved02 Streams.MitchByte
	Reserved03 Streams.MitchByte
	Price Streams.MitchPrice08
	AuctionType Streams.MitchByte
}

func NewMitchMessage_AuctionInfo()*MitchMessage_AuctionInfo {
	return &MitchMessage_AuctionInfo{}
}

// MitchMessage_AuctionInfo writer 
func Write_MitchMessage_AuctionInfo(stream Streams.IMitchWriter, value *MitchMessage_AuctionInfo) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.PairedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.ImbalanceDirection)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved03)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.AuctionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AuctionInfo reader
func Read_MitchMessage_AuctionInfo(stream Streams.IStreamReader) (value *MitchMessage_AuctionInfo, byteCount int, err error) {
	value = NewMitchMessage_AuctionInfo()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.PairedQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ImbalanceDirection, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved03, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.AuctionType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_AuctionInfo WriteMessage 
func WriteMessage_MitchMessage_AuctionInfo(stream Streams.IMitchWriter, value *MitchMessage_AuctionInfo) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x9c43b6c7)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_AuctionInfo(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_AuctionInfo reader
func ReadMessage_MitchMessage_AuctionInfo(stream Streams.IMitchReader) (*MitchMessage_AuctionInfo, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x9c43b6c7 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_AuctionInfo. Expected 0x9c43b6c7, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_AuctionInfo, n, err := Read_MitchMessage_AuctionInfo(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_AuctionInfo, byteCount, nil
}

// MitchMessage_AuctionInfo TypeCode
func TypeCode_MitchMessage_AuctionInfo() uint32 {
	return 0x9c43b6c7
}

// MitchMessage_AuctionInfo IsTypeCode
func IsTypeCode_MitchMessage_AuctionInfo(typeCode uint32) bool {
	return typeCode == 0x9c43b6c7
}

// MitchMessage_Statistics Declaration TypeCode: 0xfc258b85
type MitchMessage_Statistics struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	Reserved01 Streams.MitchByte
	Reserved02 Streams.MitchByte
	StatisticType Streams.MitchAlpha
	Price Streams.MitchPrice08
	OpenCloseIndicator Streams.MitchAlpha
	SubBook Streams.MitchUInt08
}

func NewMitchMessage_Statistics()*MitchMessage_Statistics {
	return &MitchMessage_Statistics{}
}

// MitchMessage_Statistics writer 
func Write_MitchMessage_Statistics(stream Streams.IMitchWriter, value *MitchMessage_Statistics) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.StatisticType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.OpenCloseIndicator)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt08(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_Statistics reader
func Read_MitchMessage_Statistics(stream Streams.IStreamReader) (value *MitchMessage_Statistics, byteCount int, err error) {
	value = NewMitchMessage_Statistics()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.StatisticType, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OpenCloseIndicator, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_MitchUInt08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_Statistics WriteMessage 
func WriteMessage_MitchMessage_Statistics(stream Streams.IMitchWriter, value *MitchMessage_Statistics) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xfc258b85)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_Statistics(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_Statistics reader
func ReadMessage_MitchMessage_Statistics(stream Streams.IMitchReader) (*MitchMessage_Statistics, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xfc258b85 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_Statistics. Expected 0xfc258b85, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_Statistics, n, err := Read_MitchMessage_Statistics(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_Statistics, byteCount, nil
}

// MitchMessage_Statistics TypeCode
func TypeCode_MitchMessage_Statistics() uint32 {
	return 0xfc258b85
}

// MitchMessage_Statistics IsTypeCode
func IsTypeCode_MitchMessage_Statistics(typeCode uint32) bool {
	return typeCode == 0xfc258b85
}

// MitchMessage_ExtendedStatistics Declaration TypeCode: 0x5953a419
type MitchMessage_ExtendedStatistics struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	InstrumentID Streams.MitchUInt32
	HighPrice Streams.MitchPrice08
	LowPrice Streams.MitchPrice08
	VWAP Streams.MitchPrice08
	Volume Streams.MitchUInt32
	Turnover Streams.MitchPrice04
	NumberofTrades Streams.MitchUInt32
	Reserved01 Streams.MitchAlpha
	SubBook Streams.MitchUInt08
	NotionalExposure Streams.MitchPrice08
	NotionalDeltaExposure Streams.MitchPrice08
	OpenInterest Streams.MitchPrice08
}

func NewMitchMessage_ExtendedStatistics()*MitchMessage_ExtendedStatistics {
	return &MitchMessage_ExtendedStatistics{}
}

// MitchMessage_ExtendedStatistics writer 
func Write_MitchMessage_ExtendedStatistics(stream Streams.IMitchWriter, value *MitchMessage_ExtendedStatistics) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.HighPrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.LowPrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.VWAP)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Volume)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice04(value.Turnover)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.NumberofTrades)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt08(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.NotionalExposure)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.NotionalDeltaExposure)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.OpenInterest)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_ExtendedStatistics reader
func Read_MitchMessage_ExtendedStatistics(stream Streams.IStreamReader) (value *MitchMessage_ExtendedStatistics, byteCount int, err error) {
	value = NewMitchMessage_ExtendedStatistics()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.InstrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.HighPrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.LowPrice, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.VWAP, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Volume, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Turnover, n, err = stream.Read_MitchPrice04()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.NumberofTrades, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_MitchUInt08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.NotionalExposure, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.NotionalDeltaExposure, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.OpenInterest, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_ExtendedStatistics WriteMessage 
func WriteMessage_MitchMessage_ExtendedStatistics(stream Streams.IMitchWriter, value *MitchMessage_ExtendedStatistics) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x5953a419)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_ExtendedStatistics(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_ExtendedStatistics reader
func ReadMessage_MitchMessage_ExtendedStatistics(stream Streams.IMitchReader) (*MitchMessage_ExtendedStatistics, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x5953a419 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_ExtendedStatistics. Expected 0x5953a419, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_ExtendedStatistics, n, err := Read_MitchMessage_ExtendedStatistics(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_ExtendedStatistics, byteCount, nil
}

// MitchMessage_ExtendedStatistics TypeCode
func TypeCode_MitchMessage_ExtendedStatistics() uint32 {
	return 0x5953a419
}

// MitchMessage_ExtendedStatistics IsTypeCode
func IsTypeCode_MitchMessage_ExtendedStatistics(typeCode uint32) bool {
	return typeCode == 0x5953a419
}

// MitchMessage_News Declaration TypeCode: 0x947caf09
type MitchMessage_News struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	Time Streams.MitchTime
	Urgency Streams.MitchByte
	Headline Streams.MitchAlpha
	Text Streams.MitchAlpha
	Instruments Streams.MitchAlpha
	Underlyings Streams.MitchAlpha
}

func NewMitchMessage_News()*MitchMessage_News {
	return &MitchMessage_News{}
}

// MitchMessage_News writer 
func Write_MitchMessage_News(stream Streams.IMitchWriter, value *MitchMessage_News) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchTime(value.Time)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Urgency)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Headline)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Text)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Instruments)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchAlpha(value.Underlyings)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_News reader
func Read_MitchMessage_News(stream Streams.IStreamReader) (value *MitchMessage_News, byteCount int, err error) {
	value = NewMitchMessage_News()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Time, n, err = stream.Read_MitchTime()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Urgency, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Headline, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Text, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Instruments, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Underlyings, n, err = stream.Read_MitchAlpha()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_News WriteMessage 
func WriteMessage_MitchMessage_News(stream Streams.IMitchWriter, value *MitchMessage_News) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x947caf09)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_News(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_News reader
func ReadMessage_MitchMessage_News(stream Streams.IMitchReader) (*MitchMessage_News, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x947caf09 {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_News. Expected 0x947caf09, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_News, n, err := Read_MitchMessage_News(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_News, byteCount, nil
}

// MitchMessage_News TypeCode
func TypeCode_MitchMessage_News() uint32 {
	return 0x947caf09
}

// MitchMessage_News IsTypeCode
func IsTypeCode_MitchMessage_News(typeCode uint32) bool {
	return typeCode == 0x947caf09
}

// MitchMessage_TopOfBook Declaration TypeCode: 0xc036801e
type MitchMessage_TopOfBook struct {
	Length Streams.MitchUInt16
	MessageType Streams.MitchByte
	Nanosecond Streams.MitchUInt32
	instrumentID Streams.MitchUInt32
	ReserveField Streams.MitchUInt16
	SubBook Streams.MitchBitField
	Action Streams.MitchByte
	Side Streams.MitchByte
	Price Streams.MitchPrice08
	Quantity Streams.MitchUInt32
	MarketOrderQuantity Streams.MitchUInt32
	Reserved01 Streams.MitchUInt32
	Reserved02 Streams.MitchUInt32
}

func NewMitchMessage_TopOfBook()*MitchMessage_TopOfBook {
	return &MitchMessage_TopOfBook{}
}

// MitchMessage_TopOfBook writer 
func Write_MitchMessage_TopOfBook(stream Streams.IMitchWriter, value *MitchMessage_TopOfBook) (byteCount int, err error) {
	var n int 
	n, err = stream.Write_MitchUInt16(value.Length)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.MessageType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.instrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt16(value.ReserveField)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchBitField(value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Action)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchByte(value.Side)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchPrice08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.MarketOrderQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = stream.Write_MitchUInt32(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_TopOfBook reader
func Read_MitchMessage_TopOfBook(stream Streams.IStreamReader) (value *MitchMessage_TopOfBook, byteCount int, err error) {
	value = NewMitchMessage_TopOfBook()
	var n int 
	value.Length, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Nanosecond, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.instrumentID, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.ReserveField, n, err = stream.Read_MitchUInt16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.SubBook, n, err = stream.Read_MitchBitField()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Action, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Side, n, err = stream.Read_MitchByte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Price, n, err = stream.Read_MitchPrice08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Quantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MarketOrderQuantity, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved01, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Reserved02, n, err = stream.Read_MitchUInt32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// MitchMessage_TopOfBook WriteMessage 
func WriteMessage_MitchMessage_TopOfBook(stream Streams.IMitchWriter, value *MitchMessage_TopOfBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xc036801e)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_MitchMessage_TopOfBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// MitchMessage_TopOfBook reader
func ReadMessage_MitchMessage_TopOfBook(stream Streams.IMitchReader) (*MitchMessage_TopOfBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xc036801e {
		return nil, 0, errors.New(fmt.Sprintf("typecode mismatch, while reading MitchMessage_TopOfBook. Expected 0xc036801e, got 0x%08x", typeCode))
	}
	byteCount += n
	value_MitchMessage_TopOfBook, n, err := Read_MitchMessage_TopOfBook(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_MitchMessage_TopOfBook, byteCount, nil
}

// MitchMessage_TopOfBook TypeCode
func TypeCode_MitchMessage_TopOfBook() uint32 {
	return 0xc036801e
}

// MitchMessage_TopOfBook IsTypeCode
func IsTypeCode_MitchMessage_TopOfBook(typeCode uint32) bool {
	return typeCode == 0xc036801e
}

