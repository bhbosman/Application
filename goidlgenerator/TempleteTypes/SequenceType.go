package TempleteTypes

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type SequenceType struct {
	definedTyped interfaces.IDefinedType
	length       int64
	next         interfaces.IDefinitionDeclaration
}

func (self *SequenceType) GetStreamFunctionName() string {
	return "(seq oops)"
}

func (self *SequenceType) GetPackageName() (bool, string, string) {
	return false, "", ""
}

func (self *SequenceType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *SequenceType) Predefined() bool {
	return false
}

func (self *SequenceType) Kind() interfaces.Kind {
	return interfaces.Sequence
}

func (self *SequenceType) DefaultValue() string {
	return "nil"
}

func (self *SequenceType) Length() int64 {
	return self.length
}

func (self *SequenceType) DefinedTyped() interfaces.IDefinedType {
	return self.definedTyped
}

func (self *SequenceType) GetNext() interfaces.IDefinitionDeclaration {
	return self.next
}

func (self *SequenceType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.next = typeSpec
}

func (self *SequenceType) ClearNext() {
	self.next = nil
}

func (self *SequenceType) GetName() string {
	return fmt.Sprintf("sequence<%v, %v>", self.definedTyped.GetName(), self.length)
}

func NewSequenceType(definedTyped interfaces.IDefinedType, length int64) *SequenceType {
	return &SequenceType{
		definedTyped: definedTyped,
		length:       length,
		next:         nil,
	}
}
