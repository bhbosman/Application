package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchPrice08 struct {
}

func (self *mitchPrice08) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchPrice08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchPrice08) DefaultValue() string {
	return "false"
}

func (self *mitchPrice08) Kind() interfaces.Kind {
	return interfaces.MitchPrice08
}

func (self *mitchPrice08) Predefined() bool {
	return true
}

func (*mitchPrice08) GetName() string {
	return "MitchPrice08"
}
