package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchUInt32 struct {

}


func (self *MitchUInt32) GetPackageName() (bool, string) {
	return true, "Streams"
}



func (self *MitchUInt32) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt32) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchUInt32) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchUInt32) ClearNext() {

}

func (self *MitchUInt32) GetScopeName() string {
	return self.GetName()
}

func (self *MitchUInt32) DefaultValue() string {
	return "false"
}

func (self *MitchUInt32) Kind() interfaces.Kind {
	return interfaces.MitchUInt32
}

func (self *MitchUInt32) Predefined() bool {
	return true
}

func (*MitchUInt32) GetName() string {
	return "MitchUInt32"
}

func NewMitchUInt32() *MitchUInt32 {
	return &MitchUInt32{}
}
