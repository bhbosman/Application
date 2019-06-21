package main

import "sync"

type IMessageCount interface {
	MessageType() byte
	MessageCount() int

}

type IMitchMessageHandlerRegistrar interface {
	ProcessMessage(wg *sync.WaitGroup, messageFactory IMessageFactory) error
	RegisterFeed(manager IMitchDataProcessor) error
	GetMessageCounts()[] IMessageCount
}

