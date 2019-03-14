package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt32 struct {
}

func (self *mitchUInt32) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *mitchUInt32) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchUInt32) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt32) DefaultValue() string {
	return "0"
}

func (self *mitchUInt32) Kind() interfaces.Kind {
	return interfaces.MitchUInt32
}

func (self *mitchUInt32) Predefined() bool {
	return true
}

func (self *mitchUInt32) GetName() string {
	return self.Kind().String()
}
