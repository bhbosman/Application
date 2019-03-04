{
	"Type": "native",
	"Identifier": "MitchAlpha001"
}{
	"Type": "native",
	"Identifier": "MitchAlpha004"
}{
	"Type": "native",
	"Identifier": "MitchAlpha006"
}{
	"Type": "native",
	"Identifier": "MitchAlpha012"
}{
	"Type": "native",
	"Identifier": "MitchAlpha025"
}{
	"Type": "native",
	"Identifier": "MitchAlpha030"
}{
	"Type": "native",
	"Identifier": "MitchAlpha189"
}{
	"Type": "native",
	"Identifier": "MitchBitField"
}{
	"Type": "native",
	"Identifier": "MitchByte"
}{
	"Type": "native",
	"Identifier": "MitchDate"
}{
	"Type": "native",
	"Identifier": "MitchTime"
}{
	"Type": "native",
	"Identifier": "MitchPrice04"
}{
	"Type": "native",
	"Identifier": "MitchPrice08"
}{
	"Type": "native",
	"Identifier": "MitchUInt08"
}{
	"Type": "native",
	"Identifier": "MitchUInt16"
}{
	"Type": "native",
	"Identifier": "MitchUInt32"
}{
	"Type": "native",
	"Identifier": "MitchUInt64"
}{
	"Type": "struct",
	"Identifier": "Time",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length",
				"DefaultValue": 7,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType",
				"DefaultValue": 84,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Seconds"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "TEventCode",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "EndOfDay",
			"DefaultValue": "C",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "StartOfDay",
			"DefaultValue": "O",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		}
	]
}{
	"Type": "struct",
	"Identifier": "SystemEvent",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length",
				"DefaultValue": 8,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType",
				"DefaultValue": 83,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "EventCode"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "TEventCode",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "EndOfDay",
						"DefaultValue": "C",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "StartOfDay",
						"DefaultValue": "O",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					}
				]
			}
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "SymbolDirectorySymbolStatus",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "Halted",
			"DefaultValue": "H",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "Suspended",
			"DefaultValue": "S",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "Inactive",
			"DefaultValue": "a",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "SymbolDirectoryOptionType",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "CallOption",
			"DefaultValue": "C",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "PutOption",
			"DefaultValue": "P",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		}
	]
}{
	"DefinedTyped": {
		"Type": "BitField",
		"Identifier": "BitField",
		"BitField00": "InverseOrderBook",
		"BitField01": "b1",
		"BitField02": "b2",
		"BitField03": "b3",
		"BitField04": "b4",
		"BitField05": "b5",
		"BitField06": "b6",
		"BitField07": "b7"
	},
	"Declarator": {
		"Type": "Declarator",
		"Identifier": "SymbolDirectoryFlags"
	}
}{
	"DefinedTyped": {
		"Type": "BitField",
		"Identifier": "BitField",
		"BitField00": "Regular",
		"BitField01": "OffBook",
		"BitField02": "b2",
		"BitField03": "b3",
		"BitField04": "b4",
		"BitField05": "BulletinBoard",
		"BitField06": "NegotiatedTrades",
		"BitField07": "b7"
	},
	"Declarator": {
		"Type": "Declarator",
		"Identifier": "SymbolDirectorySubBook"
	}
}{
	"Type": "EnumDecl",
	"Identifier": "SymbolDirectorySettlementMethod",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "Cash",
			"DefaultValue": "C",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "Physical",
			"DefaultValue": "P",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		}
	]
}{
	"Type": "struct",
	"Identifier": "SymbolDirectory",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType",
				"DefaultValue": 82,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SymbolStatus"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "SymbolDirectorySymbolStatus",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "Halted",
						"DefaultValue": "H",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "Suspended",
						"DefaultValue": "S",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "Inactive",
						"DefaultValue": "a",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					}
				]
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ISIN"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha012"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Symbol"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha025"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TIDM"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha012"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Segment"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "PreviousClosePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ExpirationDate"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchDate"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Underlying"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha025"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "StrikePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OptionType"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "SymbolDirectoryOptionType",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "CallOption",
						"DefaultValue": "C",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "PutOption",
						"DefaultValue": "P",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					}
				]
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Issuer"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "IssueDate"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchDate"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Coupon"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"DefinedTyped": {
					"Type": "BitField",
					"Identifier": "BitField",
					"BitField00": "InverseOrderBook",
					"BitField01": "b1",
					"BitField02": "b2",
					"BitField03": "b3",
					"BitField04": "b4",
					"BitField05": "b5",
					"BitField06": "b6",
					"BitField07": "b7"
				},
				"Declarator": {
					"Type": "Declarator",
					"Identifier": "SymbolDirectoryFlags"
				}
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"DefinedTyped": {
					"Type": "BitField",
					"Identifier": "BitField",
					"BitField00": "Regular",
					"BitField01": "OffBook",
					"BitField02": "b2",
					"BitField03": "b3",
					"BitField04": "b4",
					"BitField05": "BulletinBoard",
					"BitField06": "NegotiatedTrades",
					"BitField07": "b7"
				},
				"Declarator": {
					"Type": "Declarator",
					"Identifier": "SymbolDirectorySubBook"
				}
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "CorporateAction"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha189"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Leg1Symbol"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha025"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Leg2Symbol"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha025"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ContractMultiplier"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SettlementMethod"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "SymbolDirectorySettlementMethod",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "Cash",
						"DefaultValue": "C",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "Physical",
						"DefaultValue": "P",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					}
				]
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentSubCategory"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha030"
			}
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "SymbolStatusTradingStatus",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "Halt",
			"DefaultValue": "H",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "RegularTrading",
			"DefaultValue": "T",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "OpeningAuctionCall",
			"DefaultValue": "a",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "PostClose",
			"DefaultValue": "b",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "MarketClose",
			"DefaultValue": "c",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "ClosingAuctionCall",
			"DefaultValue": "d",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "VolatilityAuctionCall",
			"DefaultValue": "e",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "EODVolumeAuctionCall",
			"DefaultValue": "E",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "ReOpeningAuctionCall",
			"DefaultValue": "f",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "Pause",
			"DefaultValue": "l",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "FuturesCloseOut",
			"DefaultValue": "p",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "ClosingPriceCross",
			"DefaultValue": "s",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "IntraDayAuctionCall",
			"DefaultValue": "u",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "EndTradeReporting",
			"DefaultValue": "v",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "NoActiveSession",
			"DefaultValue": "w",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "EndOfPostClose",
			"DefaultValue": "x",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "StarOofTrading",
			"DefaultValue": "y",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "ClosingPricePublication",
			"DefaultValue": "z",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "SymbolStatusSessionChangeReason",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "ScheduledTransition",
			"DefaultValue": 0,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "ExtendedByMarketOps",
			"DefaultValue": 1,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "ShortenedByMarketOps",
			"DefaultValue": 2,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "MarketOrderImbalance",
			"DefaultValue": 3,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "PriceOutsideRange",
			"DefaultValue": 4,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "CircuitBreakerTripped",
			"DefaultValue": 5,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "Unavailable",
			"DefaultValue": 9,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "SymbolStatusBookType",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "OnBook",
			"DefaultValue": 1,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "OffBook",
			"DefaultValue": 2,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "BulletinBoard",
			"DefaultValue": 9,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		},
		{
			"Type": "Declarator",
			"Identifier": "NegotiatedTrades",
			"DefaultValue": 11,
			"DefaultValueType": "int64",
			"DefaultValueMaxLength": 0
		}
	]
}{
	"Type": "struct",
	"Identifier": "SymbolStatus",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType",
				"DefaultValue": 72,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradingStatus"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "SymbolStatusTradingStatus",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "Halt",
						"DefaultValue": "H",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "RegularTrading",
						"DefaultValue": "T",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "OpeningAuctionCall",
						"DefaultValue": "a",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "PostClose",
						"DefaultValue": "b",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "MarketClose",
						"DefaultValue": "c",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "ClosingAuctionCall",
						"DefaultValue": "d",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "VolatilityAuctionCall",
						"DefaultValue": "e",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "EODVolumeAuctionCall",
						"DefaultValue": "E",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "ReOpeningAuctionCall",
						"DefaultValue": "f",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "Pause",
						"DefaultValue": "l",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "FuturesCloseOut",
						"DefaultValue": "p",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "ClosingPriceCross",
						"DefaultValue": "s",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "IntraDayAuctionCall",
						"DefaultValue": "u",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "EndTradeReporting",
						"DefaultValue": "v",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "NoActiveSession",
						"DefaultValue": "w",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "EndOfPostClose",
						"DefaultValue": "x",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "StarOofTrading",
						"DefaultValue": "y",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "ClosingPricePublication",
						"DefaultValue": "z",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					}
				]
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchBitField"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reason"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha004"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SessionChangeReason"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "SymbolStatusSessionChangeReason",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "ScheduledTransition",
						"DefaultValue": 0,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "ExtendedByMarketOps",
						"DefaultValue": 1,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "ShortenedByMarketOps",
						"DefaultValue": 2,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "MarketOrderImbalance",
						"DefaultValue": 3,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "PriceOutsideRange",
						"DefaultValue": 4,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "CircuitBreakerTripped",
						"DefaultValue": 5,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "Unavailable",
						"DefaultValue": 9,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					}
				]
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "NewEndTime"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchTime"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "BookType"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "SymbolStatusBookType",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "OnBook",
						"DefaultValue": 1,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "OffBook",
						"DefaultValue": 2,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "BulletinBoard",
						"DefaultValue": 9,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					},
					{
						"Type": "Declarator",
						"Identifier": "NegotiatedTrades",
						"DefaultValue": 11,
						"DefaultValueType": "int64",
						"DefaultValueMaxLength": 0
					}
				]
			}
		}
	]
}{
	"Type": "EnumDecl",
	"Identifier": "AddOrderSide",
	"Decls": [
		{
			"Type": "Declarator",
			"Identifier": "BuyOrder",
			"DefaultValue": "B",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		},
		{
			"Type": "Declarator",
			"Identifier": "SellOrder",
			"DefaultValue": "S",
			"DefaultValueType": "string",
			"DefaultValueMaxLength": 1
		}
	]
}{
	"DefinedTyped": {
		"Type": "BitField",
		"Identifier": "BitField",
		"BitField00": "b0",
		"BitField01": "b1",
		"BitField02": "b2",
		"BitField03": "b3",
		"BitField04": "MarketOrder",
		"BitField05": "BulletinBoard",
		"BitField06": "b6",
		"BitField07": "b7"
	},
	"Declarator": {
		"Type": "Declarator",
		"Identifier": "AddOrderFlags"
	}
}{
	"Type": "struct",
	"Identifier": "AddOrder",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType",
				"DefaultValue": 65,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OrderId"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Side"
			},
			"DefinedType": {
				"Type": "EnumDecl",
				"Identifier": "AddOrderSide",
				"Decls": [
					{
						"Type": "Declarator",
						"Identifier": "BuyOrder",
						"DefaultValue": "B",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					},
					{
						"Type": "Declarator",
						"Identifier": "SellOrder",
						"DefaultValue": "S",
						"DefaultValueType": "string",
						"DefaultValueMaxLength": 1
					}
				]
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Quantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"DefinedTyped": {
					"Type": "BitField",
					"Identifier": "BitField",
					"BitField00": "b0",
					"BitField01": "b1",
					"BitField02": "b2",
					"BitField03": "b3",
					"BitField04": "MarketOrder",
					"BitField05": "BulletinBoard",
					"BitField06": "b6",
					"BitField07": "b7"
				},
				"Declarator": {
					"Type": "Declarator",
					"Identifier": "AddOrderFlags"
				}
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "AddAttributedOrder",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType",
				"DefaultValue": 70,
				"DefaultValueType": "int64",
				"DefaultValueMaxLength": 0
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OrderID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Side"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Quantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Attribution"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchBitField"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "OrderDeleted",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OrderID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "OrderModified",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OrderID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "NewQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "NewPrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchBitField"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "OrderBookClear",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "BookType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "OrderExecuted",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OrderID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ExecutedQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LastOptPx"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volatility"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "UnderlyingReferencePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "OrderExecutedWithPriceSize",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OrderID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ExecutedQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "DisplayQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Printable"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LastOptPx"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volatility"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "UnderlyingReferencePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "Trade",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ExecutedQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchBitField"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeSubType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LastOptPx"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volatility"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "UnderlyingReferencePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "AuctionTrade",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Quantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "AuctionType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LastOptPx"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volatility"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "UnderlyingReferencePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "OffBookTrade",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ExecutedQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OffBookTradeType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeTime"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchTime"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeDate"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchDate"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LastOptPx"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volatility"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "UnderlyingReferencePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "TradeBreak",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "RecoveryTrade",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ExecutedQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt64"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "AuctionType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OffBookRFQTradeType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeTime"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchTime"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "TradeDate"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchDate"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ActionType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Flags"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchBitField"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LastOptPx"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volatility"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "UnderlyingReferencePrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "AuctionInfo",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "PairedQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ImbalanceDirection"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "AuctionType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "Statistics",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Reserved"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "StatisticType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OpenCloseIndicator"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha006"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "ExtendedStatistics",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "InstrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "HighPrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "LowPrice"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "VWAP"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Volume"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Turnover"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice04"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "NumberofTrades"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ReservedField"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha001"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "NotionalExposure"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "NotionalDeltaExposure"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "OpenInterest"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "News",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Time"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchTime"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Urgency"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Headline"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha001"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Text"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha001"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Instruments"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha001"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Underlyings"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchAlpha001"
			}
		}
	]
}{
	"Type": "struct",
	"Identifier": "TopOfBook",
	"Members": [
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Length"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MessageType"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Nanosecond"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "instrumentID"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ReserveField"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt16"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "SubBook"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchBitField"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Action"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Side"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchByte"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Price"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchPrice08"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "Quantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "MarketOrderQuantity"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		},
		{
			"Type": "Member",
			"Declarator": {
				"Type": "Declarator",
				"Identifier": "ReservedField"
			},
			"DefinedType": {
				"Type": "native",
				"Identifier": "MitchUInt32"
			}
		}
	]
}