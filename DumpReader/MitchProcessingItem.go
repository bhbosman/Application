package main

import (
	"sync"
	"sync/atomic"
)

type MitchProcessingItem struct {
	wg             *sync.WaitGroup
	count          int32
	messageFactory IMessageFactory
}

func NewMitchProcessingItem(wg *sync.WaitGroup, messageFactory IMessageFactory) *MitchProcessingItem {
	return &MitchProcessingItem{
		wg:             wg,
		messageFactory: messageFactory,
	}
}

func (self *MitchProcessingItem) Message() (interface{}, error) {
	return self.messageFactory.Message()
}

func (self *MitchProcessingItem) Add() error {
	self.wg.Add(1)
	atomic.AddInt32(&self.count, 1)
	return nil
}

func (self *MitchProcessingItem) Done() error {
	self.wg.Done()
	count := atomic.AddInt32(&self.count, -1)
	if count == 0 {

	}
	return nil
}
