package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchPrice08 struct {

}






func (self *MitchPrice08) GetPackageName() (bool, string) {
	return true, "Streams"
}


func (self *MitchPrice08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchPrice08) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchPrice08) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchPrice08) ClearNext() {

}

func (self *MitchPrice08) GetScopeName() string {
	return self.GetName()
}

func (self *MitchPrice08) DefaultValue() string {
	return "false"
}

func (self *MitchPrice08) Kind() interfaces.Kind {
	return interfaces.MitchPrice08
}

func (self *MitchPrice08) Predefined() bool {
	return true
}

func (*MitchPrice08) GetName() string {
	return "MitchPrice08"
}

func NewMitchPrice08() *MitchPrice08 {
	return &MitchPrice08{}
}
