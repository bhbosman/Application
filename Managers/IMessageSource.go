package Managers

import (
	"github.com/bhbosman/Application/Messages"
	"io"
)

const MitchFeed = "MitchFeed"

type IMessageCount interface {
	MessageType() int
	MessageCount() int
}

//type IMitchMessageHandlerRegistrar interface {
//	ProcessMessage(wg IWaitGroup, messageFactory IMessageFactory, messageSource IMessageSource) error
//	RegisterFeed(manager IMitchDataProcessor) error
//	GetMessageCounts() []IMessageCount
//}

type IMessageService interface {
	Push(message Messages.IMessageServiceItem) error
}

type IMitchDataProcessor interface {
	IMessageService
	io.Closer
	DeclareInterestInMessages() []int
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
