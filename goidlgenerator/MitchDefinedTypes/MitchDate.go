package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchDate struct {

}




func (self *MitchDate) GetPackageName() (bool, string) {
	return true, "Streams"
}



func (self *MitchDate) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchDate) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchDate) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchDate) ClearNext() {

}

func (self *MitchDate) GetScopeName() string {
	return self.GetName()
}

func (self *MitchDate) DefaultValue() string {
	return "false"
}

func (self *MitchDate) Kind() interfaces.Kind {
	return interfaces.MitchDate
}

func (self *MitchDate) Predefined() bool {
	return true
}

func (*MitchDate) GetName() string {
	return "MitchDate"
}

func NewMitchDate() *MitchDate {
	return &MitchDate{}
}
