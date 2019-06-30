package main

import (
	"container/list"
	"fmt"
	"log"

)

type MessageTypeRegistrar struct {
	l            *list.List
	messageCount int
}

func (self *MessageTypeRegistrar) Add(processor IMitchDataProcessor) {
	self.l.PushBack(processor)
}

func (self *MessageTypeRegistrar) ProcessMessage(group IWaitGroup, factory IMessageFactory) error {
	self.messageCount++
	for e := self.l.Front(); e != nil; e = e.Next() {
		mitchDataProcessor, ok := e.Value.(IMitchDataProcessor)
		if !ok {
			return fmt.Errorf("Count not get interface IMitchDataProcessor\n")
		}
		func(mitchDataProcessor IMitchDataProcessor) {
			item := NewMitchProcessingItem(group, factory)
			err := item.AddOne()
			if err != nil {

			}
			defer func() {
				err := item.Done()
				if err != nil {

				}
			}()
			mitchDataProcessor.Push(item)
		}(mitchDataProcessor)
	}
	return nil
}

func NewMessageTypeRegistrar() *MessageTypeRegistrar {
	return &MessageTypeRegistrar{
		l: list.New(),
	}
}

type MitchMessageHandlerRegistrar struct {
	MessageHandlers map[byte]*MessageTypeRegistrar
	logger          *log.Logger
}

type messageCountImpl struct {
	messageType byte
	count       int
}

func (self *messageCountImpl) MessageType() byte {
	return self.messageType
}

func (self *messageCountImpl) MessageCount() int {
	return self.count
}

func (self *MitchMessageHandlerRegistrar) GetMessageCounts() []IMessageCount {
	var result []IMessageCount
	for key, value := range self.MessageHandlers {
		result = append(result, &messageCountImpl{
			messageType: key,
			count:       value.messageCount,
		})
	}
	return result
}

func (self *MitchMessageHandlerRegistrar) RegisterFeed(manager IMitchDataProcessor) error {
	messages, err := manager.DeclareInterestInMessages()
	if err != nil {
		return err
	}
	for _, messageNumber := range messages {
		messageTypeRegistrar, ok := self.MessageHandlers[messageNumber]
		if !ok {
			messageTypeRegistrar = NewMessageTypeRegistrar()
			self.MessageHandlers[messageNumber] = messageTypeRegistrar
		}
		messageTypeRegistrar.Add(manager)

	}
	return nil
}

func (self *MitchMessageHandlerRegistrar) ProcessMessage(wg IWaitGroup, messageFactory IMessageFactory, source IMessageSource) error {
	handler, ok := self.MessageHandlers[messageFactory.MessageType()]
	if !ok {

		self.logger.Printf("No message handler for 0x%x\n", messageFactory.MessageType())
		handler = NewMessageTypeRegistrar()
		self.MessageHandlers[messageFactory.MessageType()] = handler
	}
	return handler.ProcessMessage(wg, messageFactory)

}

func NewMitchMessageHandlerRegistrar(logger *log.Logger) *MitchMessageHandlerRegistrar {
	return &MitchMessageHandlerRegistrar{
		MessageHandlers: make(map[byte]*MessageTypeRegistrar),
		logger:          logger,
	}
}
