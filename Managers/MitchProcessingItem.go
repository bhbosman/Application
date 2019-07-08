package Managers

type MitchProcessingItem struct {
	wg             IWaitGroup
	messageFactory IMessageFactory
	messageSource  IMessageSource
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

func (self *MitchProcessingItem) MessageSource() IMessageSource {
	return self.messageSource
}

func NewMitchProcessingItem(wg IWaitGroup, messageFactory IMessageFactory, messageSource IMessageSource) *MitchProcessingItem {
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
