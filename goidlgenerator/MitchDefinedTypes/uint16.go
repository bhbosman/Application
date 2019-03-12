package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt16 struct {
}

func (self *mitchUInt16) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchUInt16) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt16) DefaultValue() string {
	return "false"
}

func (self *mitchUInt16) Kind() interfaces.Kind {
	return interfaces.MitchUInt16
}

func (self *mitchUInt16) Predefined() bool {
	return true
}

func (*mitchUInt16) GetName() string {
	return "MitchUInt16"
}
