
//
// Time
//
struct Time 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Seconds MitchUInt32;
	
};


//
// SystemEvent
//
struct SystemEvent 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	EventCode TEventCode;
	
};


//
// SymbolDirectory
//
struct SymbolDirectory 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	SymbolStatus SymbolDirectorySymbolStatus;
	ISIN MitchAlpha012;
	Symbol MitchAlpha025;
	TIDM MitchAlpha012;
	Segment MitchAlpha006;
	PreviousClosePrice MitchPrice08;
	ExpirationDate MitchDate;
	Underlying MitchAlpha025;
	StrikePrice MitchPrice08;
	OptionType SymbolDirectoryOptionType;
	Issuer MitchAlpha006;
	IssueDate MitchDate;
	Coupon MitchPrice08;
	Flags typedef BitField SymbolDirectoryFlags: ;
	SubBook typedef BitField SymbolDirectorySubBook: ;
	CorporateAction MitchAlpha189;
	Leg1Symbol MitchAlpha025;
	Leg2Symbol MitchAlpha025;
	ContractMultiplier MitchPrice08;
	SettlementMethod SymbolDirectorySettlementMethod;
	InstrumentSubCategory MitchAlpha030;
	
};


//
// SymbolStatus
//
struct SymbolStatus 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	TradingStatus SymbolStatusTradingStatus;
	Flags MitchBitField;
	Reason MitchAlpha004;
	SessionChangeReason SymbolStatusSessionChangeReason;
	NewEndTime MitchTime;
	BookType SymbolStatusBookType;
	
};


//
// AddOrder
//
struct AddOrder 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	OrderId MitchUInt64;
	Side AddOrderSide;
	Quantity MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	Price MitchPrice08;
	Flags typedef BitField AddOrderFlags: ;
	
};


//
// AddAttributedOrder
//
struct AddAttributedOrder 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	OrderID MitchByte;
	Side MitchUInt32;
	Quantity MitchUInt64;
	InstrumentID MitchByte;
	Price MitchUInt32;
	Attribution MitchUInt32;
	Flags MitchBitField;
	
};


//
// OrderDeleted
//
struct OrderDeleted 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	OrderID MitchUInt64;
	
};


//
// OrderModified
//
struct OrderModified 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	OrderID MitchUInt64;
	NewQuantity MitchUInt32;
	NewPrice MitchPrice08;
	Flags MitchBitField;
	
};


//
// OrderBookClear
//
struct OrderBookClear 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	InstrumentID MitchUInt32;
	SubBook MitchUInt08;
	BookType MitchByte;
	
};


//
// OrderExecuted
//
struct OrderExecuted 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	OrderID MitchUInt64;
	ExecutedQuantity MitchUInt32;
	TradeID MitchUInt64;
	LastOptPx MitchPrice08;
	Volatility MitchPrice08;
	UnderlyingReferencePrice MitchPrice08;
	
};


//
// OrderExecutedWithPriceSize
//
struct OrderExecutedWithPriceSize 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	OrderID MitchUInt64;
	ExecutedQuantity MitchUInt32;
	DisplayQuantity MitchUInt32;
	TradeID MitchUInt64;
	Printable MitchByte;
	Price MitchPrice08;
	LastOptPx MitchPrice08;
	Volatility MitchPrice08;
	UnderlyingReferencePrice MitchPrice08;
	
};


//
// Trade
//
struct Trade 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	ExecutedQuantity MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	Price MitchPrice08;
	TradeID MitchUInt64;
	SubBook MitchUInt08;
	Flags MitchBitField;
	TradeSubType MitchAlpha006;
	LastOptPx MitchPrice08;
	Volatility MitchPrice08;
	UnderlyingReferencePrice MitchPrice08;
	
};


//
// AuctionTrade
//
struct AuctionTrade 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	Quantity MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	Price MitchPrice08;
	TradeID MitchUInt64;
	AuctionType MitchByte;
	LastOptPx MitchPrice08;
	Volatility MitchPrice08;
	UnderlyingReferencePrice MitchPrice08;
	
};


//
// OffBookTrade
//
struct OffBookTrade 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	ExecutedQuantity MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	Price MitchPrice08;
	TradeID MitchUInt64;
	OffBookTradeType MitchAlpha006;
	TradeTime MitchTime;
	TradeDate MitchDate;
	LastOptPx MitchPrice08;
	Volatility MitchPrice08;
	UnderlyingReferencePrice MitchPrice08;
	
};


//
// TradeBreak
//
struct TradeBreak 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	TradeID MitchUInt64;
	TradeType MitchByte;
	
};


//
// RecoveryTrade
//
struct RecoveryTrade 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	ExecutedQuantity MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	Price MitchPrice08;
	TradeID MitchUInt64;
	AuctionType MitchByte;
	OffBookRFQTradeType MitchAlpha006;
	TradeTime MitchTime;
	TradeDate MitchDate;
	ActionType MitchByte;
	SubBook MitchUInt08;
	Flags MitchBitField;
	LastOptPx MitchPrice08;
	Volatility MitchPrice08;
	UnderlyingReferencePrice MitchPrice08;
	
};


//
// AuctionInfo
//
struct AuctionInfo 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	PairedQuantity MitchUInt32;
	Reserved MitchUInt32;
	ImbalanceDirection MitchByte;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	Price MitchPrice08;
	AuctionType MitchByte;
	
};


//
// Statistics
//
struct Statistics 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	InstrumentID MitchUInt32;
	Reserved MitchByte;
	Reserved MitchByte;
	StatisticType MitchAlpha006;
	Price MitchPrice08;
	OpenCloseIndicator MitchAlpha006;
	SubBook MitchUInt08;
	
};


//
// ExtendedStatistics
//
struct ExtendedStatistics 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	InstrumentID MitchUInt32;
	HighPrice MitchPrice08;
	LowPrice MitchPrice08;
	VWAP MitchPrice08;
	Volume MitchUInt32;
	Turnover MitchPrice04;
	NumberofTrades MitchUInt32;
	ReservedField MitchAlpha001;
	SubBook MitchUInt08;
	NotionalExposure MitchPrice08;
	NotionalDeltaExposure MitchPrice08;
	OpenInterest MitchPrice08;
	
};


//
// News
//
struct News 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	Time MitchTime;
	Urgency MitchByte;
	Headline MitchAlpha001;
	Text MitchAlpha001;
	Instruments MitchAlpha001;
	Underlyings MitchAlpha001;
	
};


//
// TopOfBook
//
struct TopOfBook 
{
	Length MitchUInt16;
	MessageType MitchByte;
	Nanosecond MitchUInt32;
	instrumentID MitchUInt32;
	ReserveField MitchUInt16;
	SubBook MitchBitField;
	Action MitchByte;
	Side MitchByte;
	Price MitchPrice08;
	Quantity MitchUInt32;
	MarketOrderQuantity MitchUInt32;
	ReservedField MitchUInt32;
	
};

