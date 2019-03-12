package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchDate struct {
}

func (self *mitchDate) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchDate) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchDate) DefaultValue() string {
	return "false"
}

func (self *mitchDate) Kind() interfaces.Kind {
	return interfaces.MitchDate
}

func (self *mitchDate) Predefined() bool {
	return true
}

func (*mitchDate) GetName() string {
	return "MitchDate"
}
