package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt64 struct {
}

func (self *mitchUInt64) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchUInt64) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt64) DefaultValue() string {
	return "false"
}

func (self *mitchUInt64) Kind() interfaces.Kind {
	return interfaces.MitchUInt64
}

func (self *mitchUInt64) Predefined() bool {
	return true
}

func (*mitchUInt64) GetName() string {
	return "MitchUInt64"
}
