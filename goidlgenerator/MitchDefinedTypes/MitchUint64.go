package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchUInt64 struct {
}





func (self *MitchUInt64) GetPackageName() (bool, string) {
	return true, "Streams"
}



func (self *MitchUInt64) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt64) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchUInt64) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchUInt64) ClearNext() {

}

func (self *MitchUInt64) GetScopeName() string {
	return self.GetName()
}

func (self *MitchUInt64) DefaultValue() string {
	return "false"
}

func (self *MitchUInt64) Kind() interfaces.Kind {
	return interfaces.MitchUInt64
}

func (self *MitchUInt64) Predefined() bool {
	return true
}

func (*MitchUInt64) GetName() string {
	return "MitchUInt64"
}

func NewMitchUInt64() *MitchUInt64 {
	return &MitchUInt64{}
}
