package yacc

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type BitField struct {
	Type       string
	BitField00 string
	BitField01 string
	BitField02 string
	BitField03 string
	BitField04 string
	BitField05 string
	BitField06 string
	BitField07 string
	Next       interfaces.IDefinitionDeclaration
}

func (self *BitField) GetName() string {
	return "BitField"
}

func (self *BitField) GetNext() interfaces.IDefinitionDeclaration {
	return self.Next
}

func (self *BitField) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.Next = typeSpec
}

func (self *BitField) ClearNext() {
	self.Next = nil
}

func (self *BitField) GetScopeName() string {
	return self.GetName()
}

func (self *BitField) MarshalJSON() ([]byte, error) {

	return json.Marshal(&struct {
		Type       string `json:"Type"`
		Identifier string `json:"Identifier"`
		BitField00 string `json:"BitField00"`
		BitField01 string `json:"BitField01"`
		BitField02 string `json:"BitField02"`
		BitField03 string `json:"BitField03"`
		BitField04 string `json:"BitField04"`
		BitField05 string `json:"BitField05"`
		BitField06 string `json:"BitField06"`
		BitField07 string `json:"BitField07"`
	}{
		Type:       "BitField",
		Identifier: "BitField",
		BitField00: self.BitField00,
		BitField01: self.BitField01,
		BitField02: self.BitField02,
		BitField03: self.BitField03,
		BitField04: self.BitField04,
		BitField05: self.BitField05,
		BitField06: self.BitField06,
		BitField07: self.BitField07,
	})
}

func NewBitField(
	bitField00 string,
	bitField01 string,
	bitField02 string,
	bitField03 string,
	bitField04 string,
	bitField05 string,
	bitField06 string,
	bitField07 string) *BitField {

	return &BitField{
		BitField00: bitField00,
		BitField01: bitField01,
		BitField02: bitField02,
		BitField03: bitField03,
		BitField04: bitField04,
		BitField05: bitField05,
		BitField06: bitField06,
		BitField07: bitField07}
}
