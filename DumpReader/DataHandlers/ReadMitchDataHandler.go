package DataHandlers

import (
	"fmt"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/Streams"
)

type ReadMitchDataHandler struct {
	GetReaderFunc func(byteStream []byte) (Streams.IMitchReader, error)
}

func (self ReadMitchDataHandler) CreateAndReadData(messageType byte, length uint16, stream []byte) (interface{}, int, error) {
	reader, err := self.GetReaderFunc(stream)
	if err != nil {
		return nil, 0, err
	}

	if reader == nil {
		return nil, 0, fmt.Errorf("GetReaderFunc is nil")
	}

	data, n, err := GeneratedFiles.CreateAndReadData(messageType, length, reader)
	if err != nil {
		if _, ok := err.(*GeneratedFiles.CreateAndReadDataNotFound); ok {
			return nil, 0, &DataHandlerErrorDidNothing{}
		}
		return nil, 0, err
	}
	return data, n, nil
}

func NewReadMitchDataHandler(reader func(byteStream []byte) (Streams.IMitchReader, error)) (*ReadMitchDataHandler, error) {
	return &ReadMitchDataHandler{
		GetReaderFunc: reader,
	}, nil
}
