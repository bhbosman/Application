package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchMessageNumber struct {
	value byte
}

func (self *mitchMessageNumber) GetStreamFunctionName() string {
	return self.Kind().String()
}

func newMitchMessageNumber(value byte) *mitchMessageNumber {
	return &mitchMessageNumber{
		value: value,
	}
}

func (self *mitchMessageNumber) GetPackageName() (bool, string, string) {
	return true, "", "byte"
}

func (self *mitchMessageNumber) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchMessageNumber) DefaultValue() string {
	return fmt.Sprintf("%v", self.value)
}

func (self *mitchMessageNumber) Kind() interfaces.Kind {
	return interfaces.MitchMessageNumber
}

func (self *mitchMessageNumber) Predefined() bool {
	return true
}

func (self *mitchMessageNumber) GetName() string {
	return self.Kind().String()
}
