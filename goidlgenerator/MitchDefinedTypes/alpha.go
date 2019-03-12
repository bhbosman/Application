package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchAlpha struct {
	length int64
}

func (self *mitchAlpha) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchAlpha) GetSequenceCount() (bool, int) {
	if self.length == 0 {
		return false, 0
	}
	return true, int(self.length)
}

func (self *mitchAlpha) DefaultValue() string {
	return "\"\""
}

func (self *mitchAlpha) Kind() interfaces.Kind {
	return interfaces.MitchAlpha
}

func (self *mitchAlpha) Predefined() bool {
	return true
}

func (*mitchAlpha) GetName() string {
	return "MitchAlpha"
}
