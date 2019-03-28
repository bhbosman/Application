package GeneratedFiles

import (
	"fmt"
	"time"
)
import "github.com/bhbosman/Application/Streams"

// Declared typed
// TimeMessage
// TEventCode
// SystemEventMessage
// SymbolDirectorySymbolStatus
// SymbolDirectoryOptionType
// SymbolDirectoryFlags
// SymbolDirectorySubBook
// SymbolDirectorySettlementMethod
// SymbolDirectoryMessage
// SymbolStatusTradingStatus
// SymbolStatusSessionChangeReason
// SymbolStatusBookType
// SymbolStatusMessageFlags
// SymbolStatusMessage
// OrderSide
// AddOrderFlags
// AddOrderMessage
// AddAttributedOrderFlags
// AddAttributedOrderMessage
// OrderDeletedMessage
// OrderModifiedFlags
// OrderModifiedMessage
// OrderBookClearSubBook
// OrderBookClearBookType
// OrderBookClearMessage
// OrderExecutedMessage
// OrderExecutedWithPriceSizeMessage
// TradeMessageFlags
// TradeMessageSubBook
// TradeMessage
// AuctionTradeMessage
// OffBookTradeMessage
// TradeBreakMessage
// RecoveryTradeMessageFlags
// RecoveryTradeMessageSubBook
// RecoveryTradeMessage
// AuctionInfoMessage
// TStatisticsMessageOpenCloseIndicator
// TStatisticsMessageSubBook
// StatisticsMessage
// ExtendedStatisticsMessage
// NewsMessage
// TopOfBookMessageSubBook
// TopOfBookMessage
//
// TimeMessage Declaration TypeCode: 0xafbd7a0e
type TimeMessage struct {
	Length      uint16 `json:"Length" xml:"Length,attr"`
	MessageType byte   `json:"MessageType" xml:"MessageType,attr"`
	Seconds     uint32 `json:"Seconds" xml:"Seconds,attr"`
}

func NewTimeMessage() *TimeMessage {
	return &TimeMessage{
		Length:      7,
		MessageType: 84,
	}
}

// TimeMessage writer
func Write_TimeMessage(stream Streams.IMitchWriter, value *TimeMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(7)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(84)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Seconds, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Seconds)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TimeMessage reader
func Read_TimeMessage(stream Streams.IMitchReader) (value *TimeMessage, byteCount int, err error) {
	value = NewTimeMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 7
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x54 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message TimeMessage was expected 0x54, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Seconds, Type: uint32
	//
	//
	value.Seconds, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// TimeMessage reset
func Reset_TimeMessage(value *TimeMessage) {
	value.Length = 7
	value.MessageType = 84
	value.Seconds = 0
}

// TimeMessage clone
func Clone_TimeMessage(value *TimeMessage) *TimeMessage {
	result := NewTimeMessage()
	result.Seconds = value.Seconds
	return result
}

// TimeMessage WriteMessage
func WriteMessage_TimeMessage(stream Streams.IMitchWriter, value *TimeMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xafbd7a0e)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TimeMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TimeMessage reader
func ReadMessage_TimeMessage(stream Streams.IMitchReader) (*TimeMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xafbd7a0e {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading TimeMessage. Expected 0xafbd7a0e, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TimeMessage, n, err := Read_TimeMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_TimeMessage, byteCount, nil
}

// TimeMessage TypeCode
func TypeCode_TimeMessage() uint32 {
	return 0xafbd7a0e
}

// TimeMessage IsTypeCode
func IsTypeCode_TimeMessage(typeCode uint32) bool {
	return typeCode == 0xafbd7a0e
}

type TEventCode byte

//noinspection ALL
const (
	TEventCode_EndOfDay   = 'C' // default value: byte('C')
	TEventCode_StartOfDay = 'O' // default value: byte('O')
)

// TEventCode Declaration TypeCode: 0x2ff833d1
func NewTEventCode() TEventCode {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading TEventCode. Expected 0x2ff833d1, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TEventCode, n, err := Read_TEventCode(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_TEventCode, byteCount, nil
}

// TEventCode writer
func Write_TEventCode(stream Streams.IMitchWriter, value TEventCode) (int, error) {
	return stream.Write_byte(byte(value))
}

// TEventCode reader
func Read_TEventCode(stream Streams.IMitchReader) (value TEventCode, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return TEventCode(v), b, e
}

// SystemEventMessage Declaration TypeCode: 0x27b7b6f7
type SystemEventMessage struct {
	Length      uint16     `json:"Length" xml:"Length,attr"`
	MessageType byte       `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond  uint32     `json:"Nanosecond" xml:"Nanosecond,attr"`
	EventCode   TEventCode `json:"EventCode" xml:"EventCode,attr"`
}

func NewSystemEventMessage() *SystemEventMessage {
	return &SystemEventMessage{
		Length:      8,
		MessageType: 83,
	}
}

// SystemEventMessage writer
func Write_SystemEventMessage(stream Streams.IMitchWriter, value *SystemEventMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(8)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(83)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: EventCode, Type: TEventCode
	//
	//
	n, err = stream.Write_byte(byte(value.EventCode))
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SystemEventMessage reader
func Read_SystemEventMessage(stream Streams.IMitchReader) (value *SystemEventMessage, byteCount int, err error) {
	value = NewSystemEventMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 8
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x53 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message SystemEventMessage was expected 0x53, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: EventCode, Type: TEventCode
	//
	//
	temp_EventCode, n, err := stream.Read_byte()
	value.EventCode = TEventCode(temp_EventCode)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// SystemEventMessage reset
func Reset_SystemEventMessage(value *SystemEventMessage) {
	value.Length = 8
	value.MessageType = 83
	value.Nanosecond = 0
	value.EventCode = 0
}

// SystemEventMessage clone
func Clone_SystemEventMessage(value *SystemEventMessage) *SystemEventMessage {
	result := NewSystemEventMessage()
	result.Nanosecond = value.Nanosecond
	result.EventCode = value.EventCode
	return result
}

// SystemEventMessage WriteMessage
func WriteMessage_SystemEventMessage(stream Streams.IMitchWriter, value *SystemEventMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x27b7b6f7)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SystemEventMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SystemEventMessage reader
func ReadMessage_SystemEventMessage(stream Streams.IMitchReader) (*SystemEventMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x27b7b6f7 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading SystemEventMessage. Expected 0x27b7b6f7, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SystemEventMessage, n, err := Read_SystemEventMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_SystemEventMessage, byteCount, nil
}

// SystemEventMessage TypeCode
func TypeCode_SystemEventMessage() uint32 {
	return 0x27b7b6f7
}

// SystemEventMessage IsTypeCode
func IsTypeCode_SystemEventMessage(typeCode uint32) bool {
	return typeCode == 0x27b7b6f7
}

type SymbolDirectorySymbolStatus byte

//noinspection ALL
const (
	SymbolDirectorySymbolStatus_Halted    = 'H' // default value: byte('H')
	SymbolDirectorySymbolStatus_Suspended = 'S' // default value: byte('S')
	SymbolDirectorySymbolStatus_Inactive  = 'a' // default value: byte('a')
)

// SymbolDirectorySymbolStatus Declaration TypeCode: 0x810d8fde
func NewSymbolDirectorySymbolStatus() SymbolDirectorySymbolStatus {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading SymbolDirectorySymbolStatus. Expected 0x810d8fde, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolDirectorySymbolStatus, n, err := Read_SymbolDirectorySymbolStatus(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolDirectorySymbolStatus, byteCount, nil
}

// SymbolDirectorySymbolStatus writer
func Write_SymbolDirectorySymbolStatus(stream Streams.IMitchWriter, value SymbolDirectorySymbolStatus) (int, error) {
	return stream.Write_byte(byte(value))
}

// SymbolDirectorySymbolStatus reader
func Read_SymbolDirectorySymbolStatus(stream Streams.IMitchReader) (value SymbolDirectorySymbolStatus, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return SymbolDirectorySymbolStatus(v), b, e
}

type SymbolDirectoryOptionType byte

//noinspection ALL
const (
	SymbolDirectoryOptionType_CallOption = 'C' // default value: byte('C')
	SymbolDirectoryOptionType_PutOption  = 'P' // default value: byte('P')
)

// SymbolDirectoryOptionType Declaration TypeCode: 0x8ca314f2
func NewSymbolDirectoryOptionType() SymbolDirectoryOptionType {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading SymbolDirectoryOptionType. Expected 0x8ca314f2, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolDirectoryOptionType, n, err := Read_SymbolDirectoryOptionType(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolDirectoryOptionType, byteCount, nil
}

// SymbolDirectoryOptionType writer
func Write_SymbolDirectoryOptionType(stream Streams.IMitchWriter, value SymbolDirectoryOptionType) (int, error) {
	return stream.Write_byte(byte(value))
}

// SymbolDirectoryOptionType reader
func Read_SymbolDirectoryOptionType(stream Streams.IMitchReader) (value SymbolDirectoryOptionType, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return SymbolDirectoryOptionType(v), b, e
}

type SymbolDirectoryFlags struct {
	Flags byte
}

func NewSymbolDirectoryFlags() SymbolDirectoryFlags {
	return SymbolDirectoryFlags{}
}

// SymbolDirectoryFlags Declaration TypeCode: 0xcc630022
// SymbolDirectoryFlags WriteMessage
func WriteMessage_SymbolDirectoryFlags(stream Streams.IMitchWriter, value SymbolDirectoryFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xcc630022)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolDirectoryFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectoryFlags reader
func ReadMessage_SymbolDirectoryFlags(stream Streams.IMitchReader) (SymbolDirectoryFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return SymbolDirectoryFlags{}, 0, err
	}
	if typeCode != 0xcc630022 {
		return SymbolDirectoryFlags{}, 0, fmt.Errorf("typecode mismatch, while reading SymbolDirectoryFlags. Expected 0xcc630022, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolDirectoryFlags, n, err := Read_SymbolDirectoryFlags(stream)
	if err != nil {
		return SymbolDirectoryFlags{}, 0, err
	}
	byteCount += n
	return value_SymbolDirectoryFlags, byteCount, nil
}

// SymbolDirectoryFlags writer
func Write_SymbolDirectoryFlags(stream Streams.IMitchWriter, value SymbolDirectoryFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// SymbolDirectoryFlags reader
func Read_SymbolDirectoryFlags(stream Streams.IMitchReader) (value SymbolDirectoryFlags, byteCount int, err error) {
	value = NewSymbolDirectoryFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

type SymbolDirectorySubBook struct {
	Flags byte
}

func NewSymbolDirectorySubBook() SymbolDirectorySubBook {
	return SymbolDirectorySubBook{}
}

// SymbolDirectorySubBook Declaration TypeCode: 0x03b3c050
// SymbolDirectorySubBook WriteMessage
func WriteMessage_SymbolDirectorySubBook(stream Streams.IMitchWriter, value SymbolDirectorySubBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x03b3c050)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolDirectorySubBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectorySubBook reader
func ReadMessage_SymbolDirectorySubBook(stream Streams.IMitchReader) (SymbolDirectorySubBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return SymbolDirectorySubBook{}, 0, err
	}
	if typeCode != 0x03b3c050 {
		return SymbolDirectorySubBook{}, 0, fmt.Errorf("typecode mismatch, while reading SymbolDirectorySubBook. Expected 0x03b3c050, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolDirectorySubBook, n, err := Read_SymbolDirectorySubBook(stream)
	if err != nil {
		return SymbolDirectorySubBook{}, 0, err
	}
	byteCount += n
	return value_SymbolDirectorySubBook, byteCount, nil
}

// SymbolDirectorySubBook writer
func Write_SymbolDirectorySubBook(stream Streams.IMitchWriter, value SymbolDirectorySubBook) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// SymbolDirectorySubBook reader
func Read_SymbolDirectorySubBook(stream Streams.IMitchReader) (value SymbolDirectorySubBook, byteCount int, err error) {
	value = NewSymbolDirectorySubBook()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

type SymbolDirectorySettlementMethod byte

//noinspection ALL
const (
	SymbolDirectorySettlementMethod_Cash     = 'C' // default value: byte('C')
	SymbolDirectorySettlementMethod_Physical = 'P' // default value: byte('P')
)

// SymbolDirectorySettlementMethod Declaration TypeCode: 0xc1ac87fd
func NewSymbolDirectorySettlementMethod() SymbolDirectorySettlementMethod {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading SymbolDirectorySettlementMethod. Expected 0xc1ac87fd, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolDirectorySettlementMethod, n, err := Read_SymbolDirectorySettlementMethod(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolDirectorySettlementMethod, byteCount, nil
}

// SymbolDirectorySettlementMethod writer
func Write_SymbolDirectorySettlementMethod(stream Streams.IMitchWriter, value SymbolDirectorySettlementMethod) (int, error) {
	return stream.Write_byte(byte(value))
}

// SymbolDirectorySettlementMethod reader
func Read_SymbolDirectorySettlementMethod(stream Streams.IMitchReader) (value SymbolDirectorySettlementMethod, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return SymbolDirectorySettlementMethod(v), b, e
}

// SymbolDirectoryMessage Declaration TypeCode: 0xcd19edf6
type SymbolDirectoryMessage struct {
	Length                uint16                          `json:"Length" xml:"Length,attr"`
	MessageType           byte                            `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond            uint32                          `json:"Nanosecond" xml:"Nanosecond,attr"`
	InstrumentID          uint32                          `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01            byte                            `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02            byte                            `json:"Reserved02" xml:"Reserved02,attr"`
	SymbolStatus          SymbolDirectorySymbolStatus     `json:"SymbolStatus" xml:"SymbolStatus,attr"`
	Isin                  string                          `json:"Isin" xml:"Isin,attr"`
	Symbol                string                          `json:"Symbol" xml:"Symbol,attr"`
	Tidm                  string                          `json:"Tidm" xml:"Tidm,attr"`
	Segment               string                          `json:"Segment" xml:"Segment,attr"`
	PreviousClosePrice    float64                         `json:"PreviousClosePrice" xml:"PreviousClosePrice,attr"`
	ExpirationDate        time.Time                       `json:"ExpirationDate" xml:"ExpirationDate,attr"`
	Underlying            string                          `json:"Underlying" xml:"Underlying,attr"`
	StrikePrice           float64                         `json:"StrikePrice" xml:"StrikePrice,attr"`
	OptionType            SymbolDirectoryOptionType       `json:"OptionType" xml:"OptionType,attr"`
	Issuer                string                          `json:"Issuer" xml:"Issuer,attr"`
	IssueDate             time.Time                       `json:"IssueDate" xml:"IssueDate,attr"`
	Coupon                float64                         `json:"Coupon" xml:"Coupon,attr"`
	Flags                 SymbolDirectoryFlags            `json:"Flags" xml:"Flags,attr"`
	SubBook               SymbolDirectorySubBook          `json:"SubBook" xml:"SubBook,attr"`
	CorporateAction       string                          `json:"CorporateAction" xml:"CorporateAction,attr"`
	Leg1Symbol            string                          `json:"Leg1Symbol" xml:"Leg1Symbol,attr"`
	Leg2Symbol            string                          `json:"Leg2Symbol" xml:"Leg2Symbol,attr"`
	ContractMultiplier    float64                         `json:"ContractMultiplier" xml:"ContractMultiplier,attr"`
	SettlementMethod      SymbolDirectorySettlementMethod `json:"SettlementMethod" xml:"SettlementMethod,attr"`
	InstrumentSubCategory string                          `json:"InstrumentSubCategory" xml:"InstrumentSubCategory,attr"`
}

func NewSymbolDirectoryMessage() *SymbolDirectoryMessage {
	return &SymbolDirectoryMessage{
		Length:      421,
		MessageType: 82,
	}
}

// SymbolDirectoryMessage writer
func Write_SymbolDirectoryMessage(stream Streams.IMitchWriter, value *SymbolDirectoryMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(421)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(82)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: SymbolStatus, Type: SymbolDirectorySymbolStatus
	//
	//
	n, err = stream.Write_byte(byte(value.SymbolStatus))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Isin, Type: string
	//
	//
	n, err = stream.Write_string(value.Isin, 12)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Symbol, Type: string
	//
	//
	n, err = stream.Write_string(value.Symbol, 25)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Tidm, Type: string
	//
	//
	n, err = stream.Write_string(value.Tidm, 12)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Segment, Type: string
	//
	//
	n, err = stream.Write_string(value.Segment, 6)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: PreviousClosePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.PreviousClosePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: ExpirationDate, Type: Time
	//
	//
	n, err = stream.Write_mitch_date(value.ExpirationDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: Underlying, Type: string
	//
	//
	n, err = stream.Write_string(value.Underlying, 25)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: StrikePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.StrikePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 15, Member Name: OptionType, Type: SymbolDirectoryOptionType
	//
	//
	n, err = stream.Write_byte(byte(value.OptionType))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 16, Member Name: Issuer, Type: string
	//
	//
	n, err = stream.Write_string(value.Issuer, 6)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 17, Member Name: IssueDate, Type: Time
	//
	//
	n, err = stream.Write_mitch_date(value.IssueDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 18, Member Name: Coupon, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Coupon)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 19, Member Name: Flags, Type: SymbolDirectoryFlags
	//
	//
	n, err = Write_SymbolDirectoryFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 20, Member Name: SubBook, Type: SymbolDirectorySubBook
	//
	//
	n, err = Write_SymbolDirectorySubBook(stream, value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 21, Member Name: CorporateAction, Type: string
	//
	//
	n, err = stream.Write_string(value.CorporateAction, 189)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 22, Member Name: Leg1Symbol, Type: string
	//
	//
	n, err = stream.Write_string(value.Leg1Symbol, 25)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 23, Member Name: Leg2Symbol, Type: string
	//
	//
	n, err = stream.Write_string(value.Leg2Symbol, 25)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 24, Member Name: ContractMultiplier, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.ContractMultiplier)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 25, Member Name: SettlementMethod, Type: SymbolDirectorySettlementMethod
	//
	//
	n, err = stream.Write_byte(byte(value.SettlementMethod))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 26, Member Name: InstrumentSubCategory, Type: string
	//
	//
	n, err = stream.Write_string(value.InstrumentSubCategory, 30)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectoryMessage reader
func Read_SymbolDirectoryMessage(stream Streams.IMitchReader) (value *SymbolDirectoryMessage, byteCount int, err error) {
	value = NewSymbolDirectoryMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 421
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x52 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message SymbolDirectoryMessage was expected 0x52, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: SymbolStatus, Type: SymbolDirectorySymbolStatus
	//
	//
	temp_SymbolStatus, n, err := stream.Read_byte()
	value.SymbolStatus = SymbolDirectorySymbolStatus(temp_SymbolStatus)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Isin, Type: string
	//
	//
	value.Isin, n, err = stream.Read_string(12)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Symbol, Type: string
	//
	//
	value.Symbol, n, err = stream.Read_string(25)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Tidm, Type: string
	//
	//
	value.Tidm, n, err = stream.Read_string(12)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Segment, Type: string
	//
	//
	value.Segment, n, err = stream.Read_string(6)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: PreviousClosePrice, Type: float64
	//
	//
	value.PreviousClosePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: ExpirationDate, Type: Time
	//
	//
	value.ExpirationDate, n, err = stream.Read_mitch_date()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: Underlying, Type: string
	//
	//
	value.Underlying, n, err = stream.Read_string(25)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: StrikePrice, Type: float64
	//
	//
	value.StrikePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 15, Member Name: OptionType, Type: SymbolDirectoryOptionType
	//
	//
	temp_OptionType, n, err := stream.Read_byte()
	value.OptionType = SymbolDirectoryOptionType(temp_OptionType)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 16, Member Name: Issuer, Type: string
	//
	//
	value.Issuer, n, err = stream.Read_string(6)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 17, Member Name: IssueDate, Type: Time
	//
	//
	value.IssueDate, n, err = stream.Read_mitch_date()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 18, Member Name: Coupon, Type: float64
	//
	//
	value.Coupon, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 19, Member Name: Flags, Type: SymbolDirectoryFlags
	//
	//
	value.Flags, n, err = Read_SymbolDirectoryFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 20, Member Name: SubBook, Type: SymbolDirectorySubBook
	//
	//
	value.SubBook, n, err = Read_SymbolDirectorySubBook(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 21, Member Name: CorporateAction, Type: string
	//
	//
	value.CorporateAction, n, err = stream.Read_string(189)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 22, Member Name: Leg1Symbol, Type: string
	//
	//
	value.Leg1Symbol, n, err = stream.Read_string(25)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 23, Member Name: Leg2Symbol, Type: string
	//
	//
	value.Leg2Symbol, n, err = stream.Read_string(25)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 24, Member Name: ContractMultiplier, Type: float64
	//
	//
	value.ContractMultiplier, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 25, Member Name: SettlementMethod, Type: SymbolDirectorySettlementMethod
	//
	//
	temp_SettlementMethod, n, err := stream.Read_byte()
	value.SettlementMethod = SymbolDirectorySettlementMethod(temp_SettlementMethod)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 26, Member Name: InstrumentSubCategory, Type: string
	//
	//
	value.InstrumentSubCategory, n, err = stream.Read_string(30)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// SymbolDirectoryMessage reset
func Reset_SymbolDirectoryMessage(value *SymbolDirectoryMessage) {
	value.Length = 421
	value.MessageType = 82
	value.Nanosecond = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.SymbolStatus = 0
	value.Isin = ""
	value.Symbol = ""
	value.Tidm = ""
	value.Segment = ""
	value.PreviousClosePrice = 0.0
	value.ExpirationDate = time.Time{}
	value.Underlying = ""
	value.StrikePrice = 0.0
	value.OptionType = 0
	value.Issuer = ""
	value.IssueDate = time.Time{}
	value.Coupon = 0.0
	value.Flags = SymbolDirectoryFlags{}
	value.SubBook = SymbolDirectorySubBook{}
	value.CorporateAction = ""
	value.Leg1Symbol = ""
	value.Leg2Symbol = ""
	value.ContractMultiplier = 0.0
	value.SettlementMethod = 0
	value.InstrumentSubCategory = ""
}

// SymbolDirectoryMessage clone
func Clone_SymbolDirectoryMessage(value *SymbolDirectoryMessage) *SymbolDirectoryMessage {
	result := NewSymbolDirectoryMessage()
	result.Nanosecond = value.Nanosecond
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.SymbolStatus = value.SymbolStatus
	result.Isin = value.Isin
	result.Symbol = value.Symbol
	result.Tidm = value.Tidm
	result.Segment = value.Segment
	result.PreviousClosePrice = value.PreviousClosePrice
	result.ExpirationDate = value.ExpirationDate
	result.Underlying = value.Underlying
	result.StrikePrice = value.StrikePrice
	result.OptionType = value.OptionType
	result.Issuer = value.Issuer
	result.IssueDate = value.IssueDate
	result.Coupon = value.Coupon
	result.Flags = value.Flags
	result.SubBook = value.SubBook
	result.CorporateAction = value.CorporateAction
	result.Leg1Symbol = value.Leg1Symbol
	result.Leg2Symbol = value.Leg2Symbol
	result.ContractMultiplier = value.ContractMultiplier
	result.SettlementMethod = value.SettlementMethod
	result.InstrumentSubCategory = value.InstrumentSubCategory
	return result
}

// SymbolDirectoryMessage WriteMessage
func WriteMessage_SymbolDirectoryMessage(stream Streams.IMitchWriter, value *SymbolDirectoryMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xcd19edf6)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolDirectoryMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolDirectoryMessage reader
func ReadMessage_SymbolDirectoryMessage(stream Streams.IMitchReader) (*SymbolDirectoryMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xcd19edf6 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading SymbolDirectoryMessage. Expected 0xcd19edf6, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolDirectoryMessage, n, err := Read_SymbolDirectoryMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_SymbolDirectoryMessage, byteCount, nil
}

// SymbolDirectoryMessage TypeCode
func TypeCode_SymbolDirectoryMessage() uint32 {
	return 0xcd19edf6
}

// SymbolDirectoryMessage IsTypeCode
func IsTypeCode_SymbolDirectoryMessage(typeCode uint32) bool {
	return typeCode == 0xcd19edf6
}

type SymbolStatusTradingStatus byte

//noinspection ALL
const (
	SymbolStatusTradingStatus_Halt                    = 'H' // default value: byte('H')
	SymbolStatusTradingStatus_RegularTrading          = 'T' // default value: byte('T')
	SymbolStatusTradingStatus_OpeningAuctionCall      = 'a' // default value: byte('a')
	SymbolStatusTradingStatus_PostClose               = 'b' // default value: byte('b')
	SymbolStatusTradingStatus_MarketClose             = 'c' // default value: byte('c')
	SymbolStatusTradingStatus_ClosingAuctionCall      = 'd' // default value: byte('d')
	SymbolStatusTradingStatus_VolatilityAuctionCall   = 'e' // default value: byte('e')
	SymbolStatusTradingStatus_EODVolumeAuctionCall    = 'E' // default value: byte('E')
	SymbolStatusTradingStatus_ReOpeningAuctionCall    = 'f' // default value: byte('f')
	SymbolStatusTradingStatus_Pause                   = 'l' // default value: byte('l')
	SymbolStatusTradingStatus_FuturesCloseOut         = 'p' // default value: byte('p')
	SymbolStatusTradingStatus_ClosingPriceCross       = 's' // default value: byte('s')
	SymbolStatusTradingStatus_IntraDayAuctionCall     = 'u' // default value: byte('u')
	SymbolStatusTradingStatus_EndTradeReporting       = 'v' // default value: byte('v')
	SymbolStatusTradingStatus_NoActiveSession         = 'w' // default value: byte('w')
	SymbolStatusTradingStatus_EndOfPostClose          = 'x' // default value: byte('x')
	SymbolStatusTradingStatus_StarOofTrading          = 'y' // default value: byte('y')
	SymbolStatusTradingStatus_ClosingPricePublication = 'z' // default value: byte('z')
)

// SymbolStatusTradingStatus Declaration TypeCode: 0xdeb92876
func NewSymbolStatusTradingStatus() SymbolStatusTradingStatus {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading SymbolStatusTradingStatus. Expected 0xdeb92876, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolStatusTradingStatus, n, err := Read_SymbolStatusTradingStatus(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolStatusTradingStatus, byteCount, nil
}

// SymbolStatusTradingStatus writer
func Write_SymbolStatusTradingStatus(stream Streams.IMitchWriter, value SymbolStatusTradingStatus) (int, error) {
	return stream.Write_byte(byte(value))
}

// SymbolStatusTradingStatus reader
func Read_SymbolStatusTradingStatus(stream Streams.IMitchReader) (value SymbolStatusTradingStatus, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return SymbolStatusTradingStatus(v), b, e
}

type SymbolStatusSessionChangeReason byte

//noinspection ALL
const (
	SymbolStatusSessionChangeReason_ScheduledTransition   = 0 // default value: byte(0)
	SymbolStatusSessionChangeReason_ExtendedByMarketOps   = 1 // default value: byte(1)
	SymbolStatusSessionChangeReason_ShortenedByMarketOps  = 2 // default value: byte(2)
	SymbolStatusSessionChangeReason_MarketOrderImbalance  = 3 // default value: byte(3)
	SymbolStatusSessionChangeReason_PriceOutsideRange     = 4 // default value: byte(4)
	SymbolStatusSessionChangeReason_CircuitBreakerTripped = 5 // default value: byte(5)
	SymbolStatusSessionChangeReason_Unavailable           = 9 // default value: byte(9)
)

// SymbolStatusSessionChangeReason Declaration TypeCode: 0xfee8af4f
func NewSymbolStatusSessionChangeReason() SymbolStatusSessionChangeReason {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading SymbolStatusSessionChangeReason. Expected 0xfee8af4f, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolStatusSessionChangeReason, n, err := Read_SymbolStatusSessionChangeReason(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolStatusSessionChangeReason, byteCount, nil
}

// SymbolStatusSessionChangeReason writer
func Write_SymbolStatusSessionChangeReason(stream Streams.IMitchWriter, value SymbolStatusSessionChangeReason) (int, error) {
	return stream.Write_byte(byte(value))
}

// SymbolStatusSessionChangeReason reader
func Read_SymbolStatusSessionChangeReason(stream Streams.IMitchReader) (value SymbolStatusSessionChangeReason, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return SymbolStatusSessionChangeReason(v), b, e
}

type SymbolStatusBookType byte

//noinspection ALL
const (
	SymbolStatusBookType_OnBook           = 1  // default value: byte(1)
	SymbolStatusBookType_OffBook          = 2  // default value: byte(2)
	SymbolStatusBookType_BulletinBoard    = 9  // default value: byte(9)
	SymbolStatusBookType_NegotiatedTrades = 11 // default value: byte(11)
)

// SymbolStatusBookType Declaration TypeCode: 0x0b2355a5
func NewSymbolStatusBookType() SymbolStatusBookType {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading SymbolStatusBookType. Expected 0x0b2355a5, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolStatusBookType, n, err := Read_SymbolStatusBookType(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_SymbolStatusBookType, byteCount, nil
}

// SymbolStatusBookType writer
func Write_SymbolStatusBookType(stream Streams.IMitchWriter, value SymbolStatusBookType) (int, error) {
	return stream.Write_byte(byte(value))
}

// SymbolStatusBookType reader
func Read_SymbolStatusBookType(stream Streams.IMitchReader) (value SymbolStatusBookType, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return SymbolStatusBookType(v), b, e
}

type SymbolStatusMessageFlags struct {
	Flags byte
}

func NewSymbolStatusMessageFlags() SymbolStatusMessageFlags {
	return SymbolStatusMessageFlags{}
}

// SymbolStatusMessageFlags Declaration TypeCode: 0x5cadcb8d
// SymbolStatusMessageFlags WriteMessage
func WriteMessage_SymbolStatusMessageFlags(stream Streams.IMitchWriter, value SymbolStatusMessageFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x5cadcb8d)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolStatusMessageFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolStatusMessageFlags reader
func ReadMessage_SymbolStatusMessageFlags(stream Streams.IMitchReader) (SymbolStatusMessageFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return SymbolStatusMessageFlags{}, 0, err
	}
	if typeCode != 0x5cadcb8d {
		return SymbolStatusMessageFlags{}, 0, fmt.Errorf("typecode mismatch, while reading SymbolStatusMessageFlags. Expected 0x5cadcb8d, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolStatusMessageFlags, n, err := Read_SymbolStatusMessageFlags(stream)
	if err != nil {
		return SymbolStatusMessageFlags{}, 0, err
	}
	byteCount += n
	return value_SymbolStatusMessageFlags, byteCount, nil
}

// SymbolStatusMessageFlags writer
func Write_SymbolStatusMessageFlags(stream Streams.IMitchWriter, value SymbolStatusMessageFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// SymbolStatusMessageFlags reader
func Read_SymbolStatusMessageFlags(stream Streams.IMitchReader) (value SymbolStatusMessageFlags, byteCount int, err error) {
	value = NewSymbolStatusMessageFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

// SymbolStatusMessage Declaration TypeCode: 0x5533b1cc
type SymbolStatusMessage struct {
	Length              uint16                          `json:"Length" xml:"Length,attr"`
	MessageType         byte                            `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond          uint32                          `json:"Nanosecond" xml:"Nanosecond,attr"`
	InstrumentID        uint32                          `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01          byte                            `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02          byte                            `json:"Reserved02" xml:"Reserved02,attr"`
	TradingStatus       SymbolStatusTradingStatus       `json:"TradingStatus" xml:"TradingStatus,attr"`
	Flags               SymbolStatusMessageFlags        `json:"Flags" xml:"Flags,attr"`
	Reason              string                          `json:"Reason" xml:"Reason,attr"`
	SessionChangeReason SymbolStatusSessionChangeReason `json:"SessionChangeReason" xml:"SessionChangeReason,attr"`
	NewEndTime          time.Time                       `json:"NewEndTime" xml:"NewEndTime,attr"`
	BookType            SymbolStatusBookType            `json:"BookType" xml:"BookType,attr"`
}

func NewSymbolStatusMessage() *SymbolStatusMessage {
	return &SymbolStatusMessage{
		Length:      29,
		MessageType: 72,
	}
}

// SymbolStatusMessage writer
func Write_SymbolStatusMessage(stream Streams.IMitchWriter, value *SymbolStatusMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(29)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(72)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: TradingStatus, Type: SymbolStatusTradingStatus
	//
	//
	n, err = stream.Write_byte(byte(value.TradingStatus))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Flags, Type: SymbolStatusMessageFlags
	//
	//
	n, err = Write_SymbolStatusMessageFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Reason, Type: string
	//
	//
	n, err = stream.Write_string(value.Reason, 4)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: SessionChangeReason, Type: SymbolStatusSessionChangeReason
	//
	//
	n, err = stream.Write_byte(byte(value.SessionChangeReason))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: NewEndTime, Type: Time
	//
	//
	n, err = stream.Write_mitch_time(value.NewEndTime)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: BookType, Type: SymbolStatusBookType
	//
	//
	n, err = stream.Write_byte(byte(value.BookType))
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolStatusMessage reader
func Read_SymbolStatusMessage(stream Streams.IMitchReader) (value *SymbolStatusMessage, byteCount int, err error) {
	value = NewSymbolStatusMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 29
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x48 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message SymbolStatusMessage was expected 0x48, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: TradingStatus, Type: SymbolStatusTradingStatus
	//
	//
	temp_TradingStatus, n, err := stream.Read_byte()
	value.TradingStatus = SymbolStatusTradingStatus(temp_TradingStatus)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Flags, Type: SymbolStatusMessageFlags
	//
	//
	value.Flags, n, err = Read_SymbolStatusMessageFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Reason, Type: string
	//
	//
	value.Reason, n, err = stream.Read_string(4)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: SessionChangeReason, Type: SymbolStatusSessionChangeReason
	//
	//
	temp_SessionChangeReason, n, err := stream.Read_byte()
	value.SessionChangeReason = SymbolStatusSessionChangeReason(temp_SessionChangeReason)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: NewEndTime, Type: Time
	//
	//
	value.NewEndTime, n, err = stream.Read_mitch_time()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: BookType, Type: SymbolStatusBookType
	//
	//
	temp_BookType, n, err := stream.Read_byte()
	value.BookType = SymbolStatusBookType(temp_BookType)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// SymbolStatusMessage reset
func Reset_SymbolStatusMessage(value *SymbolStatusMessage) {
	value.Length = 29
	value.MessageType = 72
	value.Nanosecond = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.TradingStatus = 0
	value.Flags = SymbolStatusMessageFlags{}
	value.Reason = ""
	value.SessionChangeReason = 0
	value.NewEndTime = time.Time{}
	value.BookType = 0
}

// SymbolStatusMessage clone
func Clone_SymbolStatusMessage(value *SymbolStatusMessage) *SymbolStatusMessage {
	result := NewSymbolStatusMessage()
	result.Nanosecond = value.Nanosecond
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.TradingStatus = value.TradingStatus
	result.Flags = value.Flags
	result.Reason = value.Reason
	result.SessionChangeReason = value.SessionChangeReason
	result.NewEndTime = value.NewEndTime
	result.BookType = value.BookType
	return result
}

// SymbolStatusMessage WriteMessage
func WriteMessage_SymbolStatusMessage(stream Streams.IMitchWriter, value *SymbolStatusMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x5533b1cc)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_SymbolStatusMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// SymbolStatusMessage reader
func ReadMessage_SymbolStatusMessage(stream Streams.IMitchReader) (*SymbolStatusMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x5533b1cc {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading SymbolStatusMessage. Expected 0x5533b1cc, got 0x%08x", typeCode)
	}
	byteCount += n
	value_SymbolStatusMessage, n, err := Read_SymbolStatusMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_SymbolStatusMessage, byteCount, nil
}

// SymbolStatusMessage TypeCode
func TypeCode_SymbolStatusMessage() uint32 {
	return 0x5533b1cc
}

// SymbolStatusMessage IsTypeCode
func IsTypeCode_SymbolStatusMessage(typeCode uint32) bool {
	return typeCode == 0x5533b1cc
}

type OrderSide byte

//noinspection ALL
const (
	OrderSide_BuyOrder  = 'B' // default value: byte('B')
	OrderSide_SellOrder = 'S' // default value: byte('S')
)

// OrderSide Declaration TypeCode: 0x22a70cb8
func NewOrderSide() OrderSide {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading OrderSide. Expected 0x22a70cb8, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderSide, n, err := Read_OrderSide(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_OrderSide, byteCount, nil
}

// OrderSide writer
func Write_OrderSide(stream Streams.IMitchWriter, value OrderSide) (int, error) {
	return stream.Write_byte(byte(value))
}

// OrderSide reader
func Read_OrderSide(stream Streams.IMitchReader) (value OrderSide, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return OrderSide(v), b, e
}

type AddOrderFlags struct {
	Flags byte
}

func NewAddOrderFlags() AddOrderFlags {
	return AddOrderFlags{}
}

// AddOrderFlags Declaration TypeCode: 0xf14d9a43
// AddOrderFlags WriteMessage
func WriteMessage_AddOrderFlags(stream Streams.IMitchWriter, value AddOrderFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xf14d9a43)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_AddOrderFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AddOrderFlags reader
func ReadMessage_AddOrderFlags(stream Streams.IMitchReader) (AddOrderFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return AddOrderFlags{}, 0, err
	}
	if typeCode != 0xf14d9a43 {
		return AddOrderFlags{}, 0, fmt.Errorf("typecode mismatch, while reading AddOrderFlags. Expected 0xf14d9a43, got 0x%08x", typeCode)
	}
	byteCount += n
	value_AddOrderFlags, n, err := Read_AddOrderFlags(stream)
	if err != nil {
		return AddOrderFlags{}, 0, err
	}
	byteCount += n
	return value_AddOrderFlags, byteCount, nil
}

// AddOrderFlags writer
func Write_AddOrderFlags(stream Streams.IMitchWriter, value AddOrderFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// AddOrderFlags reader
func Read_AddOrderFlags(stream Streams.IMitchReader) (value AddOrderFlags, byteCount int, err error) {
	value = NewAddOrderFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

// AddOrderMessage Declaration TypeCode: 0x4ace6131
type AddOrderMessage struct {
	Length       uint16        `json:"Length" xml:"Length,attr"`
	MessageType  byte          `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond   uint32        `json:"Nanosecond" xml:"Nanosecond,attr"`
	OrderId      uint64        `json:"OrderId" xml:"OrderId,attr"`
	Side         OrderSide     `json:"Side" xml:"Side,attr"`
	Quantity     uint32        `json:"Quantity" xml:"Quantity,attr"`
	InstrumentID uint32        `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01   byte          `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02   byte          `json:"Reserved02" xml:"Reserved02,attr"`
	Price        float64       `json:"Price" xml:"Price,attr"`
	Flags        AddOrderFlags `json:"Flags" xml:"Flags,attr"`
}

func NewAddOrderMessage() *AddOrderMessage {
	return &AddOrderMessage{
		Length:      35,
		MessageType: 65,
	}
}

// AddOrderMessage writer
func Write_AddOrderMessage(stream Streams.IMitchWriter, value *AddOrderMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(35)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(65)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderId, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.OrderId)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Side, Type: OrderSide
	//
	//
	n, err = stream.Write_byte(byte(value.Side))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Quantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Flags, Type: AddOrderFlags
	//
	//
	n, err = Write_AddOrderFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AddOrderMessage reader
func Read_AddOrderMessage(stream Streams.IMitchReader) (value *AddOrderMessage, byteCount int, err error) {
	value = NewAddOrderMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 35
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x41 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message AddOrderMessage was expected 0x41, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderId, Type: uint64
	//
	//
	value.OrderId, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Side, Type: OrderSide
	//
	//
	temp_Side, n, err := stream.Read_byte()
	value.Side = OrderSide(temp_Side)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Quantity, Type: uint32
	//
	//
	value.Quantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Flags, Type: AddOrderFlags
	//
	//
	value.Flags, n, err = Read_AddOrderFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// AddOrderMessage reset
func Reset_AddOrderMessage(value *AddOrderMessage) {
	value.Length = 35
	value.MessageType = 65
	value.Nanosecond = 0
	value.OrderId = 0
	value.Side = 0
	value.Quantity = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.Price = 0.0
	value.Flags = AddOrderFlags{}
}

// AddOrderMessage clone
func Clone_AddOrderMessage(value *AddOrderMessage) *AddOrderMessage {
	result := NewAddOrderMessage()
	result.Nanosecond = value.Nanosecond
	result.OrderId = value.OrderId
	result.Side = value.Side
	result.Quantity = value.Quantity
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.Price = value.Price
	result.Flags = value.Flags
	return result
}

// AddOrderMessage WriteMessage
func WriteMessage_AddOrderMessage(stream Streams.IMitchWriter, value *AddOrderMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x4ace6131)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_AddOrderMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AddOrderMessage reader
func ReadMessage_AddOrderMessage(stream Streams.IMitchReader) (*AddOrderMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x4ace6131 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading AddOrderMessage. Expected 0x4ace6131, got 0x%08x", typeCode)
	}
	byteCount += n
	value_AddOrderMessage, n, err := Read_AddOrderMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_AddOrderMessage, byteCount, nil
}

// AddOrderMessage TypeCode
func TypeCode_AddOrderMessage() uint32 {
	return 0x4ace6131
}

// AddOrderMessage IsTypeCode
func IsTypeCode_AddOrderMessage(typeCode uint32) bool {
	return typeCode == 0x4ace6131
}

type AddAttributedOrderFlags struct {
	Flags byte
}

func NewAddAttributedOrderFlags() AddAttributedOrderFlags {
	return AddAttributedOrderFlags{}
}

// AddAttributedOrderFlags Declaration TypeCode: 0xcf3c7356
// AddAttributedOrderFlags WriteMessage
func WriteMessage_AddAttributedOrderFlags(stream Streams.IMitchWriter, value AddAttributedOrderFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xcf3c7356)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_AddAttributedOrderFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AddAttributedOrderFlags reader
func ReadMessage_AddAttributedOrderFlags(stream Streams.IMitchReader) (AddAttributedOrderFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return AddAttributedOrderFlags{}, 0, err
	}
	if typeCode != 0xcf3c7356 {
		return AddAttributedOrderFlags{}, 0, fmt.Errorf("typecode mismatch, while reading AddAttributedOrderFlags. Expected 0xcf3c7356, got 0x%08x", typeCode)
	}
	byteCount += n
	value_AddAttributedOrderFlags, n, err := Read_AddAttributedOrderFlags(stream)
	if err != nil {
		return AddAttributedOrderFlags{}, 0, err
	}
	byteCount += n
	return value_AddAttributedOrderFlags, byteCount, nil
}

// AddAttributedOrderFlags writer
func Write_AddAttributedOrderFlags(stream Streams.IMitchWriter, value AddAttributedOrderFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// AddAttributedOrderFlags reader
func Read_AddAttributedOrderFlags(stream Streams.IMitchReader) (value AddAttributedOrderFlags, byteCount int, err error) {
	value = NewAddAttributedOrderFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

// AddAttributedOrderMessage Declaration TypeCode: 0x2ac24905
type AddAttributedOrderMessage struct {
	Length       uint16                  `json:"Length" xml:"Length,attr"`
	MessageType  byte                    `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond   uint32                  `json:"Nanosecond" xml:"Nanosecond,attr"`
	OrderID      uint64                  `json:"OrderID" xml:"OrderID,attr"`
	Side         OrderSide               `json:"Side" xml:"Side,attr"`
	Quantity     uint32                  `json:"Quantity" xml:"Quantity,attr"`
	InstrumentID uint32                  `json:"InstrumentID" xml:"InstrumentID,attr"`
	Price        float64                 `json:"Price" xml:"Price,attr"`
	Attribution  string                  `json:"Attribution" xml:"Attribution,attr"`
	Flags        AddAttributedOrderFlags `json:"Flags" xml:"Flags,attr"`
}

func NewAddAttributedOrderMessage() *AddAttributedOrderMessage {
	return &AddAttributedOrderMessage{
		Length:      44,
		MessageType: 70,
	}
}

// AddAttributedOrderMessage writer
func Write_AddAttributedOrderMessage(stream Streams.IMitchWriter, value *AddAttributedOrderMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(44)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(70)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Side, Type: OrderSide
	//
	//
	n, err = stream.Write_byte(byte(value.Side))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Quantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Attribution, Type: string
	//
	//
	n, err = stream.Write_string(value.Attribution, 11)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Flags, Type: AddAttributedOrderFlags
	//
	//
	n, err = Write_AddAttributedOrderFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AddAttributedOrderMessage reader
func Read_AddAttributedOrderMessage(stream Streams.IMitchReader) (value *AddAttributedOrderMessage, byteCount int, err error) {
	value = NewAddAttributedOrderMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 44
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x46 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message AddAttributedOrderMessage was expected 0x46, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	value.OrderID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Side, Type: OrderSide
	//
	//
	temp_Side, n, err := stream.Read_byte()
	value.Side = OrderSide(temp_Side)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Quantity, Type: uint32
	//
	//
	value.Quantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Attribution, Type: string
	//
	//
	value.Attribution, n, err = stream.Read_string(11)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Flags, Type: AddAttributedOrderFlags
	//
	//
	value.Flags, n, err = Read_AddAttributedOrderFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// AddAttributedOrderMessage reset
func Reset_AddAttributedOrderMessage(value *AddAttributedOrderMessage) {
	value.Length = 44
	value.MessageType = 70
	value.Nanosecond = 0
	value.OrderID = 0
	value.Side = 0
	value.Quantity = 0
	value.InstrumentID = 0
	value.Price = 0.0
	value.Attribution = ""
	value.Flags = AddAttributedOrderFlags{}
}

// AddAttributedOrderMessage clone
func Clone_AddAttributedOrderMessage(value *AddAttributedOrderMessage) *AddAttributedOrderMessage {
	result := NewAddAttributedOrderMessage()
	result.Nanosecond = value.Nanosecond
	result.OrderID = value.OrderID
	result.Side = value.Side
	result.Quantity = value.Quantity
	result.InstrumentID = value.InstrumentID
	result.Price = value.Price
	result.Attribution = value.Attribution
	result.Flags = value.Flags
	return result
}

// AddAttributedOrderMessage WriteMessage
func WriteMessage_AddAttributedOrderMessage(stream Streams.IMitchWriter, value *AddAttributedOrderMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x2ac24905)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_AddAttributedOrderMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AddAttributedOrderMessage reader
func ReadMessage_AddAttributedOrderMessage(stream Streams.IMitchReader) (*AddAttributedOrderMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x2ac24905 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading AddAttributedOrderMessage. Expected 0x2ac24905, got 0x%08x", typeCode)
	}
	byteCount += n
	value_AddAttributedOrderMessage, n, err := Read_AddAttributedOrderMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_AddAttributedOrderMessage, byteCount, nil
}

// AddAttributedOrderMessage TypeCode
func TypeCode_AddAttributedOrderMessage() uint32 {
	return 0x2ac24905
}

// AddAttributedOrderMessage IsTypeCode
func IsTypeCode_AddAttributedOrderMessage(typeCode uint32) bool {
	return typeCode == 0x2ac24905
}

// OrderDeletedMessage Declaration TypeCode: 0x71ebfeee
type OrderDeletedMessage struct {
	Length      uint16 `json:"Length" xml:"Length,attr"`
	MessageType byte   `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond  uint32 `json:"Nanosecond" xml:"Nanosecond,attr"`
	OrderID     uint64 `json:"OrderID" xml:"OrderID,attr"`
}

func NewOrderDeletedMessage() *OrderDeletedMessage {
	return &OrderDeletedMessage{
		Length:      15,
		MessageType: 68,
	}
}

// OrderDeletedMessage writer
func Write_OrderDeletedMessage(stream Streams.IMitchWriter, value *OrderDeletedMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(15)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(68)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderDeletedMessage reader
func Read_OrderDeletedMessage(stream Streams.IMitchReader) (value *OrderDeletedMessage, byteCount int, err error) {
	value = NewOrderDeletedMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 15
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x44 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message OrderDeletedMessage was expected 0x44, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	value.OrderID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// OrderDeletedMessage reset
func Reset_OrderDeletedMessage(value *OrderDeletedMessage) {
	value.Length = 15
	value.MessageType = 68
	value.Nanosecond = 0
	value.OrderID = 0
}

// OrderDeletedMessage clone
func Clone_OrderDeletedMessage(value *OrderDeletedMessage) *OrderDeletedMessage {
	result := NewOrderDeletedMessage()
	result.Nanosecond = value.Nanosecond
	result.OrderID = value.OrderID
	return result
}

// OrderDeletedMessage WriteMessage
func WriteMessage_OrderDeletedMessage(stream Streams.IMitchWriter, value *OrderDeletedMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x71ebfeee)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderDeletedMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderDeletedMessage reader
func ReadMessage_OrderDeletedMessage(stream Streams.IMitchReader) (*OrderDeletedMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x71ebfeee {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading OrderDeletedMessage. Expected 0x71ebfeee, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderDeletedMessage, n, err := Read_OrderDeletedMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_OrderDeletedMessage, byteCount, nil
}

// OrderDeletedMessage TypeCode
func TypeCode_OrderDeletedMessage() uint32 {
	return 0x71ebfeee
}

// OrderDeletedMessage IsTypeCode
func IsTypeCode_OrderDeletedMessage(typeCode uint32) bool {
	return typeCode == 0x71ebfeee
}

type OrderModifiedFlags struct {
	Flags byte
}

func NewOrderModifiedFlags() OrderModifiedFlags {
	return OrderModifiedFlags{}
}

// OrderModifiedFlags Declaration TypeCode: 0x4bfbc52c
// OrderModifiedFlags WriteMessage
func WriteMessage_OrderModifiedFlags(stream Streams.IMitchWriter, value OrderModifiedFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x4bfbc52c)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderModifiedFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderModifiedFlags reader
func ReadMessage_OrderModifiedFlags(stream Streams.IMitchReader) (OrderModifiedFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return OrderModifiedFlags{}, 0, err
	}
	if typeCode != 0x4bfbc52c {
		return OrderModifiedFlags{}, 0, fmt.Errorf("typecode mismatch, while reading OrderModifiedFlags. Expected 0x4bfbc52c, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderModifiedFlags, n, err := Read_OrderModifiedFlags(stream)
	if err != nil {
		return OrderModifiedFlags{}, 0, err
	}
	byteCount += n
	return value_OrderModifiedFlags, byteCount, nil
}

// OrderModifiedFlags writer
func Write_OrderModifiedFlags(stream Streams.IMitchWriter, value OrderModifiedFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// OrderModifiedFlags reader
func Read_OrderModifiedFlags(stream Streams.IMitchReader) (value OrderModifiedFlags, byteCount int, err error) {
	value = NewOrderModifiedFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

// OrderModifiedMessage Declaration TypeCode: 0x34838f04
type OrderModifiedMessage struct {
	Length      uint16             `json:"Length" xml:"Length,attr"`
	MessageType byte               `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond  uint32             `json:"Nanosecond" xml:"Nanosecond,attr"`
	OrderID     uint64             `json:"OrderID" xml:"OrderID,attr"`
	NewQuantity uint32             `json:"NewQuantity" xml:"NewQuantity,attr"`
	NewPrice    float64            `json:"NewPrice" xml:"NewPrice,attr"`
	Flags       OrderModifiedFlags `json:"Flags" xml:"Flags,attr"`
}

func NewOrderModifiedMessage() *OrderModifiedMessage {
	return &OrderModifiedMessage{
		Length:      28,
		MessageType: 85,
	}
}

// OrderModifiedMessage writer
func Write_OrderModifiedMessage(stream Streams.IMitchWriter, value *OrderModifiedMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(28)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(85)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: NewQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.NewQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: NewPrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.NewPrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Flags, Type: OrderModifiedFlags
	//
	//
	n, err = Write_OrderModifiedFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderModifiedMessage reader
func Read_OrderModifiedMessage(stream Streams.IMitchReader) (value *OrderModifiedMessage, byteCount int, err error) {
	value = NewOrderModifiedMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 28
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x55 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message OrderModifiedMessage was expected 0x55, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	value.OrderID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: NewQuantity, Type: uint32
	//
	//
	value.NewQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: NewPrice, Type: float64
	//
	//
	value.NewPrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Flags, Type: OrderModifiedFlags
	//
	//
	value.Flags, n, err = Read_OrderModifiedFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// OrderModifiedMessage reset
func Reset_OrderModifiedMessage(value *OrderModifiedMessage) {
	value.Length = 28
	value.MessageType = 85
	value.Nanosecond = 0
	value.OrderID = 0
	value.NewQuantity = 0
	value.NewPrice = 0.0
	value.Flags = OrderModifiedFlags{}
}

// OrderModifiedMessage clone
func Clone_OrderModifiedMessage(value *OrderModifiedMessage) *OrderModifiedMessage {
	result := NewOrderModifiedMessage()
	result.Nanosecond = value.Nanosecond
	result.OrderID = value.OrderID
	result.NewQuantity = value.NewQuantity
	result.NewPrice = value.NewPrice
	result.Flags = value.Flags
	return result
}

// OrderModifiedMessage WriteMessage
func WriteMessage_OrderModifiedMessage(stream Streams.IMitchWriter, value *OrderModifiedMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x34838f04)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderModifiedMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderModifiedMessage reader
func ReadMessage_OrderModifiedMessage(stream Streams.IMitchReader) (*OrderModifiedMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x34838f04 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading OrderModifiedMessage. Expected 0x34838f04, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderModifiedMessage, n, err := Read_OrderModifiedMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_OrderModifiedMessage, byteCount, nil
}

// OrderModifiedMessage TypeCode
func TypeCode_OrderModifiedMessage() uint32 {
	return 0x34838f04
}

// OrderModifiedMessage IsTypeCode
func IsTypeCode_OrderModifiedMessage(typeCode uint32) bool {
	return typeCode == 0x34838f04
}

type OrderBookClearSubBook byte

//noinspection ALL
const (
	OrderBookClearSubBook_OnBook           = 1  // default value: byte(1)
	OrderBookClearSubBook_OffBook          = 2  // default value: byte(2)
	OrderBookClearSubBook_BulletinBoard    = 9  // default value: byte(9)
	OrderBookClearSubBook_NegotiatedTrades = 11 // default value: byte(11)
)

// OrderBookClearSubBook Declaration TypeCode: 0x5ead6996
func NewOrderBookClearSubBook() OrderBookClearSubBook {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading OrderBookClearSubBook. Expected 0x5ead6996, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderBookClearSubBook, n, err := Read_OrderBookClearSubBook(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_OrderBookClearSubBook, byteCount, nil
}

// OrderBookClearSubBook writer
func Write_OrderBookClearSubBook(stream Streams.IMitchWriter, value OrderBookClearSubBook) (int, error) {
	return stream.Write_byte(byte(value))
}

// OrderBookClearSubBook reader
func Read_OrderBookClearSubBook(stream Streams.IMitchReader) (value OrderBookClearSubBook, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return OrderBookClearSubBook(v), b, e
}

type OrderBookClearBookType byte

//noinspection ALL
const (
	OrderBookClearBookType_MBO       = 0 // default value: byte(0)
	OrderBookClearBookType_TopOfBook = 1 // default value: byte(1)
)

// OrderBookClearBookType Declaration TypeCode: 0x7d3b6454
func NewOrderBookClearBookType() OrderBookClearBookType {
	return 0
}

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
		return 0, 0, fmt.Errorf("typecode mismatch, while reading OrderBookClearBookType. Expected 0x7d3b6454, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderBookClearBookType, n, err := Read_OrderBookClearBookType(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_OrderBookClearBookType, byteCount, nil
}

// OrderBookClearBookType writer
func Write_OrderBookClearBookType(stream Streams.IMitchWriter, value OrderBookClearBookType) (int, error) {
	return stream.Write_byte(byte(value))
}

// OrderBookClearBookType reader
func Read_OrderBookClearBookType(stream Streams.IMitchReader) (value OrderBookClearBookType, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return OrderBookClearBookType(v), b, e
}

// OrderBookClearMessage Declaration TypeCode: 0x90074430
type OrderBookClearMessage struct {
	Length       uint16                 `json:"Length" xml:"Length,attr"`
	MessageType  byte                   `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond   uint32                 `json:"Nanosecond" xml:"Nanosecond,attr"`
	InstrumentID uint32                 `json:"InstrumentID" xml:"InstrumentID,attr"`
	SubBook      OrderBookClearSubBook  `json:"SubBook" xml:"SubBook,attr"`
	BookType     OrderBookClearBookType `json:"BookType" xml:"BookType,attr"`
}

func NewOrderBookClearMessage() *OrderBookClearMessage {
	return &OrderBookClearMessage{
		Length:      13,
		MessageType: 121,
	}
}

// OrderBookClearMessage writer
func Write_OrderBookClearMessage(stream Streams.IMitchWriter, value *OrderBookClearMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(13)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(121)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: SubBook, Type: OrderBookClearSubBook
	//
	//
	n, err = stream.Write_byte(byte(value.SubBook))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: BookType, Type: OrderBookClearBookType
	//
	//
	n, err = stream.Write_byte(byte(value.BookType))
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderBookClearMessage reader
func Read_OrderBookClearMessage(stream Streams.IMitchReader) (value *OrderBookClearMessage, byteCount int, err error) {
	value = NewOrderBookClearMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 13
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x79 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message OrderBookClearMessage was expected 0x79, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: SubBook, Type: OrderBookClearSubBook
	//
	//
	temp_SubBook, n, err := stream.Read_byte()
	value.SubBook = OrderBookClearSubBook(temp_SubBook)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: BookType, Type: OrderBookClearBookType
	//
	//
	temp_BookType, n, err := stream.Read_byte()
	value.BookType = OrderBookClearBookType(temp_BookType)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// OrderBookClearMessage reset
func Reset_OrderBookClearMessage(value *OrderBookClearMessage) {
	value.Length = 13
	value.MessageType = 121
	value.Nanosecond = 0
	value.InstrumentID = 0
	value.SubBook = 0
	value.BookType = 0
}

// OrderBookClearMessage clone
func Clone_OrderBookClearMessage(value *OrderBookClearMessage) *OrderBookClearMessage {
	result := NewOrderBookClearMessage()
	result.Nanosecond = value.Nanosecond
	result.InstrumentID = value.InstrumentID
	result.SubBook = value.SubBook
	result.BookType = value.BookType
	return result
}

// OrderBookClearMessage WriteMessage
func WriteMessage_OrderBookClearMessage(stream Streams.IMitchWriter, value *OrderBookClearMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x90074430)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderBookClearMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderBookClearMessage reader
func ReadMessage_OrderBookClearMessage(stream Streams.IMitchReader) (*OrderBookClearMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x90074430 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading OrderBookClearMessage. Expected 0x90074430, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderBookClearMessage, n, err := Read_OrderBookClearMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_OrderBookClearMessage, byteCount, nil
}

// OrderBookClearMessage TypeCode
func TypeCode_OrderBookClearMessage() uint32 {
	return 0x90074430
}

// OrderBookClearMessage IsTypeCode
func IsTypeCode_OrderBookClearMessage(typeCode uint32) bool {
	return typeCode == 0x90074430
}

// OrderExecutedMessage Declaration TypeCode: 0xf694c0e7
type OrderExecutedMessage struct {
	Length                   uint16  `json:"Length" xml:"Length,attr"`
	MessageType              byte    `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond               uint32  `json:"Nanosecond" xml:"Nanosecond,attr"`
	OrderID                  uint64  `json:"OrderID" xml:"OrderID,attr"`
	ExecutedQuantity         uint32  `json:"ExecutedQuantity" xml:"ExecutedQuantity,attr"`
	TradeID                  uint64  `json:"TradeID" xml:"TradeID,attr"`
	LastOptPx                float64 `json:"LastOptPx" xml:"LastOptPx,attr"`
	Volatility               float64 `json:"Volatility" xml:"Volatility,attr"`
	UnderlyingReferencePrice float64 `json:"UnderlyingReferencePrice" xml:"UnderlyingReferencePrice,attr"`
}

func NewOrderExecutedMessage() *OrderExecutedMessage {
	return &OrderExecutedMessage{
		Length:      51,
		MessageType: 69,
	}
}

// OrderExecutedMessage writer
func Write_OrderExecutedMessage(stream Streams.IMitchWriter, value *OrderExecutedMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(51)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(69)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: LastOptPx, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Volatility, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderExecutedMessage reader
func Read_OrderExecutedMessage(stream Streams.IMitchReader) (value *OrderExecutedMessage, byteCount int, err error) {
	value = NewOrderExecutedMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 51
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x45 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message OrderExecutedMessage was expected 0x45, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	value.OrderID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	value.ExecutedQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: LastOptPx, Type: float64
	//
	//
	value.LastOptPx, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Volatility, Type: float64
	//
	//
	value.Volatility, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	value.UnderlyingReferencePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// OrderExecutedMessage reset
func Reset_OrderExecutedMessage(value *OrderExecutedMessage) {
	value.Length = 51
	value.MessageType = 69
	value.Nanosecond = 0
	value.OrderID = 0
	value.ExecutedQuantity = 0
	value.TradeID = 0
	value.LastOptPx = 0.0
	value.Volatility = 0.0
	value.UnderlyingReferencePrice = 0.0
}

// OrderExecutedMessage clone
func Clone_OrderExecutedMessage(value *OrderExecutedMessage) *OrderExecutedMessage {
	result := NewOrderExecutedMessage()
	result.Nanosecond = value.Nanosecond
	result.OrderID = value.OrderID
	result.ExecutedQuantity = value.ExecutedQuantity
	result.TradeID = value.TradeID
	result.LastOptPx = value.LastOptPx
	result.Volatility = value.Volatility
	result.UnderlyingReferencePrice = value.UnderlyingReferencePrice
	return result
}

// OrderExecutedMessage WriteMessage
func WriteMessage_OrderExecutedMessage(stream Streams.IMitchWriter, value *OrderExecutedMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xf694c0e7)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderExecutedMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderExecutedMessage reader
func ReadMessage_OrderExecutedMessage(stream Streams.IMitchReader) (*OrderExecutedMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xf694c0e7 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading OrderExecutedMessage. Expected 0xf694c0e7, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderExecutedMessage, n, err := Read_OrderExecutedMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_OrderExecutedMessage, byteCount, nil
}

// OrderExecutedMessage TypeCode
func TypeCode_OrderExecutedMessage() uint32 {
	return 0xf694c0e7
}

// OrderExecutedMessage IsTypeCode
func IsTypeCode_OrderExecutedMessage(typeCode uint32) bool {
	return typeCode == 0xf694c0e7
}

// OrderExecutedWithPriceSizeMessage Declaration TypeCode: 0xca49e936
type OrderExecutedWithPriceSizeMessage struct {
	Length                   uint16  `json:"Length" xml:"Length,attr"`
	MessageType              byte    `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond               uint32  `json:"Nanosecond" xml:"Nanosecond,attr"`
	OrderID                  uint64  `json:"OrderID" xml:"OrderID,attr"`
	ExecutedQuantity         uint32  `json:"ExecutedQuantity" xml:"ExecutedQuantity,attr"`
	DisplayQuantity          uint32  `json:"DisplayQuantity" xml:"DisplayQuantity,attr"`
	TradeID                  uint64  `json:"TradeID" xml:"TradeID,attr"`
	Printable                byte    `json:"Printable" xml:"Printable,attr"`
	Price                    float64 `json:"Price" xml:"Price,attr"`
	LastOptPx                float64 `json:"LastOptPx" xml:"LastOptPx,attr"`
	Volatility               float64 `json:"Volatility" xml:"Volatility,attr"`
	UnderlyingReferencePrice float64 `json:"UnderlyingReferencePrice" xml:"UnderlyingReferencePrice,attr"`
}

func NewOrderExecutedWithPriceSizeMessage() *OrderExecutedWithPriceSizeMessage {
	return &OrderExecutedWithPriceSizeMessage{
		Length:      64,
		MessageType: 67,
	}
}

// OrderExecutedWithPriceSizeMessage writer
func Write_OrderExecutedWithPriceSizeMessage(stream Streams.IMitchWriter, value *OrderExecutedWithPriceSizeMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(64)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(67)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.OrderID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: DisplayQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.DisplayQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Printable, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Printable)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: LastOptPx, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Volatility, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderExecutedWithPriceSizeMessage reader
func Read_OrderExecutedWithPriceSizeMessage(stream Streams.IMitchReader) (value *OrderExecutedWithPriceSizeMessage, byteCount int, err error) {
	value = NewOrderExecutedWithPriceSizeMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 64
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x43 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message OrderExecutedWithPriceSizeMessage was expected 0x43, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: OrderID, Type: uint64
	//
	//
	value.OrderID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	value.ExecutedQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: DisplayQuantity, Type: uint32
	//
	//
	value.DisplayQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Printable, Type: byte
	//
	//
	value.Printable, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: LastOptPx, Type: float64
	//
	//
	value.LastOptPx, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Volatility, Type: float64
	//
	//
	value.Volatility, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	value.UnderlyingReferencePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// OrderExecutedWithPriceSizeMessage reset
func Reset_OrderExecutedWithPriceSizeMessage(value *OrderExecutedWithPriceSizeMessage) {
	value.Length = 64
	value.MessageType = 67
	value.Nanosecond = 0
	value.OrderID = 0
	value.ExecutedQuantity = 0
	value.DisplayQuantity = 0
	value.TradeID = 0
	value.Printable = 0
	value.Price = 0.0
	value.LastOptPx = 0.0
	value.Volatility = 0.0
	value.UnderlyingReferencePrice = 0.0
}

// OrderExecutedWithPriceSizeMessage clone
func Clone_OrderExecutedWithPriceSizeMessage(value *OrderExecutedWithPriceSizeMessage) *OrderExecutedWithPriceSizeMessage {
	result := NewOrderExecutedWithPriceSizeMessage()
	result.Nanosecond = value.Nanosecond
	result.OrderID = value.OrderID
	result.ExecutedQuantity = value.ExecutedQuantity
	result.DisplayQuantity = value.DisplayQuantity
	result.TradeID = value.TradeID
	result.Printable = value.Printable
	result.Price = value.Price
	result.LastOptPx = value.LastOptPx
	result.Volatility = value.Volatility
	result.UnderlyingReferencePrice = value.UnderlyingReferencePrice
	return result
}

// OrderExecutedWithPriceSizeMessage WriteMessage
func WriteMessage_OrderExecutedWithPriceSizeMessage(stream Streams.IMitchWriter, value *OrderExecutedWithPriceSizeMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xca49e936)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OrderExecutedWithPriceSizeMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OrderExecutedWithPriceSizeMessage reader
func ReadMessage_OrderExecutedWithPriceSizeMessage(stream Streams.IMitchReader) (*OrderExecutedWithPriceSizeMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xca49e936 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading OrderExecutedWithPriceSizeMessage. Expected 0xca49e936, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OrderExecutedWithPriceSizeMessage, n, err := Read_OrderExecutedWithPriceSizeMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_OrderExecutedWithPriceSizeMessage, byteCount, nil
}

// OrderExecutedWithPriceSizeMessage TypeCode
func TypeCode_OrderExecutedWithPriceSizeMessage() uint32 {
	return 0xca49e936
}

// OrderExecutedWithPriceSizeMessage IsTypeCode
func IsTypeCode_OrderExecutedWithPriceSizeMessage(typeCode uint32) bool {
	return typeCode == 0xca49e936
}

type TradeMessageFlags struct {
	Flags byte
}

func NewTradeMessageFlags() TradeMessageFlags {
	return TradeMessageFlags{}
}

// TradeMessageFlags Declaration TypeCode: 0x8fa6f67d
// TradeMessageFlags WriteMessage
func WriteMessage_TradeMessageFlags(stream Streams.IMitchWriter, value TradeMessageFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x8fa6f67d)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TradeMessageFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TradeMessageFlags reader
func ReadMessage_TradeMessageFlags(stream Streams.IMitchReader) (TradeMessageFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return TradeMessageFlags{}, 0, err
	}
	if typeCode != 0x8fa6f67d {
		return TradeMessageFlags{}, 0, fmt.Errorf("typecode mismatch, while reading TradeMessageFlags. Expected 0x8fa6f67d, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TradeMessageFlags, n, err := Read_TradeMessageFlags(stream)
	if err != nil {
		return TradeMessageFlags{}, 0, err
	}
	byteCount += n
	return value_TradeMessageFlags, byteCount, nil
}

// TradeMessageFlags writer
func Write_TradeMessageFlags(stream Streams.IMitchWriter, value TradeMessageFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// TradeMessageFlags reader
func Read_TradeMessageFlags(stream Streams.IMitchReader) (value TradeMessageFlags, byteCount int, err error) {
	value = NewTradeMessageFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

type TradeMessageSubBook byte

//noinspection ALL
const (
	TradeMessageSubBook_Regular          = 1  // default value: byte(1)
	TradeMessageSubBook_NegotiatedTrades = 11 // default value: byte(11)
)

// TradeMessageSubBook Declaration TypeCode: 0x09c46d4d
func NewTradeMessageSubBook() TradeMessageSubBook {
	return 0
}

// TradeMessageSubBook WriteMessage
func WriteMessage_TradeMessageSubBook(stream Streams.IMitchWriter, value TradeMessageSubBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x09c46d4d)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TradeMessageSubBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TradeMessageSubBook reader
func ReadMessage_TradeMessageSubBook(stream Streams.IMitchReader) (TradeMessageSubBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x09c46d4d {
		return 0, 0, fmt.Errorf("typecode mismatch, while reading TradeMessageSubBook. Expected 0x09c46d4d, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TradeMessageSubBook, n, err := Read_TradeMessageSubBook(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_TradeMessageSubBook, byteCount, nil
}

// TradeMessageSubBook writer
func Write_TradeMessageSubBook(stream Streams.IMitchWriter, value TradeMessageSubBook) (int, error) {
	return stream.Write_byte(byte(value))
}

// TradeMessageSubBook reader
func Read_TradeMessageSubBook(stream Streams.IMitchReader) (value TradeMessageSubBook, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return TradeMessageSubBook(v), b, e
}

// TradeMessage Declaration TypeCode: 0xb1ce427e
type TradeMessage struct {
	Length                   uint16              `json:"Length" xml:"Length,attr"`
	MessageType              byte                `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond               uint32              `json:"Nanosecond" xml:"Nanosecond,attr"`
	ExecutedQuantity         uint32              `json:"ExecutedQuantity" xml:"ExecutedQuantity,attr"`
	InstrumentID             uint32              `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01               byte                `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02               byte                `json:"Reserved02" xml:"Reserved02,attr"`
	Price                    float64             `json:"Price" xml:"Price,attr"`
	TradeID                  uint64              `json:"TradeID" xml:"TradeID,attr"`
	SubBook                  TradeMessageSubBook `json:"SubBook" xml:"SubBook,attr"`
	Flags                    TradeMessageFlags   `json:"Flags" xml:"Flags,attr"`
	TradeSubType             string              `json:"TradeSubType" xml:"TradeSubType,attr"`
	LastOptPx                float64             `json:"LastOptPx" xml:"LastOptPx,attr"`
	Volatility               float64             `json:"Volatility" xml:"Volatility,attr"`
	UnderlyingReferencePrice float64             `json:"UnderlyingReferencePrice" xml:"UnderlyingReferencePrice,attr"`
}

func NewTradeMessage() *TradeMessage {
	return &TradeMessage{
		Length:      63,
		MessageType: 80,
	}
}

// TradeMessage writer
func Write_TradeMessage(stream Streams.IMitchWriter, value *TradeMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(63)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(80)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: SubBook, Type: TradeMessageSubBook
	//
	//
	n, err = stream.Write_byte(byte(value.SubBook))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Flags, Type: TradeMessageFlags
	//
	//
	n, err = Write_TradeMessageFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: TradeSubType, Type: string
	//
	//
	n, err = stream.Write_string(value.TradeSubType, 4)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: LastOptPx, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: Volatility, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TradeMessage reader
func Read_TradeMessage(stream Streams.IMitchReader) (value *TradeMessage, byteCount int, err error) {
	value = NewTradeMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 63
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x50 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message TradeMessage was expected 0x50, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	value.ExecutedQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: SubBook, Type: TradeMessageSubBook
	//
	//
	temp_SubBook, n, err := stream.Read_byte()
	value.SubBook = TradeMessageSubBook(temp_SubBook)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Flags, Type: TradeMessageFlags
	//
	//
	value.Flags, n, err = Read_TradeMessageFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: TradeSubType, Type: string
	//
	//
	value.TradeSubType, n, err = stream.Read_string(4)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: LastOptPx, Type: float64
	//
	//
	value.LastOptPx, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: Volatility, Type: float64
	//
	//
	value.Volatility, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	value.UnderlyingReferencePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// TradeMessage reset
func Reset_TradeMessage(value *TradeMessage) {
	value.Length = 63
	value.MessageType = 80
	value.Nanosecond = 0
	value.ExecutedQuantity = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.Price = 0.0
	value.TradeID = 0
	value.SubBook = 0
	value.Flags = TradeMessageFlags{}
	value.TradeSubType = ""
	value.LastOptPx = 0.0
	value.Volatility = 0.0
	value.UnderlyingReferencePrice = 0.0
}

// TradeMessage clone
func Clone_TradeMessage(value *TradeMessage) *TradeMessage {
	result := NewTradeMessage()
	result.Nanosecond = value.Nanosecond
	result.ExecutedQuantity = value.ExecutedQuantity
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.Price = value.Price
	result.TradeID = value.TradeID
	result.SubBook = value.SubBook
	result.Flags = value.Flags
	result.TradeSubType = value.TradeSubType
	result.LastOptPx = value.LastOptPx
	result.Volatility = value.Volatility
	result.UnderlyingReferencePrice = value.UnderlyingReferencePrice
	return result
}

// TradeMessage WriteMessage
func WriteMessage_TradeMessage(stream Streams.IMitchWriter, value *TradeMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xb1ce427e)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TradeMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TradeMessage reader
func ReadMessage_TradeMessage(stream Streams.IMitchReader) (*TradeMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xb1ce427e {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading TradeMessage. Expected 0xb1ce427e, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TradeMessage, n, err := Read_TradeMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_TradeMessage, byteCount, nil
}

// TradeMessage TypeCode
func TypeCode_TradeMessage() uint32 {
	return 0xb1ce427e
}

// TradeMessage IsTypeCode
func IsTypeCode_TradeMessage(typeCode uint32) bool {
	return typeCode == 0xb1ce427e
}

// AuctionTradeMessage Declaration TypeCode: 0x8569a219
type AuctionTradeMessage struct {
	Length                   uint16  `json:"Length" xml:"Length,attr"`
	MessageType              byte    `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond               uint32  `json:"Nanosecond" xml:"Nanosecond,attr"`
	Quantity                 uint32  `json:"Quantity" xml:"Quantity,attr"`
	InstrumentID             uint32  `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01               byte    `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02               byte    `json:"Reserved02" xml:"Reserved02,attr"`
	Price                    float64 `json:"Price" xml:"Price,attr"`
	TradeID                  uint64  `json:"TradeID" xml:"TradeID,attr"`
	AuctionType              byte    `json:"AuctionType" xml:"AuctionType,attr"`
	LastOptPx                float64 `json:"LastOptPx" xml:"LastOptPx,attr"`
	Volatility               float64 `json:"Volatility" xml:"Volatility,attr"`
	UnderlyingReferencePrice float64 `json:"UnderlyingReferencePrice" xml:"UnderlyingReferencePrice,attr"`
}

func NewAuctionTradeMessage() *AuctionTradeMessage {
	return &AuctionTradeMessage{
		Length:      58,
		MessageType: 81,
	}
}

// AuctionTradeMessage writer
func Write_AuctionTradeMessage(stream Streams.IMitchWriter, value *AuctionTradeMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(58)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(81)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: Quantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: AuctionType, Type: byte
	//
	//
	n, err = stream.Write_byte(value.AuctionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: LastOptPx, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: Volatility, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AuctionTradeMessage reader
func Read_AuctionTradeMessage(stream Streams.IMitchReader) (value *AuctionTradeMessage, byteCount int, err error) {
	value = NewAuctionTradeMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 58
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x51 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message AuctionTradeMessage was expected 0x51, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: Quantity, Type: uint32
	//
	//
	value.Quantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: AuctionType, Type: byte
	//
	//
	value.AuctionType, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: LastOptPx, Type: float64
	//
	//
	value.LastOptPx, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: Volatility, Type: float64
	//
	//
	value.Volatility, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	value.UnderlyingReferencePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// AuctionTradeMessage reset
func Reset_AuctionTradeMessage(value *AuctionTradeMessage) {
	value.Length = 58
	value.MessageType = 81
	value.Nanosecond = 0
	value.Quantity = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.Price = 0.0
	value.TradeID = 0
	value.AuctionType = 0
	value.LastOptPx = 0.0
	value.Volatility = 0.0
	value.UnderlyingReferencePrice = 0.0
}

// AuctionTradeMessage clone
func Clone_AuctionTradeMessage(value *AuctionTradeMessage) *AuctionTradeMessage {
	result := NewAuctionTradeMessage()
	result.Nanosecond = value.Nanosecond
	result.Quantity = value.Quantity
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.Price = value.Price
	result.TradeID = value.TradeID
	result.AuctionType = value.AuctionType
	result.LastOptPx = value.LastOptPx
	result.Volatility = value.Volatility
	result.UnderlyingReferencePrice = value.UnderlyingReferencePrice
	return result
}

// AuctionTradeMessage WriteMessage
func WriteMessage_AuctionTradeMessage(stream Streams.IMitchWriter, value *AuctionTradeMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x8569a219)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_AuctionTradeMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AuctionTradeMessage reader
func ReadMessage_AuctionTradeMessage(stream Streams.IMitchReader) (*AuctionTradeMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x8569a219 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading AuctionTradeMessage. Expected 0x8569a219, got 0x%08x", typeCode)
	}
	byteCount += n
	value_AuctionTradeMessage, n, err := Read_AuctionTradeMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_AuctionTradeMessage, byteCount, nil
}

// AuctionTradeMessage TypeCode
func TypeCode_AuctionTradeMessage() uint32 {
	return 0x8569a219
}

// AuctionTradeMessage IsTypeCode
func IsTypeCode_AuctionTradeMessage(typeCode uint32) bool {
	return typeCode == 0x8569a219
}

// OffBookTradeMessage Declaration TypeCode: 0xa7b5c913
type OffBookTradeMessage struct {
	Length                   uint16    `json:"Length" xml:"Length,attr"`
	MessageType              byte      `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond               uint32    `json:"Nanosecond" xml:"Nanosecond,attr"`
	ExecutedQuantity         uint32    `json:"ExecutedQuantity" xml:"ExecutedQuantity,attr"`
	InstrumentID             uint32    `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01               byte      `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02               byte      `json:"Reserved02" xml:"Reserved02,attr"`
	Price                    float64   `json:"Price" xml:"Price,attr"`
	TradeID                  uint64    `json:"TradeID" xml:"TradeID,attr"`
	OffBookTradeType         string    `json:"OffBookTradeType" xml:"OffBookTradeType,attr"`
	TradeTime                time.Time `json:"TradeTime" xml:"TradeTime,attr"`
	TradeDate                time.Time `json:"TradeDate" xml:"TradeDate,attr"`
	LastOptPx                float64   `json:"LastOptPx" xml:"LastOptPx,attr"`
	Volatility               float64   `json:"Volatility" xml:"Volatility,attr"`
	UnderlyingReferencePrice float64   `json:"UnderlyingReferencePrice" xml:"UnderlyingReferencePrice,attr"`
}

func NewOffBookTradeMessage() *OffBookTradeMessage {
	return &OffBookTradeMessage{
		Length:      77,
		MessageType: 120,
	}
}

// OffBookTradeMessage writer
func Write_OffBookTradeMessage(stream Streams.IMitchWriter, value *OffBookTradeMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(77)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(120)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: OffBookTradeType, Type: string
	//
	//
	n, err = stream.Write_string(value.OffBookTradeType, 4)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: TradeTime, Type: Time
	//
	//
	n, err = stream.Write_mitch_time(value.TradeTime)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: TradeDate, Type: Time
	//
	//
	n, err = stream.Write_mitch_date(value.TradeDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: LastOptPx, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: Volatility, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OffBookTradeMessage reader
func Read_OffBookTradeMessage(stream Streams.IMitchReader) (value *OffBookTradeMessage, byteCount int, err error) {
	value = NewOffBookTradeMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 77
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x78 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message OffBookTradeMessage was expected 0x78, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	value.ExecutedQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: OffBookTradeType, Type: string
	//
	//
	value.OffBookTradeType, n, err = stream.Read_string(4)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: TradeTime, Type: Time
	//
	//
	value.TradeTime, n, err = stream.Read_mitch_time()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: TradeDate, Type: Time
	//
	//
	value.TradeDate, n, err = stream.Read_mitch_date()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: LastOptPx, Type: float64
	//
	//
	value.LastOptPx, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: Volatility, Type: float64
	//
	//
	value.Volatility, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	value.UnderlyingReferencePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// OffBookTradeMessage reset
func Reset_OffBookTradeMessage(value *OffBookTradeMessage) {
	value.Length = 77
	value.MessageType = 120
	value.Nanosecond = 0
	value.ExecutedQuantity = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.Price = 0.0
	value.TradeID = 0
	value.OffBookTradeType = ""
	value.TradeTime = time.Time{}
	value.TradeDate = time.Time{}
	value.LastOptPx = 0.0
	value.Volatility = 0.0
	value.UnderlyingReferencePrice = 0.0
}

// OffBookTradeMessage clone
func Clone_OffBookTradeMessage(value *OffBookTradeMessage) *OffBookTradeMessage {
	result := NewOffBookTradeMessage()
	result.Nanosecond = value.Nanosecond
	result.ExecutedQuantity = value.ExecutedQuantity
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.Price = value.Price
	result.TradeID = value.TradeID
	result.OffBookTradeType = value.OffBookTradeType
	result.TradeTime = value.TradeTime
	result.TradeDate = value.TradeDate
	result.LastOptPx = value.LastOptPx
	result.Volatility = value.Volatility
	result.UnderlyingReferencePrice = value.UnderlyingReferencePrice
	return result
}

// OffBookTradeMessage WriteMessage
func WriteMessage_OffBookTradeMessage(stream Streams.IMitchWriter, value *OffBookTradeMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xa7b5c913)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_OffBookTradeMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// OffBookTradeMessage reader
func ReadMessage_OffBookTradeMessage(stream Streams.IMitchReader) (*OffBookTradeMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xa7b5c913 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading OffBookTradeMessage. Expected 0xa7b5c913, got 0x%08x", typeCode)
	}
	byteCount += n
	value_OffBookTradeMessage, n, err := Read_OffBookTradeMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_OffBookTradeMessage, byteCount, nil
}

// OffBookTradeMessage TypeCode
func TypeCode_OffBookTradeMessage() uint32 {
	return 0xa7b5c913
}

// OffBookTradeMessage IsTypeCode
func IsTypeCode_OffBookTradeMessage(typeCode uint32) bool {
	return typeCode == 0xa7b5c913
}

// TradeBreakMessage Declaration TypeCode: 0x6c216472
type TradeBreakMessage struct {
	Length      uint16 `json:"Length" xml:"Length,attr"`
	MessageType byte   `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond  uint32 `json:"Nanosecond" xml:"Nanosecond,attr"`
	TradeID     uint64 `json:"TradeID" xml:"TradeID,attr"`
	TradeType   byte   `json:"TradeType" xml:"TradeType,attr"`
}

func NewTradeBreakMessage() *TradeBreakMessage {
	return &TradeBreakMessage{
		Length:      16,
		MessageType: 66,
	}
}

// TradeBreakMessage writer
func Write_TradeBreakMessage(stream Streams.IMitchWriter, value *TradeBreakMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(16)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(66)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: TradeType, Type: byte
	//
	//
	n, err = stream.Write_byte(value.TradeType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TradeBreakMessage reader
func Read_TradeBreakMessage(stream Streams.IMitchReader) (value *TradeBreakMessage, byteCount int, err error) {
	value = NewTradeBreakMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 16
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x42 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message TradeBreakMessage was expected 0x42, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: TradeType, Type: byte
	//
	//
	value.TradeType, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// TradeBreakMessage reset
func Reset_TradeBreakMessage(value *TradeBreakMessage) {
	value.Length = 16
	value.MessageType = 66
	value.Nanosecond = 0
	value.TradeID = 0
	value.TradeType = 0
}

// TradeBreakMessage clone
func Clone_TradeBreakMessage(value *TradeBreakMessage) *TradeBreakMessage {
	result := NewTradeBreakMessage()
	result.Nanosecond = value.Nanosecond
	result.TradeID = value.TradeID
	result.TradeType = value.TradeType
	return result
}

// TradeBreakMessage WriteMessage
func WriteMessage_TradeBreakMessage(stream Streams.IMitchWriter, value *TradeBreakMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x6c216472)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TradeBreakMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TradeBreakMessage reader
func ReadMessage_TradeBreakMessage(stream Streams.IMitchReader) (*TradeBreakMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x6c216472 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading TradeBreakMessage. Expected 0x6c216472, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TradeBreakMessage, n, err := Read_TradeBreakMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_TradeBreakMessage, byteCount, nil
}

// TradeBreakMessage TypeCode
func TypeCode_TradeBreakMessage() uint32 {
	return 0x6c216472
}

// TradeBreakMessage IsTypeCode
func IsTypeCode_TradeBreakMessage(typeCode uint32) bool {
	return typeCode == 0x6c216472
}

type RecoveryTradeMessageFlags struct {
	Flags byte
}

func NewRecoveryTradeMessageFlags() RecoveryTradeMessageFlags {
	return RecoveryTradeMessageFlags{}
}

// RecoveryTradeMessageFlags Declaration TypeCode: 0x041c16c8
// RecoveryTradeMessageFlags WriteMessage
func WriteMessage_RecoveryTradeMessageFlags(stream Streams.IMitchWriter, value RecoveryTradeMessageFlags) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x041c16c8)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_RecoveryTradeMessageFlags(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// RecoveryTradeMessageFlags reader
func ReadMessage_RecoveryTradeMessageFlags(stream Streams.IMitchReader) (RecoveryTradeMessageFlags, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return RecoveryTradeMessageFlags{}, 0, err
	}
	if typeCode != 0x041c16c8 {
		return RecoveryTradeMessageFlags{}, 0, fmt.Errorf("typecode mismatch, while reading RecoveryTradeMessageFlags. Expected 0x041c16c8, got 0x%08x", typeCode)
	}
	byteCount += n
	value_RecoveryTradeMessageFlags, n, err := Read_RecoveryTradeMessageFlags(stream)
	if err != nil {
		return RecoveryTradeMessageFlags{}, 0, err
	}
	byteCount += n
	return value_RecoveryTradeMessageFlags, byteCount, nil
}

// RecoveryTradeMessageFlags writer
func Write_RecoveryTradeMessageFlags(stream Streams.IMitchWriter, value RecoveryTradeMessageFlags) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// RecoveryTradeMessageFlags reader
func Read_RecoveryTradeMessageFlags(stream Streams.IMitchReader) (value RecoveryTradeMessageFlags, byteCount int, err error) {
	value = NewRecoveryTradeMessageFlags()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

type RecoveryTradeMessageSubBook byte

//noinspection ALL
const (
	RecoveryTradeMessageSubBook_Regular          = 1  // default value: byte(1)
	RecoveryTradeMessageSubBook_NegotiatedTrades = 11 // default value: byte(11)
)

// RecoveryTradeMessageSubBook Declaration TypeCode: 0x8d246869
func NewRecoveryTradeMessageSubBook() RecoveryTradeMessageSubBook {
	return 0
}

// RecoveryTradeMessageSubBook WriteMessage
func WriteMessage_RecoveryTradeMessageSubBook(stream Streams.IMitchWriter, value RecoveryTradeMessageSubBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x8d246869)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_RecoveryTradeMessageSubBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// RecoveryTradeMessageSubBook reader
func ReadMessage_RecoveryTradeMessageSubBook(stream Streams.IMitchReader) (RecoveryTradeMessageSubBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x8d246869 {
		return 0, 0, fmt.Errorf("typecode mismatch, while reading RecoveryTradeMessageSubBook. Expected 0x8d246869, got 0x%08x", typeCode)
	}
	byteCount += n
	value_RecoveryTradeMessageSubBook, n, err := Read_RecoveryTradeMessageSubBook(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_RecoveryTradeMessageSubBook, byteCount, nil
}

// RecoveryTradeMessageSubBook writer
func Write_RecoveryTradeMessageSubBook(stream Streams.IMitchWriter, value RecoveryTradeMessageSubBook) (int, error) {
	return stream.Write_byte(byte(value))
}

// RecoveryTradeMessageSubBook reader
func Read_RecoveryTradeMessageSubBook(stream Streams.IMitchReader) (value RecoveryTradeMessageSubBook, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return RecoveryTradeMessageSubBook(v), b, e
}

// RecoveryTradeMessage Declaration TypeCode: 0x800fdbc2
type RecoveryTradeMessage struct {
	Length                   uint16                      `json:"Length" xml:"Length,attr"`
	MessageType              byte                        `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond               uint32                      `json:"Nanosecond" xml:"Nanosecond,attr"`
	ExecutedQuantity         uint32                      `json:"ExecutedQuantity" xml:"ExecutedQuantity,attr"`
	InstrumentID             uint32                      `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01               byte                        `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02               byte                        `json:"Reserved02" xml:"Reserved02,attr"`
	Price                    float64                     `json:"Price" xml:"Price,attr"`
	TradeID                  uint64                      `json:"TradeID" xml:"TradeID,attr"`
	AuctionType              byte                        `json:"AuctionType" xml:"AuctionType,attr"`
	OffBookRFQTradeType      string                      `json:"OffBookRFQTradeType" xml:"OffBookRFQTradeType,attr"`
	TradeTime                time.Time                   `json:"TradeTime" xml:"TradeTime,attr"`
	TradeDate                time.Time                   `json:"TradeDate" xml:"TradeDate,attr"`
	ActionType               byte                        `json:"ActionType" xml:"ActionType,attr"`
	SubBook                  RecoveryTradeMessageSubBook `json:"SubBook" xml:"SubBook,attr"`
	Flags                    RecoveryTradeMessageFlags   `json:"Flags" xml:"Flags,attr"`
	LastOptPx                float64                     `json:"LastOptPx" xml:"LastOptPx,attr"`
	Volatility               float64                     `json:"Volatility" xml:"Volatility,attr"`
	UnderlyingReferencePrice float64                     `json:"UnderlyingReferencePrice" xml:"UnderlyingReferencePrice,attr"`
}

func NewRecoveryTradeMessage() *RecoveryTradeMessage {
	return &RecoveryTradeMessage{
		Length:      81,
		MessageType: 118,
	}
}

// RecoveryTradeMessage writer
func Write_RecoveryTradeMessage(stream Streams.IMitchWriter, value *RecoveryTradeMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(81)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(118)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.ExecutedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	n, err = stream.Write_uint64(value.TradeID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: AuctionType, Type: byte
	//
	//
	n, err = stream.Write_byte(value.AuctionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: OffBookRFQTradeType, Type: string
	//
	//
	n, err = stream.Write_string(value.OffBookRFQTradeType, 4)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: TradeTime, Type: Time
	//
	//
	n, err = stream.Write_mitch_time(value.TradeTime)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: TradeDate, Type: Time
	//
	//
	n, err = stream.Write_mitch_date(value.TradeDate)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: ActionType, Type: byte
	//
	//
	n, err = stream.Write_byte(value.ActionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: SubBook, Type: RecoveryTradeMessageSubBook
	//
	//
	n, err = stream.Write_byte(byte(value.SubBook))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 15, Member Name: Flags, Type: RecoveryTradeMessageFlags
	//
	//
	n, err = Write_RecoveryTradeMessageFlags(stream, value.Flags)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 16, Member Name: LastOptPx, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LastOptPx)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 17, Member Name: Volatility, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Volatility)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 18, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.UnderlyingReferencePrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// RecoveryTradeMessage reader
func Read_RecoveryTradeMessage(stream Streams.IMitchReader) (value *RecoveryTradeMessage, byteCount int, err error) {
	value = NewRecoveryTradeMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 81
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x76 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message RecoveryTradeMessage was expected 0x76, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: ExecutedQuantity, Type: uint32
	//
	//
	value.ExecutedQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: TradeID, Type: uint64
	//
	//
	value.TradeID, n, err = stream.Read_uint64()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: AuctionType, Type: byte
	//
	//
	value.AuctionType, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: OffBookRFQTradeType, Type: string
	//
	//
	value.OffBookRFQTradeType, n, err = stream.Read_string(4)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: TradeTime, Type: Time
	//
	//
	value.TradeTime, n, err = stream.Read_mitch_time()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: TradeDate, Type: Time
	//
	//
	value.TradeDate, n, err = stream.Read_mitch_date()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: ActionType, Type: byte
	//
	//
	value.ActionType, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: SubBook, Type: RecoveryTradeMessageSubBook
	//
	//
	temp_SubBook, n, err := stream.Read_byte()
	value.SubBook = RecoveryTradeMessageSubBook(temp_SubBook)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 15, Member Name: Flags, Type: RecoveryTradeMessageFlags
	//
	//
	value.Flags, n, err = Read_RecoveryTradeMessageFlags(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 16, Member Name: LastOptPx, Type: float64
	//
	//
	value.LastOptPx, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 17, Member Name: Volatility, Type: float64
	//
	//
	value.Volatility, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 18, Member Name: UnderlyingReferencePrice, Type: float64
	//
	//
	value.UnderlyingReferencePrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// RecoveryTradeMessage reset
func Reset_RecoveryTradeMessage(value *RecoveryTradeMessage) {
	value.Length = 81
	value.MessageType = 118
	value.Nanosecond = 0
	value.ExecutedQuantity = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.Price = 0.0
	value.TradeID = 0
	value.AuctionType = 0
	value.OffBookRFQTradeType = ""
	value.TradeTime = time.Time{}
	value.TradeDate = time.Time{}
	value.ActionType = 0
	value.SubBook = 0
	value.Flags = RecoveryTradeMessageFlags{}
	value.LastOptPx = 0.0
	value.Volatility = 0.0
	value.UnderlyingReferencePrice = 0.0
}

// RecoveryTradeMessage clone
func Clone_RecoveryTradeMessage(value *RecoveryTradeMessage) *RecoveryTradeMessage {
	result := NewRecoveryTradeMessage()
	result.Nanosecond = value.Nanosecond
	result.ExecutedQuantity = value.ExecutedQuantity
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.Price = value.Price
	result.TradeID = value.TradeID
	result.AuctionType = value.AuctionType
	result.OffBookRFQTradeType = value.OffBookRFQTradeType
	result.TradeTime = value.TradeTime
	result.TradeDate = value.TradeDate
	result.ActionType = value.ActionType
	result.SubBook = value.SubBook
	result.Flags = value.Flags
	result.LastOptPx = value.LastOptPx
	result.Volatility = value.Volatility
	result.UnderlyingReferencePrice = value.UnderlyingReferencePrice
	return result
}

// RecoveryTradeMessage WriteMessage
func WriteMessage_RecoveryTradeMessage(stream Streams.IMitchWriter, value *RecoveryTradeMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x800fdbc2)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_RecoveryTradeMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// RecoveryTradeMessage reader
func ReadMessage_RecoveryTradeMessage(stream Streams.IMitchReader) (*RecoveryTradeMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x800fdbc2 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading RecoveryTradeMessage. Expected 0x800fdbc2, got 0x%08x", typeCode)
	}
	byteCount += n
	value_RecoveryTradeMessage, n, err := Read_RecoveryTradeMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_RecoveryTradeMessage, byteCount, nil
}

// RecoveryTradeMessage TypeCode
func TypeCode_RecoveryTradeMessage() uint32 {
	return 0x800fdbc2
}

// RecoveryTradeMessage IsTypeCode
func IsTypeCode_RecoveryTradeMessage(typeCode uint32) bool {
	return typeCode == 0x800fdbc2
}

// AuctionInfoMessage Declaration TypeCode: 0x12617b15
type AuctionInfoMessage struct {
	Length             uint16  `json:"Length" xml:"Length,attr"`
	MessageType        byte    `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond         uint32  `json:"Nanosecond" xml:"Nanosecond,attr"`
	PairedQuantity     uint32  `json:"PairedQuantity" xml:"PairedQuantity,attr"`
	Reserved01         uint32  `json:"Reserved01" xml:"Reserved01,attr"`
	ImbalanceDirection byte    `json:"ImbalanceDirection" xml:"ImbalanceDirection,attr"`
	InstrumentID       uint32  `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved02         byte    `json:"Reserved02" xml:"Reserved02,attr"`
	Reserved03         byte    `json:"Reserved03" xml:"Reserved03,attr"`
	Price              float64 `json:"Price" xml:"Price,attr"`
	AuctionType        byte    `json:"AuctionType" xml:"AuctionType,attr"`
}

func NewAuctionInfoMessage() *AuctionInfoMessage {
	return &AuctionInfoMessage{
		Length:      31,
		MessageType: 73,
	}
}

// AuctionInfoMessage writer
func Write_AuctionInfoMessage(stream Streams.IMitchWriter, value *AuctionInfoMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(31)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(73)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: PairedQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.PairedQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: ImbalanceDirection, Type: byte
	//
	//
	n, err = stream.Write_byte(value.ImbalanceDirection)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Reserved03, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved03)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: AuctionType, Type: byte
	//
	//
	n, err = stream.Write_byte(value.AuctionType)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AuctionInfoMessage reader
func Read_AuctionInfoMessage(stream Streams.IMitchReader) (value *AuctionInfoMessage, byteCount int, err error) {
	value = NewAuctionInfoMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 31
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x49 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message AuctionInfoMessage was expected 0x49, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: PairedQuantity, Type: uint32
	//
	//
	value.PairedQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: uint32
	//
	//
	value.Reserved01, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: ImbalanceDirection, Type: byte
	//
	//
	value.ImbalanceDirection, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Reserved03, Type: byte
	//
	//
	value.Reserved03, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: AuctionType, Type: byte
	//
	//
	value.AuctionType, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// AuctionInfoMessage reset
func Reset_AuctionInfoMessage(value *AuctionInfoMessage) {
	value.Length = 31
	value.MessageType = 73
	value.Nanosecond = 0
	value.PairedQuantity = 0
	value.Reserved01 = 0
	value.ImbalanceDirection = 0
	value.InstrumentID = 0
	value.Reserved02 = 0
	value.Reserved03 = 0
	value.Price = 0.0
	value.AuctionType = 0
}

// AuctionInfoMessage clone
func Clone_AuctionInfoMessage(value *AuctionInfoMessage) *AuctionInfoMessage {
	result := NewAuctionInfoMessage()
	result.Nanosecond = value.Nanosecond
	result.PairedQuantity = value.PairedQuantity
	result.Reserved01 = value.Reserved01
	result.ImbalanceDirection = value.ImbalanceDirection
	result.InstrumentID = value.InstrumentID
	result.Reserved02 = value.Reserved02
	result.Reserved03 = value.Reserved03
	result.Price = value.Price
	result.AuctionType = value.AuctionType
	return result
}

// AuctionInfoMessage WriteMessage
func WriteMessage_AuctionInfoMessage(stream Streams.IMitchWriter, value *AuctionInfoMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x12617b15)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_AuctionInfoMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// AuctionInfoMessage reader
func ReadMessage_AuctionInfoMessage(stream Streams.IMitchReader) (*AuctionInfoMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x12617b15 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading AuctionInfoMessage. Expected 0x12617b15, got 0x%08x", typeCode)
	}
	byteCount += n
	value_AuctionInfoMessage, n, err := Read_AuctionInfoMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_AuctionInfoMessage, byteCount, nil
}

// AuctionInfoMessage TypeCode
func TypeCode_AuctionInfoMessage() uint32 {
	return 0x12617b15
}

// AuctionInfoMessage IsTypeCode
func IsTypeCode_AuctionInfoMessage(typeCode uint32) bool {
	return typeCode == 0x12617b15
}

type TStatisticsMessageOpenCloseIndicator byte

//noinspection ALL
const (
	TStatisticsMessageOpenCloseIndicator_UT             = 'A' // default value: byte('A')
	TStatisticsMessageOpenCloseIndicator_AT             = 'B' // default value: byte('B')
	TStatisticsMessageOpenCloseIndicator_MidOfBBO       = 'C' // default value: byte('C')
	TStatisticsMessageOpenCloseIndicator_LastAT         = 'D' // default value: byte('D')
	TStatisticsMessageOpenCloseIndicator_LastUT         = 'E' // default value: byte('E')
	TStatisticsMessageOpenCloseIndicator_Manual         = 'F' // default value: byte('F')
	TStatisticsMessageOpenCloseIndicator_Vwap           = 'H' // default value: byte('H')
	TStatisticsMessageOpenCloseIndicator_PreviousClose  = 'I' // default value: byte('I')
	TStatisticsMessageOpenCloseIndicator_Zero           = 'J' // default value: byte('J')
	TStatisticsMessageOpenCloseIndicator_BestBid        = 'U' // default value: byte('U')
	TStatisticsMessageOpenCloseIndicator_BestOffer      = 'V' // default value: byte('V')
	TStatisticsMessageOpenCloseIndicator_ReferencePrice = 'Y' // default value: byte('Y')
)

// TStatisticsMessageOpenCloseIndicator Declaration TypeCode: 0x9313c377
func NewTStatisticsMessageOpenCloseIndicator() TStatisticsMessageOpenCloseIndicator {
	return 0
}

// TStatisticsMessageOpenCloseIndicator WriteMessage
func WriteMessage_TStatisticsMessageOpenCloseIndicator(stream Streams.IMitchWriter, value TStatisticsMessageOpenCloseIndicator) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x9313c377)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TStatisticsMessageOpenCloseIndicator(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TStatisticsMessageOpenCloseIndicator reader
func ReadMessage_TStatisticsMessageOpenCloseIndicator(stream Streams.IMitchReader) (TStatisticsMessageOpenCloseIndicator, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x9313c377 {
		return 0, 0, fmt.Errorf("typecode mismatch, while reading TStatisticsMessageOpenCloseIndicator. Expected 0x9313c377, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TStatisticsMessageOpenCloseIndicator, n, err := Read_TStatisticsMessageOpenCloseIndicator(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_TStatisticsMessageOpenCloseIndicator, byteCount, nil
}

// TStatisticsMessageOpenCloseIndicator writer
func Write_TStatisticsMessageOpenCloseIndicator(stream Streams.IMitchWriter, value TStatisticsMessageOpenCloseIndicator) (int, error) {
	return stream.Write_byte(byte(value))
}

// TStatisticsMessageOpenCloseIndicator reader
func Read_TStatisticsMessageOpenCloseIndicator(stream Streams.IMitchReader) (value TStatisticsMessageOpenCloseIndicator, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return TStatisticsMessageOpenCloseIndicator(v), b, e
}

type TStatisticsMessageSubBook byte

//noinspection ALL
const (
	TStatisticsMessageSubBook_Regular       = 1 // default value: byte(1)
	TStatisticsMessageSubBook_OffBook       = 2 // default value: byte(2)
	TStatisticsMessageSubBook_BulletinBoard = 9 // default value: byte(9)
)

// TStatisticsMessageSubBook Declaration TypeCode: 0x7362a31a
func NewTStatisticsMessageSubBook() TStatisticsMessageSubBook {
	return 0
}

// TStatisticsMessageSubBook WriteMessage
func WriteMessage_TStatisticsMessageSubBook(stream Streams.IMitchWriter, value TStatisticsMessageSubBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x7362a31a)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TStatisticsMessageSubBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TStatisticsMessageSubBook reader
func ReadMessage_TStatisticsMessageSubBook(stream Streams.IMitchReader) (TStatisticsMessageSubBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return 0, 0, err
	}
	if typeCode != 0x7362a31a {
		return 0, 0, fmt.Errorf("typecode mismatch, while reading TStatisticsMessageSubBook. Expected 0x7362a31a, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TStatisticsMessageSubBook, n, err := Read_TStatisticsMessageSubBook(stream)
	if err != nil {
		return 0, 0, err
	}
	byteCount += n
	return value_TStatisticsMessageSubBook, byteCount, nil
}

// TStatisticsMessageSubBook writer
func Write_TStatisticsMessageSubBook(stream Streams.IMitchWriter, value TStatisticsMessageSubBook) (int, error) {
	return stream.Write_byte(byte(value))
}

// TStatisticsMessageSubBook reader
func Read_TStatisticsMessageSubBook(stream Streams.IMitchReader) (value TStatisticsMessageSubBook, byteCount int, err error) {
	v, b, e := stream.Read_byte()
	return TStatisticsMessageSubBook(v), b, e
}

// StatisticsMessage Declaration TypeCode: 0xc87d7e5a
type StatisticsMessage struct {
	Length             uint16                               `json:"Length" xml:"Length,attr"`
	MessageType        byte                                 `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond         uint32                               `json:"Nanosecond" xml:"Nanosecond,attr"`
	InstrumentID       uint32                               `json:"InstrumentID" xml:"InstrumentID,attr"`
	Reserved01         byte                                 `json:"Reserved01" xml:"Reserved01,attr"`
	Reserved02         byte                                 `json:"Reserved02" xml:"Reserved02,attr"`
	StatisticType      string                               `json:"StatisticType" xml:"StatisticType,attr"`
	Price              float64                              `json:"Price" xml:"Price,attr"`
	OpenCloseIndicator TStatisticsMessageOpenCloseIndicator `json:"OpenCloseIndicator" xml:"OpenCloseIndicator,attr"`
	SubBook            TStatisticsMessageSubBook            `json:"SubBook" xml:"SubBook,attr"`
}

func NewStatisticsMessage() *StatisticsMessage {
	return &StatisticsMessage{
		Length:      24,
		MessageType: 119,
	}
}

// StatisticsMessage writer
func Write_StatisticsMessage(stream Streams.IMitchWriter, value *StatisticsMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(24)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(119)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved02, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Reserved02)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: StatisticType, Type: string
	//
	//
	n, err = stream.Write_string(value.StatisticType, 1)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: OpenCloseIndicator, Type: TStatisticsMessageOpenCloseIndicator
	//
	//
	n, err = stream.Write_byte(byte(value.OpenCloseIndicator))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: SubBook, Type: TStatisticsMessageSubBook
	//
	//
	n, err = stream.Write_byte(byte(value.SubBook))
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// StatisticsMessage reader
func Read_StatisticsMessage(stream Streams.IMitchReader) (value *StatisticsMessage, byteCount int, err error) {
	value = NewStatisticsMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 24
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x77 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message StatisticsMessage was expected 0x77, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Reserved01, Type: byte
	//
	//
	value.Reserved01, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Reserved02, Type: byte
	//
	//
	value.Reserved02, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: StatisticType, Type: string
	//
	//
	value.StatisticType, n, err = stream.Read_string(1)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: OpenCloseIndicator, Type: TStatisticsMessageOpenCloseIndicator
	//
	//
	temp_OpenCloseIndicator, n, err := stream.Read_byte()
	value.OpenCloseIndicator = TStatisticsMessageOpenCloseIndicator(temp_OpenCloseIndicator)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: SubBook, Type: TStatisticsMessageSubBook
	//
	//
	temp_SubBook, n, err := stream.Read_byte()
	value.SubBook = TStatisticsMessageSubBook(temp_SubBook)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// StatisticsMessage reset
func Reset_StatisticsMessage(value *StatisticsMessage) {
	value.Length = 24
	value.MessageType = 119
	value.Nanosecond = 0
	value.InstrumentID = 0
	value.Reserved01 = 0
	value.Reserved02 = 0
	value.StatisticType = ""
	value.Price = 0.0
	value.OpenCloseIndicator = 0
	value.SubBook = 0
}

// StatisticsMessage clone
func Clone_StatisticsMessage(value *StatisticsMessage) *StatisticsMessage {
	result := NewStatisticsMessage()
	result.Nanosecond = value.Nanosecond
	result.InstrumentID = value.InstrumentID
	result.Reserved01 = value.Reserved01
	result.Reserved02 = value.Reserved02
	result.StatisticType = value.StatisticType
	result.Price = value.Price
	result.OpenCloseIndicator = value.OpenCloseIndicator
	result.SubBook = value.SubBook
	return result
}

// StatisticsMessage WriteMessage
func WriteMessage_StatisticsMessage(stream Streams.IMitchWriter, value *StatisticsMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xc87d7e5a)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_StatisticsMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// StatisticsMessage reader
func ReadMessage_StatisticsMessage(stream Streams.IMitchReader) (*StatisticsMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xc87d7e5a {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading StatisticsMessage. Expected 0xc87d7e5a, got 0x%08x", typeCode)
	}
	byteCount += n
	value_StatisticsMessage, n, err := Read_StatisticsMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_StatisticsMessage, byteCount, nil
}

// StatisticsMessage TypeCode
func TypeCode_StatisticsMessage() uint32 {
	return 0xc87d7e5a
}

// StatisticsMessage IsTypeCode
func IsTypeCode_StatisticsMessage(typeCode uint32) bool {
	return typeCode == 0xc87d7e5a
}

// ExtendedStatisticsMessage Declaration TypeCode: 0xff7aeb65
type ExtendedStatisticsMessage struct {
	Length                uint16                    `json:"Length" xml:"Length,attr"`
	MessageType           byte                      `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond            uint32                    `json:"Nanosecond" xml:"Nanosecond,attr"`
	InstrumentID          uint32                    `json:"InstrumentID" xml:"InstrumentID,attr"`
	HighPrice             float64                   `json:"HighPrice" xml:"HighPrice,attr"`
	LowPrice              float64                   `json:"LowPrice" xml:"LowPrice,attr"`
	Vwap                  float64                   `json:"Vwap" xml:"Vwap,attr"`
	Volume                uint32                    `json:"Volume" xml:"Volume,attr"`
	Turnover              float64                   `json:"Turnover" xml:"Turnover,attr"`
	NumberOfTrades        uint32                    `json:"NumberOfTrades" xml:"NumberOfTrades,attr"`
	Reserved01            string                    `json:"Reserved01" xml:"Reserved01,attr"`
	SubBook               TStatisticsMessageSubBook `json:"SubBook" xml:"SubBook,attr"`
	NotionalExposure      float64                   `json:"NotionalExposure" xml:"NotionalExposure,attr"`
	NotionalDeltaExposure float64                   `json:"NotionalDeltaExposure" xml:"NotionalDeltaExposure,attr"`
	OpenInterest          float64                   `json:"OpenInterest" xml:"OpenInterest,attr"`
}

func NewExtendedStatisticsMessage() *ExtendedStatisticsMessage {
	return &ExtendedStatisticsMessage{
		Length:      84,
		MessageType: 128,
	}
}

// ExtendedStatisticsMessage writer
func Write_ExtendedStatisticsMessage(stream Streams.IMitchWriter, value *ExtendedStatisticsMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(84)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(128)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: HighPrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.HighPrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: LowPrice, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.LowPrice)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Vwap, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Vwap)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Volume, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Volume)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Turnover, Type: float64
	//
	//
	n, err = stream.Write_mitch_price04(value.Turnover)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: NumberOfTrades, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.NumberOfTrades)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Reserved01, Type: string
	//
	//
	n, err = stream.Write_string(value.Reserved01, 8)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: SubBook, Type: TStatisticsMessageSubBook
	//
	//
	n, err = stream.Write_byte(byte(value.SubBook))
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: NotionalExposure, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.NotionalExposure)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: NotionalDeltaExposure, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.NotionalDeltaExposure)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: OpenInterest, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.OpenInterest)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// ExtendedStatisticsMessage reader
func Read_ExtendedStatisticsMessage(stream Streams.IMitchReader) (value *ExtendedStatisticsMessage, byteCount int, err error) {
	value = NewExtendedStatisticsMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 84
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x80 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message ExtendedStatisticsMessage was expected 0x80, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: HighPrice, Type: float64
	//
	//
	value.HighPrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: LowPrice, Type: float64
	//
	//
	value.LowPrice, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Vwap, Type: float64
	//
	//
	value.Vwap, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Volume, Type: uint32
	//
	//
	value.Volume, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Turnover, Type: float64
	//
	//
	value.Turnover, n, err = stream.Read_mitch_price04()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: NumberOfTrades, Type: uint32
	//
	//
	value.NumberOfTrades, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: Reserved01, Type: string
	//
	//
	value.Reserved01, n, err = stream.Read_string(8)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: SubBook, Type: TStatisticsMessageSubBook
	//
	//
	temp_SubBook, n, err := stream.Read_byte()
	value.SubBook = TStatisticsMessageSubBook(temp_SubBook)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 12, Member Name: NotionalExposure, Type: float64
	//
	//
	value.NotionalExposure, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 13, Member Name: NotionalDeltaExposure, Type: float64
	//
	//
	value.NotionalDeltaExposure, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 14, Member Name: OpenInterest, Type: float64
	//
	//
	value.OpenInterest, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// ExtendedStatisticsMessage reset
func Reset_ExtendedStatisticsMessage(value *ExtendedStatisticsMessage) {
	value.Length = 84
	value.MessageType = 128
	value.Nanosecond = 0
	value.InstrumentID = 0
	value.HighPrice = 0.0
	value.LowPrice = 0.0
	value.Vwap = 0.0
	value.Volume = 0
	value.Turnover = 0.0
	value.NumberOfTrades = 0
	value.Reserved01 = ""
	value.SubBook = 0
	value.NotionalExposure = 0.0
	value.NotionalDeltaExposure = 0.0
	value.OpenInterest = 0.0
}

// ExtendedStatisticsMessage clone
func Clone_ExtendedStatisticsMessage(value *ExtendedStatisticsMessage) *ExtendedStatisticsMessage {
	result := NewExtendedStatisticsMessage()
	result.Nanosecond = value.Nanosecond
	result.InstrumentID = value.InstrumentID
	result.HighPrice = value.HighPrice
	result.LowPrice = value.LowPrice
	result.Vwap = value.Vwap
	result.Volume = value.Volume
	result.Turnover = value.Turnover
	result.NumberOfTrades = value.NumberOfTrades
	result.Reserved01 = value.Reserved01
	result.SubBook = value.SubBook
	result.NotionalExposure = value.NotionalExposure
	result.NotionalDeltaExposure = value.NotionalDeltaExposure
	result.OpenInterest = value.OpenInterest
	return result
}

// ExtendedStatisticsMessage WriteMessage
func WriteMessage_ExtendedStatisticsMessage(stream Streams.IMitchWriter, value *ExtendedStatisticsMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xff7aeb65)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_ExtendedStatisticsMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// ExtendedStatisticsMessage reader
func ReadMessage_ExtendedStatisticsMessage(stream Streams.IMitchReader) (*ExtendedStatisticsMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xff7aeb65 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading ExtendedStatisticsMessage. Expected 0xff7aeb65, got 0x%08x", typeCode)
	}
	byteCount += n
	value_ExtendedStatisticsMessage, n, err := Read_ExtendedStatisticsMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_ExtendedStatisticsMessage, byteCount, nil
}

// ExtendedStatisticsMessage TypeCode
func TypeCode_ExtendedStatisticsMessage() uint32 {
	return 0xff7aeb65
}

// ExtendedStatisticsMessage IsTypeCode
func IsTypeCode_ExtendedStatisticsMessage(typeCode uint32) bool {
	return typeCode == 0xff7aeb65
}

// NewsMessage Declaration TypeCode: 0x8a0c79e9
type NewsMessage struct {
	Length      uint16    `json:"Length" xml:"Length,attr"`
	MessageType byte      `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond  uint32    `json:"Nanosecond" xml:"Nanosecond,attr"`
	Time        time.Time `json:"Time" xml:"Time,attr"`
	Urgency     byte      `json:"Urgency" xml:"Urgency,attr"`
	Headline    string    `json:"Headline" xml:"Headline,attr"`
	Text        string    `json:"Text" xml:"Text,attr"`
	Instruments string    `json:"Instruments" xml:"Instruments,attr"`
	Underlyings string    `json:"Underlyings" xml:"Underlyings,attr"`
}

func NewNewsMessage() *NewsMessage {
	return &NewsMessage{
		Length:      1066,
		MessageType: 117,
	}
}

// NewsMessage writer
func Write_NewsMessage(stream Streams.IMitchWriter, value *NewsMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(1066)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(117)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: Time, Type: Time
	//
	//
	n, err = stream.Write_mitch_time(value.Time)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Urgency, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Urgency)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Headline, Type: string
	//
	//
	n, err = stream.Write_string(value.Headline, 100)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Text, Type: string
	//
	//
	n, err = stream.Write_string(value.Text, 750)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Instruments, Type: string
	//
	//
	n, err = stream.Write_string(value.Instruments, 100)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Underlyings, Type: string
	//
	//
	n, err = stream.Write_string(value.Underlyings, 100)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// NewsMessage reader
func Read_NewsMessage(stream Streams.IMitchReader) (value *NewsMessage, byteCount int, err error) {
	value = NewNewsMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 1066
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x75 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message NewsMessage was expected 0x75, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: Time, Type: Time
	//
	//
	value.Time, n, err = stream.Read_mitch_time()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: Urgency, Type: byte
	//
	//
	value.Urgency, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: Headline, Type: string
	//
	//
	value.Headline, n, err = stream.Read_string(100)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Text, Type: string
	//
	//
	value.Text, n, err = stream.Read_string(750)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Instruments, Type: string
	//
	//
	value.Instruments, n, err = stream.Read_string(100)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Underlyings, Type: string
	//
	//
	value.Underlyings, n, err = stream.Read_string(100)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// NewsMessage reset
func Reset_NewsMessage(value *NewsMessage) {
	value.Length = 1066
	value.MessageType = 117
	value.Nanosecond = 0
	value.Time = time.Time{}
	value.Urgency = 0
	value.Headline = ""
	value.Text = ""
	value.Instruments = ""
	value.Underlyings = ""
}

// NewsMessage clone
func Clone_NewsMessage(value *NewsMessage) *NewsMessage {
	result := NewNewsMessage()
	result.Nanosecond = value.Nanosecond
	result.Time = value.Time
	result.Urgency = value.Urgency
	result.Headline = value.Headline
	result.Text = value.Text
	result.Instruments = value.Instruments
	result.Underlyings = value.Underlyings
	return result
}

// NewsMessage WriteMessage
func WriteMessage_NewsMessage(stream Streams.IMitchWriter, value *NewsMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0x8a0c79e9)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_NewsMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// NewsMessage reader
func ReadMessage_NewsMessage(stream Streams.IMitchReader) (*NewsMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0x8a0c79e9 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading NewsMessage. Expected 0x8a0c79e9, got 0x%08x", typeCode)
	}
	byteCount += n
	value_NewsMessage, n, err := Read_NewsMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_NewsMessage, byteCount, nil
}

// NewsMessage TypeCode
func TypeCode_NewsMessage() uint32 {
	return 0x8a0c79e9
}

// NewsMessage IsTypeCode
func IsTypeCode_NewsMessage(typeCode uint32) bool {
	return typeCode == 0x8a0c79e9
}

type TopOfBookMessageSubBook struct {
	Flags byte
}

func NewTopOfBookMessageSubBook() TopOfBookMessageSubBook {
	return TopOfBookMessageSubBook{}
}

// TopOfBookMessageSubBook Declaration TypeCode: 0xb416f305
// TopOfBookMessageSubBook WriteMessage
func WriteMessage_TopOfBookMessageSubBook(stream Streams.IMitchWriter, value TopOfBookMessageSubBook) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xb416f305)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TopOfBookMessageSubBook(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TopOfBookMessageSubBook reader
func ReadMessage_TopOfBookMessageSubBook(stream Streams.IMitchReader) (TopOfBookMessageSubBook, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return TopOfBookMessageSubBook{}, 0, err
	}
	if typeCode != 0xb416f305 {
		return TopOfBookMessageSubBook{}, 0, fmt.Errorf("typecode mismatch, while reading TopOfBookMessageSubBook. Expected 0xb416f305, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TopOfBookMessageSubBook, n, err := Read_TopOfBookMessageSubBook(stream)
	if err != nil {
		return TopOfBookMessageSubBook{}, 0, err
	}
	byteCount += n
	return value_TopOfBookMessageSubBook, byteCount, nil
}

// TopOfBookMessageSubBook writer
func Write_TopOfBookMessageSubBook(stream Streams.IMitchWriter, value TopOfBookMessageSubBook) (byteCount int, err error) {
	return stream.Write_byte(byte(value.Flags))
}

// TopOfBookMessageSubBook reader
func Read_TopOfBookMessageSubBook(stream Streams.IMitchReader) (value TopOfBookMessageSubBook, byteCount int, err error) {
	value = NewTopOfBookMessageSubBook()
	value.Flags, byteCount, err = stream.Read_byte()
	return value, byteCount, err
}

// TopOfBookMessage Declaration TypeCode: 0xac36e385
type TopOfBookMessage struct {
	Length              uint16                  `json:"Length" xml:"Length,attr"`
	MessageType         byte                    `json:"MessageType" xml:"MessageType,attr"`
	Nanosecond          uint32                  `json:"Nanosecond" xml:"Nanosecond,attr"`
	InstrumentID        uint32                  `json:"InstrumentID" xml:"InstrumentID,attr"`
	ReserveField        uint16                  `json:"ReserveField" xml:"ReserveField,attr"`
	SubBook             TopOfBookMessageSubBook `json:"SubBook" xml:"SubBook,attr"`
	Action              byte                    `json:"Action" xml:"Action,attr"`
	Side                byte                    `json:"Side" xml:"Side,attr"`
	Price               float64                 `json:"Price" xml:"Price,attr"`
	Quantity            uint32                  `json:"Quantity" xml:"Quantity,attr"`
	MarketOrderQuantity uint32                  `json:"MarketOrderQuantity" xml:"MarketOrderQuantity,attr"`
	Reserved01          uint16                  `json:"Reserved01" xml:"Reserved01,attr"`
}

func NewTopOfBookMessage() *TopOfBookMessage {
	return &TopOfBookMessage{
		Length:      34,
		MessageType: 113,
	}
}

// TopOfBookMessage writer
func Write_TopOfBookMessage(stream Streams.IMitchWriter, value *TopOfBookMessage) (byteCount int, err error) {
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	n, err = stream.Write_uint16(34)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	n, err = stream.Write_byte(113)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Nanosecond)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.InstrumentID)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: ReserveField, Type: uint16
	//
	//
	n, err = stream.Write_uint16(value.ReserveField)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: SubBook, Type: TopOfBookMessageSubBook
	//
	//
	n, err = Write_TopOfBookMessageSubBook(stream, value.SubBook)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Action, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Action)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Side, Type: byte
	//
	//
	n, err = stream.Write_byte(value.Side)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Price, Type: float64
	//
	//
	n, err = stream.Write_mitch_price08(value.Price)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Quantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.Quantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: MarketOrderQuantity, Type: uint32
	//
	//
	n, err = stream.Write_uint32(value.MarketOrderQuantity)
	if err != nil {
		return 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: Reserved01, Type: uint16
	//
	//
	n, err = stream.Write_uint16(value.Reserved01)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TopOfBookMessage reader
func Read_TopOfBookMessage(stream Streams.IMitchReader) (value *TopOfBookMessage, byteCount int, err error) {
	value = NewTopOfBookMessage()
	var n int
	//
	//
	// Index: 0, Member Name: Length, Type: uint16
	//
	//
	_, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.Length = 34
	//
	//
	// Index: 1, Member Name: MessageType, Type: byte
	//
	//
	b, n, err := stream.Read_byte()
	if b != 0x71 {
		return nil, 0, fmt.Errorf("message type numbers does not match up. For Message TopOfBookMessage was expected 0x71, but 0x%x was found.)", b)
	}
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	value.MessageType = b
	//
	//
	// Index: 2, Member Name: Nanosecond, Type: uint32
	//
	//
	value.Nanosecond, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 3, Member Name: InstrumentID, Type: uint32
	//
	//
	value.InstrumentID, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 4, Member Name: ReserveField, Type: uint16
	//
	//
	value.ReserveField, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 5, Member Name: SubBook, Type: TopOfBookMessageSubBook
	//
	//
	value.SubBook, n, err = Read_TopOfBookMessageSubBook(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 6, Member Name: Action, Type: byte
	//
	//
	value.Action, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 7, Member Name: Side, Type: byte
	//
	//
	value.Side, n, err = stream.Read_byte()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 8, Member Name: Price, Type: float64
	//
	//
	value.Price, n, err = stream.Read_mitch_price08()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 9, Member Name: Quantity, Type: uint32
	//
	//
	value.Quantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 10, Member Name: MarketOrderQuantity, Type: uint32
	//
	//
	value.MarketOrderQuantity, n, err = stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	//
	//
	// Index: 11, Member Name: Reserved01, Type: uint16
	//
	//
	value.Reserved01, n, err = stream.Read_uint16()
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value, byteCount, nil
}

// TopOfBookMessage reset
func Reset_TopOfBookMessage(value *TopOfBookMessage) {
	value.Length = 34
	value.MessageType = 113
	value.Nanosecond = 0
	value.InstrumentID = 0
	value.ReserveField = 0
	value.SubBook = TopOfBookMessageSubBook{}
	value.Action = 0
	value.Side = 0
	value.Price = 0.0
	value.Quantity = 0
	value.MarketOrderQuantity = 0
	value.Reserved01 = 0
}

// TopOfBookMessage clone
func Clone_TopOfBookMessage(value *TopOfBookMessage) *TopOfBookMessage {
	result := NewTopOfBookMessage()
	result.Nanosecond = value.Nanosecond
	result.InstrumentID = value.InstrumentID
	result.ReserveField = value.ReserveField
	result.SubBook = value.SubBook
	result.Action = value.Action
	result.Side = value.Side
	result.Price = value.Price
	result.Quantity = value.Quantity
	result.MarketOrderQuantity = value.MarketOrderQuantity
	result.Reserved01 = value.Reserved01
	return result
}

// TopOfBookMessage WriteMessage
func WriteMessage_TopOfBookMessage(stream Streams.IMitchWriter, value *TopOfBookMessage) (int, error) {
	byteCount := 0
	n, err := stream.Write_uint32(0xac36e385)
	if err != nil {
		return 0, err
	}
	byteCount += n
	n, err = Write_TopOfBookMessage(stream, value)
	if err != nil {
		return 0, err
	}
	byteCount += n
	return byteCount, nil
}

// TopOfBookMessage reader
func ReadMessage_TopOfBookMessage(stream Streams.IMitchReader) (*TopOfBookMessage, int, error) {
	byteCount := 0
	typeCode, n, err := stream.Read_uint32()
	if err != nil {
		return nil, 0, err
	}
	if typeCode != 0xac36e385 {
		return nil, 0, fmt.Errorf("typecode mismatch, while reading TopOfBookMessage. Expected 0xac36e385, got 0x%08x", typeCode)
	}
	byteCount += n
	value_TopOfBookMessage, n, err := Read_TopOfBookMessage(stream)
	if err != nil {
		return nil, 0, err
	}
	byteCount += n
	return value_TopOfBookMessage, byteCount, nil
}

// TopOfBookMessage TypeCode
func TypeCode_TopOfBookMessage() uint32 {
	return 0xac36e385
}

// TopOfBookMessage IsTypeCode
func IsTypeCode_TopOfBookMessage(typeCode uint32) bool {
	return typeCode == 0xac36e385
}
