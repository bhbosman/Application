package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchBitField struct {
}

func (self *mitchBitField) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *mitchBitField) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchBitField) DefaultValue() string {
	return "false"
}

func (self *mitchBitField) Kind() interfaces.Kind {
	return interfaces.MitchBitField
}

func (self *mitchBitField) Predefined() bool {
	return true
}

func (*mitchBitField) GetName() string {
	return "MitchBitField"
}
