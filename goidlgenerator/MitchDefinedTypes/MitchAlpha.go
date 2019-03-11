package MitchDefinedTypes

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type MitchAlpha struct {
	length      int64

}

func (self *MitchAlpha) GetPackageName() (bool, string) {
	return true, "Streams"
}

func (self *MitchAlpha) GetSequenceCount() (bool, int) {
	if self.length == 0 {
		return false, 0
	}
	return true, int(self.length)
}

func (self *MitchAlpha) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *MitchAlpha) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *MitchAlpha) ClearNext() {

}

func (self *MitchAlpha) GetScopeName() string {
	return self.GetName()
}

func (self *MitchAlpha) DefaultValue() string {
	return "\"\""
}

func (self *MitchAlpha) Kind() interfaces.Kind {
	return interfaces.MitchAlpha
}

func (self *MitchAlpha) Predefined() bool {
	return true
}

func (*MitchAlpha) GetName() string {
	return "MitchAlpha"
}

func NewMitchAlpha(length int64) *MitchAlpha {
	return &MitchAlpha{
		length: length,
	}
}
