package StringDirectory

import "hash/crc32"

type StringRep struct {
	v uint32
}

func NewStringRep(s string) *StringRep {

	v := crc32.ChecksumIEEE([]byte(s))


	return &StringRep{v:v}
}
