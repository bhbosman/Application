package Extansions

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

func TypeValueForDefinedType(DefinedType interfaces.IDefinedType) string {
	value := func() string {
		b, s := DefinedType.GetPackageName()
		if b {
			return fmt.Sprintf("%v.%v", s, DefinedType.GetName())
		}
		return DefinedType.GetName()
	}()
	return func(value string) string {
		if DefinedType.Kind() == interfaces.Struct {
			return "*" + value
		}
		return value
	}(value)
}
