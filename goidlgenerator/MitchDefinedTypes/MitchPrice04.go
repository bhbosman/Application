package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchPrice04 struct {

}





func (self *MitchPrice04) GetPackageName() (bool, string) {
	return true, "Streams"
}


func (self *MitchPrice04) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchPrice04) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchPrice04) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchPrice04) ClearNext() {

}

func (self *MitchPrice04) GetScopeName() string {
	return self.GetName()
}

func (self *MitchPrice04) DefaultValue() string {
	return "false"
}

func (self *MitchPrice04) Kind() interfaces.Kind {
	return interfaces.MitchPrice04
}

func (self *MitchPrice04) Predefined() bool {
	return true
}

func (*MitchPrice04) GetName() string {
	return "MitchPrice04"
}

func NewMitchPrice04() *MitchPrice04 {
	return &MitchPrice04{}
}
