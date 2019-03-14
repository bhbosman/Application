package Extansions

import (
	"fmt"
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
)

func TypeValueForDefinedType(DefinedType interfaces.IDefinedType) string {
	value := func() string {
		_, packageName, typeName := DefinedType.GetPackageName()
		if packageName != "" {
			return fmt.Sprintf("%v.%v", packageName, typeName)
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
