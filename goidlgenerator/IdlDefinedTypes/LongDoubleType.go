package IdlDefinedTypes

import (
	"encoding/json"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

type LongDoubleType struct {
}

func (self *LongDoubleType) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *LongDoubleType) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *LongDoubleType) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *LongDoubleType) DefaultValue() string {
	return "0.0"
}

func (self *LongDoubleType) GetNext() interfaces.IDefinitionDeclaration {
	return nil
}

func (self *LongDoubleType) SetNext(typeSpec interfaces.IDefinitionDeclaration) {

}

func (self *LongDoubleType) ClearNext() {

}

func (self *LongDoubleType) GetScopeName() string {
	return self.GetName()
}

func (self *LongDoubleType) Kind() interfaces.Kind {
	return interfaces.LongDouble
}
func (self *LongDoubleType) Predefined() bool {
	return true
}

func (*LongDoubleType) GetName() string {
	return "IDLLongDoubleType"
}

func (self *LongDoubleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"Type"`
	}{
		Type: self.GetName(),
	})
}
