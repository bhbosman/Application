package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt16 struct {
}

func (self *mitchUInt16) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *mitchUInt16) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchUInt16) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt16) DefaultValue() string {
	return "0"
}

func (self *mitchUInt16) Kind() interfaces.Kind {
	return interfaces.MitchUInt16
}

func (self *mitchUInt16) Predefined() bool {
	return true
}

func (self *mitchUInt16) GetName() string {
	return self.Kind().String()
}
