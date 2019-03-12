package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchUInt08 struct {
}

func (self *mitchUInt08) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchUInt08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchUInt08) DefaultValue() string {
	return "false"
}

func (self *mitchUInt08) Kind() interfaces.Kind {
	return interfaces.MitchUInt08
}

func (self *mitchUInt08) Predefined() bool {
	return true
}

func (*mitchUInt08) GetName() string {
	return "MitchUInt08"
}
