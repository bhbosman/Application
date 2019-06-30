package main

type MitchProcessingItem struct {
	wg             IWaitGroup
	messageFactory IMessageFactory
}

func NewMitchProcessingItem(wg IWaitGroup, messageFactory IMessageFactory) *MitchProcessingItem {
	return &MitchProcessingItem{
		wg:             wg,
		messageFactory: messageFactory,
	}
}

func (self *MitchProcessingItem) Message() (interface{}, error) {
	return self.messageFactory.Message()
}

func (self *MitchProcessingItem) AddOne() error{
	return self.wg.AddOne()
}

func (self *MitchProcessingItem) Done()  error{
	return self.wg.Done()
}
