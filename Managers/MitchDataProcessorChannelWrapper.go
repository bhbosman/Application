package Managers

import (
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/PubSub"
	"strconv"
)

type MitchDataProcessorChannelWrapper struct {
	publisher PubSub.IPublisher
	next      IMitchDataProcessor
	ch        chan Messages.IMessageServiceItem
}

func (self *MitchDataProcessorChannelWrapper) FeedStopped() error {
	return nil
}

func (self *MitchDataProcessorChannelWrapper) Handle(waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error {
	waitGroup.AddOne()
	self.ch <- NewMitchProcessingItem(waitGroup, messageSource, data)

	return nil
}

func (self *MitchDataProcessorChannelWrapper) Process() error {
	return nil
}

func (self *MitchDataProcessorChannelWrapper) Close() error {
	close(self.ch)
	self.publisher.UnRegisterReceiver(self)
	return self.next.Close()
}

func NewMitchDataProcessorChannelWrapper(next IMitchDataProcessor, publisher PubSub.IPublisher) IMitchDataProcessor {
	result := &MitchDataProcessorChannelWrapper{
		publisher: publisher,
		next:      next,
		ch:        make(chan Messages.IMessageServiceItem, 1024),
	}
	msgTypes := next.DeclareInterestInMessages()
	for _, msgType := range msgTypes {

		_, _ = publisher.Register(MitchFeed, strconv.Itoa(msgType), result)
	}
	go func(mitchDataProcessorChannelWrapper *MitchDataProcessorChannelWrapper) {
		for message := range mitchDataProcessorChannelWrapper.ch {
			func(message Messages.IMessageServiceItem) {
				defer func() {
					err := message.Done()
					if err != nil {

					}
				}()
				err := result.next.Push(message)
				if err != nil {
				}
				if len(mitchDataProcessorChannelWrapper.ch) == 0 {
					err := result.next.Process()
					if err != nil {

					}
				}
			}(message)
		}
	}(result)
	return result
}

func (self *MitchDataProcessorChannelWrapper) Push(message Messages.IMessageServiceItem) error {
	err := message.AddOne()
	if err != nil {

	}
	self.ch <- message

	return nil
}

func (self *MitchDataProcessorChannelWrapper) DeclareInterestInMessages() []int {
	return self.next.DeclareInterestInMessages()
}
