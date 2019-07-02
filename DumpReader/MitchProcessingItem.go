package main

import (
	"github.com/bhbosman/Application/Managers"
)

type MitchProcessingItem struct {
	wg             IWaitGroup
	messageFactory IMessageFactory
	messageSource  Managers.IMessageSource
}

func (self *MitchProcessingItem) MessageSource() Managers.IMessageSource {
	return self.messageSource
}

func NewMitchProcessingItem(wg IWaitGroup, messageFactory IMessageFactory, messageSource Managers.IMessageSource) *MitchProcessingItem {
	return &MitchProcessingItem{
		wg:             wg,
		messageFactory: messageFactory,
		messageSource:  messageSource,
	}
}

func (self *MitchProcessingItem) Message() (interface{}, error) {
	return self.messageFactory.Message()
}

func (self *MitchProcessingItem) AddOne() error {
	return self.wg.AddOne()
}

func (self *MitchProcessingItem) Done() error {
	return self.wg.Done()
}
