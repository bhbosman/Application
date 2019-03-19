package yacc

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type constantValue struct {
	value     interface{}
	valueKind interfaces.Kind
	maxLength int
}

func GetExportValue(value interfaces.IConstantValue) string {
	if value.ValueKind() == interfaces.String {
		if value.MaxLength() == 1 {
			return fmt.Sprintf("'%v'", value.Value())
		}
		return fmt.Sprintf("\"%v\"", value.Value())
	}

	if value.ValueKind() == interfaces.Char {
		return fmt.Sprintf("'%v'", value.Value())
	}

	return fmt.Sprintf("%v", value.Value())
}


func (self *constantValue) GetExportValue() string {
	return GetExportValue(self)
}

func newConstantValue(value interface{}, valueKind interfaces.Kind, maxLength int) interfaces.IConstantValue {
	return &constantValue{
		value:     value,
		valueKind: valueKind,
		maxLength: maxLength,
	}
}

func newConstantValueWithNoLength(value interface{}, valueKind interfaces.Kind) interfaces.IConstantValue {
	return newConstantValue(value, valueKind, -1)
}

func (self *constantValue) Value() interface{} {
	return self.value
}

func (self *constantValue) ValueKind() interfaces.Kind {
	return self.valueKind
}

func (self *constantValue) MaxLength() int {
	return self.maxLength
}
