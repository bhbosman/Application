package main

type IMessageCount interface {
	MessageType() byte
	MessageCount() int
}

type IMessageSource interface {
	Sequence() int
	Source() string
	FeedName() string
}

type IMitchMessageHandlerRegistrar interface {
	ProcessMessage(wg IWaitGroup, messageFactory IMessageFactory, messageSource IMessageSource) error
	RegisterFeed(manager IMitchDataProcessor) error
	GetMessageCounts() []IMessageCount
}
