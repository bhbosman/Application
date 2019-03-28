package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchMessageLength struct {
	value uint16
}

func (self *mitchMessageLength) GetStreamFunctionName() string {
	return self.Kind().String()
}

func newMitchMessageLength(value uint16) *mitchMessageLength {
	return &mitchMessageLength{
		value: value,
	}
}

func (self *mitchMessageLength) GetPackageName() (bool, string, string) {
	return true, "", "uint16"
}

func (self *mitchMessageLength) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchMessageLength) DefaultValue() string {
	return fmt.Sprintf("%v", self.value)
}

func (self *mitchMessageLength) Kind() interfaces.Kind {
	return interfaces.MitchMessageLength
}

func (self *mitchMessageLength) Predefined() bool {
	return true
}

func (self *mitchMessageLength) GetName() string {
	return self.Kind().String()
}
