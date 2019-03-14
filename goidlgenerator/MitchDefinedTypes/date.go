package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchDate struct {
}

func (self *mitchDate) GetStreamFunctionName() string {
	return "mitch_date"
}

func (self *mitchDate) GetPackageName() (bool, string, string) {
	return true, "time", self.Kind().String()
}

func (self *mitchDate) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchDate) DefaultValue() string {
	_, s, _ := self.GetPackageName()
	name := self.GetName()
	return fmt.Sprintf("%v.%v{}", s, name)
}

func (self *mitchDate) Kind() interfaces.Kind {
	return interfaces.MitchDate
}

func (self *mitchDate) Predefined() bool {
	return true
}

func (self *mitchDate) GetName() string {
	return "date"
}
