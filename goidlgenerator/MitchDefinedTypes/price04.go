package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchPrice04 struct {
}

func (self *mitchPrice04) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchPrice04) GetStreamFunctionName() string {
	return "mitch_price04"
}

func (self *mitchPrice04) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchPrice04) DefaultValue() string {
	return "0.0"
}

func (self *mitchPrice04) Kind() interfaces.Kind {
	return interfaces.MitchPrice04
}

func (self *mitchPrice04) Predefined() bool {
	return true
}

func (self *mitchPrice04) GetName() string {
	return self.Kind().String()
}
