package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchTime struct {
}

func (self *mitchTime) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchTime) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchTime) DefaultValue() string {
	return "false"
}

func (self *mitchTime) Kind() interfaces.Kind {
	return interfaces.MitchTime
}

func (self *mitchTime) Predefined() bool {
	return true
}

func (*mitchTime) GetName() string {
	return "MitchTime"
}
