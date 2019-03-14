package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt08 struct {
}

func (self *mitchUInt08) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *mitchUInt08) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchUInt08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt08) DefaultValue() string {
	return "0"
}

func (self *mitchUInt08) Kind() interfaces.Kind {
	return interfaces.MitchUInt08
}

func (self *mitchUInt08) Predefined() bool {
	return true
}

func (self *mitchUInt08) GetName() string {
	return self.Kind().String()
}
