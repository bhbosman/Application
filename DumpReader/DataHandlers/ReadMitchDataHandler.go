package DataHandlers

import (
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/Streams"
	"sync"
)

type ReadMitchDataHandler struct {
	factory Streams.IMitchReaderFactory
}
type IMessageFactory interface {
	Message() (interface{}, error)
	MessageType() byte
	Length() uint16
}
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
	factory Streams.IMitchReaderFactory) *MessageFactory {
	return &MessageFactory{
		mutex:          sync.Mutex{},
		messageType:    messageType,
		length:         length,
		stream:         stream,
		factory:        factory,
		createdMessage: nil,
	}
}

func (self ReadMitchDataHandler) CreateMessageFactory(messageType byte, length uint16, stream IStreamData) (IMessageFactory, error) {
	return NewMessageFactory(messageType, length, stream, self.factory), nil
}

func NewReadMitchDataHandler(factory Streams.IMitchReaderFactory) (*ReadMitchDataHandler, error) {
	return &ReadMitchDataHandler{
		factory: factory,
	}, nil
}
