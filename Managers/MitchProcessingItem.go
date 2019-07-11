package Managers

import "github.com/bhbosman/Application/Messages"

type MitchProcessingItem struct {
	wg            Messages.IWaitGroup
	data          interface{}
	messageSource Messages.IMessageSource
}

func NewMitchProcessingItem(wg Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) *MitchProcessingItem {
	return &MitchProcessingItem{
		wg:            wg,
		data:          data,
		messageSource: messageSource,
	}
}

func (self *MitchProcessingItem) GetWaitGroup() Messages.IWaitGroup {
	return self.wg
}

func (self *MitchProcessingItem) Message() interface{} {
	return self.data
}

func (self *MitchProcessingItem) MessageSource() Messages.IMessageSource {
	return self.messageSource
}

func (self *MitchProcessingItem) AddOne() error {
	return self.wg.AddOne()
}

func (self *MitchProcessingItem) Done() error {
	return self.wg.Done()
}
