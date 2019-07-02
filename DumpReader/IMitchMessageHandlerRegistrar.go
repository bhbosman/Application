package main

import (
	"github.com/bhbosman/Application/Managers"
)

type IMessageCount interface {
	MessageType() byte
	MessageCount() int
}

type IMitchMessageHandlerRegistrar interface {
	ProcessMessage(wg IWaitGroup, messageFactory IMessageFactory, messageSource Managers.IMessageSource) error
	RegisterFeed(manager Managers.IMitchDataProcessor) error
	GetMessageCounts() []IMessageCount
}
