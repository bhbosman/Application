package yacc

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

//publishGo:generate goyacc -o idl.publishGo -p "IdlExpr" idl.y

type Declarator struct {
	identifier   string
	defaultValue *ConstantValue
	next         interfaces.IDeclarator `json:"-"`
}

func (self *Declarator) GetDefaultValue() interfaces.IConstantValue {
	return self.defaultValue
}

func (self *Declarator) ClearNext() {
	self.next = nil
}

func (self *Declarator) GetNext() interfaces.IDeclarator {
	return self.next
}

func (self *Declarator) Identifier() string {
	return self.identifier
}

func (self *Declarator) Next() interfaces.IDeclarator {
	return self.next
}

func (self *Declarator) SetNext(next interfaces.IDeclarator) {
	self.next = next
}

func (receiver *Declarator) String() string {
	return fmt.Sprintf("Declarator => Name: %v\n", receiver.Identifier)
}

func (self *Declarator) MarshalJSON() ([]byte, error) {

	if self.defaultValue == nil {
		return json.Marshal(&struct {
			Type       string `json:"Type"`
			Identifier string `json:"Identifier"`
		}{
			Type:       "Declarator",
			Identifier: self.identifier,
		})
	}

	return json.Marshal(&struct {
		Type                  string      `json:"Type"`
		Identifier            string      `json:"Identifier"`
		DefaultValue          interface{} `json:"DefaultValue"`
		DefaultValueType      string      `json:"DefaultValueType"`
		DefaultValueMaxLength int         `json:"DefaultValueMaxLength"`
	}{
		Type:                  "Declarator",
		Identifier:            self.identifier,
		DefaultValue:          self.defaultValue.Value,
		DefaultValueType:      self.defaultValue.Type.String(),
		DefaultValueMaxLength: self.defaultValue.MaxLength,
	})
}

func NewDeclarator(Identifier string, defaultValue *ConstantValue) *Declarator {
	return &Declarator{
		identifier:   Identifier,
		defaultValue: defaultValue,
		next:         nil,
	}
}
