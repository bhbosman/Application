package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt64 struct {
}

func (self *mitchUInt64) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *mitchUInt64) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchUInt64) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt64) DefaultValue() string {
	return "0"
}

func (self *mitchUInt64) Kind() interfaces.Kind {
	return interfaces.MitchUInt64
}

func (self *mitchUInt64) Predefined() bool {
	return true
}

func (self *mitchUInt64) GetName() string {
	return self.Kind().String()
}
