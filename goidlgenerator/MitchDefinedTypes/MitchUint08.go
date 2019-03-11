package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchUInt08 struct {

}




func (self *MitchUInt08) GetPackageName() (bool, string) {
	return true, "Streams"
}



func (self *MitchUInt08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt08) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchUInt08) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchUInt08) ClearNext() {

}

func (self *MitchUInt08) GetScopeName() string {
	return self.GetName()
}

func (self *MitchUInt08) DefaultValue() string {
	return "false"
}

func (self *MitchUInt08) Kind() interfaces.Kind {
	return interfaces.MitchUInt08
}

func (self *MitchUInt08) Predefined() bool {
	return true
}

func (*MitchUInt08) GetName() string {
	return "MitchUInt08"
}

func NewMitchUInt08() *MitchUInt08 {
	return &MitchUInt08{}
}
