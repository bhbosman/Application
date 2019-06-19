package main

import (
	"github.com/bhbosman/Application/Streams"
)

type ReadMitchDataHandler struct {
	factory Streams.IMitchReaderFactory
}

func (self ReadMitchDataHandler) CreateMessageFactory(messageType byte, length uint16, stream IStreamData) (IMessageFactory, error) {
	return NewMessageFactory(messageType, length, stream, self.factory)
}

func NewReadMitchDataHandler(factory Streams.IMitchReaderFactory) (*ReadMitchDataHandler, error) {
	return &ReadMitchDataHandler{
		factory: factory,
	}, nil
}
