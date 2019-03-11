package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchBitField struct {

}



func (self *MitchBitField) GetPackageName() (bool, string) {
	return true, "Streams"
}




func (self *MitchBitField) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchBitField) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchBitField) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchBitField) ClearNext() {

}

func (self *MitchBitField) GetScopeName() string {
	return self.GetName()
}

func (self *MitchBitField) DefaultValue() string {
	return "false"
}

func (self *MitchBitField) Kind() interfaces.Kind {
	return interfaces.MitchBitField
}

func (self *MitchBitField) Predefined() bool {
	return true
}

func (*MitchBitField) GetName() string {
	return "MitchBitField"
}

func NewMitchBitField() *MitchBitField {
	return &MitchBitField{}
}
