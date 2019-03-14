package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchByte struct {
}

func (self *mitchByte) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *mitchByte) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *mitchByte) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchByte) DefaultValue() string {
	return "0"
}

func (self *mitchByte) Kind() interfaces.Kind {
	return interfaces.MitchByte
}

func (self *mitchByte) Predefined() bool {
	return true
}

func (self *mitchByte) GetName() string {
	return self.Kind().String()
}
