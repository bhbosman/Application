package yacc

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"reflect"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type constantValue struct {
	value     interface{}
	valueKind reflect.Kind
	maxLength int
}

func newConstantValue(value interface{}, maxLength int) interfaces.IConstantValue {
	return &constantValue{
		value:     value,
		valueKind: reflect.TypeOf(value).Kind(),
		maxLength: maxLength,
	}
}

func newConstantValueWithNoLength(value interface{}) interfaces.IConstantValue {
	return newConstantValue(value, -1)
}

func (self *constantValue) Value() interface{} {
	return self.value
}

func (self *constantValue) ValueKind() reflect.Kind {
	return self.valueKind
}

func (self *constantValue) MaxLength() int {
	return self.maxLength
}
