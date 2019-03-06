package yacc

import (
	"reflect"
)



type ConstantValue struct {
	Value     interface{}
	Type      reflect.Type
	MaxLength int
}


func (self *ConstantValue) GetValue() interface{} {
	return self.Value
}

func (self *ConstantValue) GetType() reflect.Type {
	return self.Type
}

func (self *ConstantValue) GetMaxLength() int {
	return self.MaxLength
}


