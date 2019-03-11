package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchUInt16 struct {

}




func (self *MitchUInt16) GetPackageName() (bool, string) {
	return true, "Streams"
}



func (self *MitchUInt16) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt16) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchUInt16) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchUInt16) ClearNext() {

}

func (self *MitchUInt16) GetScopeName() string {
	return self.GetName()
}

func (self *MitchUInt16) DefaultValue() string {
	return "false"
}

func (self *MitchUInt16) Kind() interfaces.Kind {
	return interfaces.MitchUInt16
}

func (self *MitchUInt16) Predefined() bool {
	return true
}

func (*MitchUInt16) GetName() string {
	return "MitchUInt16"
}

func NewMitchUInt16() *MitchUInt16 {
	return &MitchUInt16{}
}
