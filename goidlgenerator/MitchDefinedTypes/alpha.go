package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchAlpha struct {
	length int64
}

func (self *mitchAlpha) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *mitchAlpha) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
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

func (self *mitchAlpha) GetName() string {
	return self.Kind().String()
}
