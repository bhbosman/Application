package Managers

import "io"

type IMessageFactory interface {
	Message() (interface{}, error)
	MessageType() int
	Length() uint16
}

type IMessageCount interface {
	MessageType() int
	MessageCount() int
}

type IMitchMessageHandlerRegistrar interface {
	ProcessMessage(wg IWaitGroup, messageFactory IMessageFactory, messageSource IMessageSource) error
	RegisterFeed(manager IMitchDataProcessor) error
	GetMessageCounts() []IMessageCount
}

type IMessageSource interface {
	Sequence() uint64
	Source() string
	FeedName() string
}

type IWaitGroup interface {
	AddOne() error
	Done() error
}

type IMessageServiceItem interface {
	IWaitGroup
	IMessageFactory
	MessageSource() IMessageSource
}

type IMessageService interface {
	Push(message IMessageServiceItem) error
}

type IMitchDataProcessor interface {
	IMessageService
	io.Closer
	DeclareInterestInMessages() ([]int, error)
	Process() error
}

type MessageSource struct {
	sequenceNumber uint64
	source         string
	feedName       string
}

func NewMessageSource(sequenceNumber uint64, source string, feedName string) *MessageSource {
	return &MessageSource{
		sequenceNumber: sequenceNumber,
		source:         source,
		feedName:       feedName,
	}
}

func (self MessageSource) Sequence() uint64 {
	return self.sequenceNumber
}

func (self MessageSource) Source() string {
	return self.source
}

func (self MessageSource) FeedName() string {
	return self.feedName
}
