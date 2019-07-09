package Managers

import "github.com/bhbosman/Application/Messages"

type MitchProcessingItem struct {
	wg             Messages.IWaitGroup
	messageFactory Messages.IMessageFactory
	messageSource  Messages.IMessageSource
}

func (self *MitchProcessingItem) Message() (interface{}, error) {
	return self.messageFactory.Message()
}

func (self *MitchProcessingItem) MessageType() int {
	return self.messageFactory.MessageType()
}

func (self *MitchProcessingItem) Length() uint16 {
	return self.messageFactory.Length()
}

func (self *MitchProcessingItem) MessageSource() Messages.IMessageSource {
	return self.messageSource
}

func NewMitchProcessingItem(wg Messages.IWaitGroup, messageFactory Messages.IMessageFactory, messageSource Messages.IMessageSource) *MitchProcessingItem {
	return &MitchProcessingItem{
		wg:             wg,
		messageFactory: messageFactory,
		messageSource:  messageSource,
	}
}

func (self *MitchProcessingItem) AddOne() error {
	return self.wg.AddOne()
}

func (self *MitchProcessingItem) Done() error {
	return self.wg.Done()
}
