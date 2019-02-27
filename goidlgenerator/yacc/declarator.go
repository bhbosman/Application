package yacc

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/gofintech/goidlgenerator/interfaces"
	"reflect"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type Declarator struct {
	identifier   string
	defaultValue interface{}
	next         interfaces.IDeclarator `json:"-"`
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
		Type             string      `json:"Type"`
		Identifier       string      `json:"Identifier"`
		DefaultValue     interface{} `json:"DefaultValue"`
		DefaultValueType string      `json:"DefaultValueType"`
	}{
		Type:             "Declarator",
		Identifier:       self.identifier,
		DefaultValue:     self.defaultValue,
		DefaultValueType: reflect.TypeOf(self.defaultValue).String(),
	})
}

func NewDeclarator(Identifier string, defaultValue interface{}) *Declarator {
	return &Declarator{
		identifier:   Identifier,
		defaultValue: defaultValue,
		next:         nil,
	}
}
