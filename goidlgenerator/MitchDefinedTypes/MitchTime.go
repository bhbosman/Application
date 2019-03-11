package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchTime struct {

}





func (self *MitchTime) GetPackageName() (bool, string) {
	return true, "Streams"
}



func (self *MitchTime) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchTime) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchTime) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchTime) ClearNext() {

}

func (self *MitchTime) GetScopeName() string {
	return self.GetName()
}

func (self *MitchTime) DefaultValue() string {
	return "false"
}

func (self *MitchTime) Kind() interfaces.Kind {
	return interfaces.MitchTime
}

func (self *MitchTime) Predefined() bool {
	return true
}

func (*MitchTime) GetName() string {
	return "MitchTime"
}

func NewMitchTime() *MitchTime {
	return &MitchTime{}
}
