package main

import "sync"

type IMitchMessageHandlerRegistrar interface {
	ProcessMessage(wg *sync.WaitGroup, messageFactory IMessageFactory) error
	RegisterFeed(manager IMitchDataProcessor) error
}

