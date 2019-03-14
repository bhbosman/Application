package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type mitchTime struct {
}

func (self *mitchTime) GetStreamFunctionName() string {
	return "mitch_time"
}

func (self *mitchTime) GetPackageName() (bool, string, string) {
	return true, "time", self.Kind().String()
}

func (self *mitchTime) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *mitchTime) DefaultValue() string {
	_, s, _ := self.GetPackageName()
	name := self.GetName()
	return fmt.Sprintf("%v.%v{}", s, name)
}

func (self *mitchTime) Kind() interfaces.Kind {
	return interfaces.MitchTime
}

func (self *mitchTime) Predefined() bool {
	return true
}

func (self *mitchTime) GetName() string {
	return "time"
}
