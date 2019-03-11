package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchByte struct {

}





func (self *MitchByte) GetPackageName() (bool, string) {
	return true, "Streams"
}


func (self *MitchByte) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchByte) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchByte) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchByte) ClearNext() {

}

func (self *MitchByte) GetScopeName() string {
	return self.GetName()
}

func (self *MitchByte) DefaultValue() string {
	return "false"
}

func (self *MitchByte) Kind() interfaces.Kind {
	return interfaces.MitchByte
}

func (self *MitchByte) Predefined() bool {
	return true
}

func (*MitchByte) GetName() string {
	return "MitchByte"
}

func NewMitchByte() *MitchByte {
	return &MitchByte{}
}
