package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

//
//type BitField struct {
//	Identifier string
//	BitField00 string
//	BitField01 string
//	BitField02 string
//	BitField03 string
//	BitField04 string
//	BitField05 string
//	BitField06 string
//	BitField07 string
//	BitsUsed   byte
//	Next       interfaces.IDefinitionDeclaration
//}
//
//func (self *BitField) GetSequenceCount() (bool, int) {
//	return false, 0
//}
//
//func (self *BitField) GetPackageName() (bool, string, string) {
//	return false, ""
//}
//
//func (self *BitField) Kind() interfaces.Kind {
//	return interfaces.BitField
//}
//
//func (self *BitField) DefaultValue() string {
//	return "nil"
//}
//
//func (self *BitField) Predefined() bool {
//	return false
//}
//
//var bitFieldCounter = 0
//
//func getNumber() int {
//	bitFieldCounter += 1
//	return bitFieldCounter
//}
//
//func (self *BitField) GetName() string {
//	return self.Identifier
//}
//
//func (self *BitField) GetNext() interfaces.IDefinitionDeclaration {
//	return self.Next
//}
//
//func (self *BitField) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
//	self.Next = typeSpec
//}
//
//func (self *BitField) ClearNext() {
//	self.Next = nil
//}
//
//func (self *BitField) GetScopeName() string {
//	return self.GetName()
//}
//
//func (self *BitField) MarshalJSON() ([]byte, error) {
//
//	return json.Marshal(&struct {
//		Type       string `json:"Type"`
//		Identifier string `json:"Identifier"`
//		BitField00 string `json:"BitField00"`
//		BitField01 string `json:"BitField01"`
//		BitField02 string `json:"BitField02"`
//		BitField03 string `json:"BitField03"`
//		BitField04 string `json:"BitField04"`
//		BitField05 string `json:"BitField05"`
//		BitField06 string `json:"BitField06"`
//		BitField07 string `json:"BitField07"`
//		BitsUsed   byte   `json:"BitsUsed"`
//	}{
//		Type:       "BitField",
//		Identifier: self.Identifier,
//		BitField00: self.BitField00,
//		BitField01: self.BitField01,
//		BitField02: self.BitField02,
//		BitField03: self.BitField03,
//		BitField04: self.BitField04,
//		BitField05: self.BitField05,
//		BitField06: self.BitField06,
//		BitField07: self.BitField07,
//		BitsUsed:   self.BitsUsed,
//	})
//}
//
//func NewBitField(
//	bitField00 string,
//	bitField01 string,
//	bitField02 string,
//	bitField03 string,
//	bitField04 string,
//	bitField05 string,
//	bitField06 string,
//	bitField07 string) *BitField {
//
//	used := byte(0)
//	if bitField00 != "b0" {
//		used = used | 0x01
//	}
//	if bitField01 != "b1" {
//		used = used | 0x02
//	}
//	if bitField02 != "b2" {
//		used = used | 0x04
//	}
//	if bitField03 != "b3" {
//		used = used | 0x08
//	}
//	if bitField04 != "b4" {
//		used = used | 0x10
//	}
//	if bitField05 != "b5" {
//		used = used | 0x20
//	}
//	if bitField06 != "b6" {
//		used = used | 0x40
//	}
//	if bitField07 != "b7" {
//		used = used | 0x80
//	}
//
//	number := getNumber()
//
//	return &BitField{
//		Identifier: fmt.Sprintf("bitField%04d", number),
//		BitField00: bitField00,
//		BitField01: bitField01,
//		BitField02: bitField02,
//		BitField03: bitField03,
//		BitField04: bitField04,
//		BitField05: bitField05,
//		BitField06: bitField06,
//		BitField07: bitField07,
//		BitsUsed:   used,
//		Next:       nil,
//	}
//}
