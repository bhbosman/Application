package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchPrice08 struct {
}

func (self *mitchPrice08) GetStreamFunctionName() string {
	return "mitch_price08"
}

func (self *mitchPrice08) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchPrice08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchPrice08) DefaultValue() string {
	return "0.0"
}

func (self *mitchPrice08) Kind() interfaces.Kind {
	return interfaces.MitchPrice08
}

func (self *mitchPrice08) Predefined() bool {
	return true
}

func (self *mitchPrice08) GetName() string {
	return self.Kind().String()
}
