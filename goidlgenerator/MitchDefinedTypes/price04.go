package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchPrice04 struct {
}

func (self *mitchPrice04) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchPrice04) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchPrice04) DefaultValue() string {
	return "false"
}

func (self *mitchPrice04) Kind() interfaces.Kind {
	return interfaces.MitchPrice04
}

func (self *mitchPrice04) Predefined() bool {
	return true
}

func (*mitchPrice04) GetName() string {
	return "MitchPrice04"
}
