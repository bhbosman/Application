package main

import "github.com/bhbosman/Application/Managers"

type MitchDataProcessorChannelWrapper struct {
	next Managers.IMitchDataProcessor
	ch   chan Managers.IMessageServiceItem
}

func (self *MitchDataProcessorChannelWrapper) Close() error {
	close(self.ch)
	return self.next.Close()
}

func NewMitchDataProcessorChannelWrapper(next Managers.IMitchDataProcessor) (Managers.IMitchDataProcessor, error) {
	result := &MitchDataProcessorChannelWrapper{
		next: next,
		ch:   make(chan Managers.IMessageServiceItem, 1024),
	}
	go func(mitchDataProcessorChannelWrapper *MitchDataProcessorChannelWrapper) {
		for message := range mitchDataProcessorChannelWrapper.ch {
			func(message Managers.IMessageServiceItem) {
				defer func() {
					err := message.Done()
					if err != nil {

					}
				}()
				err := result.next.Push(message)
				if err != nil {

				}
			}(message)
		}
	}(result)
	return result, nil
}

func (self *MitchDataProcessorChannelWrapper) Push(message Managers.IMessageServiceItem) error {
	err := message.AddOne()
	if err != nil {

	}
	self.ch <- message

	return nil
}

func (self *MitchDataProcessorChannelWrapper) DeclareInterestInMessages() ([]byte, error) {
	return self.next.DeclareInterestInMessages()
}
