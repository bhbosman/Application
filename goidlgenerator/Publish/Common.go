package Publish

import (
	"errors"
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
	Export(outputStream io.Writer, packageName string, declaredTypes []interfaces.IDefinitionDeclaration) error
}

var registrations map[OutputType]IPublish

func init() {
	registrations = make(map[OutputType]IPublish)
}

func Register(outputType OutputType, publish IPublish) {
	registrations[outputType] = publish
}

func HasOutputType(outputType OutputType) (IPublish, error) {
	result, ok := registrations[outputType]
	if ok {
		return result, nil
	}
	return nil, errors.New("Could not find publisher")
}

func PublishOutputType(outputType OutputType, packageName string, declaredTypes []interfaces.IDefinitionDeclaration, writer io.Writer) error {
	result, ok := registrations[outputType]
	if ok {
		err := result.Export(writer, packageName, declaredTypes)
		if err != nil {
			return err
		}
	}
	return nil
}
