package main

import (
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/Streams"
	"sync"
)

type MessageFactory struct {
	mutex          sync.Mutex
	messageType    byte
	length         uint16
	stream         IStreamData
	factory        Streams.IMitchReaderFactory
	createdMessage interface{}
}

func (self *MessageFactory) MessageType() byte {
	return self.messageType
}

func (self *MessageFactory) Length() uint16 {
	return self.length
}

func (self *MessageFactory) Message() (interface{}, error) {
	if self.createdMessage != nil {
		return self.createdMessage, nil
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	if self.createdMessage != nil {
		return self.createdMessage, nil
	}
	reader, err := self.factory.Create(self.stream.Data())
	if err != nil {
		return nil, err
	}
	self.stream.Close()
	self.stream = nil

	data, _, err := GeneratedFiles.CreateAndReadData(self.messageType, self.length, reader)
	if err != nil {
		if _, ok := err.(*GeneratedFiles.CreateAndReadDataNotFound); ok {
			return nil, &DataHandlerErrorDidNothing{}
		}
		return nil, err
	}
	return data, nil
}

func NewMessageFactory(
	messageType byte,
	length uint16,
	stream IStreamData,
	factory Streams.IMitchReaderFactory) (*MessageFactory, error) {
	return &MessageFactory{
		mutex:          sync.Mutex{},
		messageType:    messageType,
		length:         length,
		stream:         stream,
		factory:        factory,
		createdMessage: nil,
	}, nil
}
