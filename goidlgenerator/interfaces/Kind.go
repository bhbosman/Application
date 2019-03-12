package interfaces

import "strconv"

type Kind byte

const (
	Invalid Kind = iota
	Sequence
	Native
	Struct
	Uint08
	Uint16
	Uint32
	Uint64
	Int08
	Int16
	Int32
	Int64
	String
	Bool
	WChar
	Char
	Double
	Fixed
	Float
	LongDouble
	Octet
	WideString
	Enum
	BitField
	TypeDeclarator
	MitchAlpha
	MitchBitField
	MitchByte
	MitchDate
	MitchTime
	MitchPrice04
	MitchPrice08
	MitchUInt08
	MitchUInt16
	MitchUInt32
	MitchUInt64
)

// String returns the name of k.
func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var kindNames = []string{
	Invalid:        "Invalid",
	Sequence:       "Sequence",
	Native:         "Native",
	Struct:         "Struct",
	Uint08:         "Uint08",
	Uint16:         "Uint16",
	Uint32:         "Uint32",
	Uint64:         "Uint64",
	Int08:          "Int08",
	Int16:          "Int16",
	Int32:          "Int32",
	Int64:          "Int64",
	String:         "String",
	Bool:           "Bool",
	WChar:          "WChar",
	Char:           "Char",
	Double:         "Double",
	Fixed:          "Fixed",
	Float:          "Float",
	LongDouble:     "LongDouble",
	Octet:          "Octet",
	WideString:     "WideString",
	Enum:           "Enum",
	BitField:       "BitField",
	TypeDeclarator: "TypeDeclarator",
	MitchAlpha:     "MitchAlpha",
	MitchBitField:  "MitchBitField",
	MitchByte:      "MitchByte",
	MitchDate:      "MitchDate",
	MitchTime:      "MitchTime",
	MitchPrice04:   "MitchPrice04",
	MitchPrice08:   "MitchPrice08",
	MitchUInt08:    "MitchUInt08",
	MitchUInt16:    "MitchUInt16",
	MitchUInt32:    "MitchUInt32",
	MitchUInt64:    "MitchUInt64",
}
