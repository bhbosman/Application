package Publish

import (
	"github.com/bhbosman/Application/goidlgenerator/interfaces"
	"io"
)

type OutputType int

const (
	Unknown OutputType = iota
	Json
	Go
)

func ToOutputType(s string) OutputType {
	switch s {
	case "json":
		return Json
	case "go":
		return Go
	default:
		return Unknown
	}
}

type IPublish interface {
	Export(outputStream io.Writer, declaredTypes []interfaces.IDefinitionDeclaration)
}

var registrations map[OutputType]IPublish

func init() {
	registrations = make(map[OutputType]IPublish)
}

func Register(outputType OutputType, publish IPublish) {
	registrations[outputType] = publish
}

func HasOutputType(outputType OutputType) IPublish {
	result, ok := registrations[outputType]
	if ok {
		return result
	}
	return nil
}

func PublishOutputType(outputType OutputType, declaredTypes []interfaces.IDefinitionDeclaration, writer io.Writer) error {
	result, ok := registrations[outputType]
	if ok {
		result.Export(writer, declaredTypes)
	}
	return nil
}
