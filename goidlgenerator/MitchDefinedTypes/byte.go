package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchByte struct {
}

func (self *mitchByte) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchByte) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchByte) DefaultValue() string {
	return "false"
}

func (self *mitchByte) Kind() interfaces.Kind {
	return interfaces.MitchByte
}

func (self *mitchByte) Predefined() bool {
	return true
}

func (*mitchByte) GetName() string {
	return "MitchByte"
}
