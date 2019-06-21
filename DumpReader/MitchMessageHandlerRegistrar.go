package main

import (
	"container/list"
	"fmt"
	"log"
	"sync"
)

type MessageTypeRegistrar struct {
	l *list.List
}

func (self *MessageTypeRegistrar) Add(processor IMitchDataProcessor) {
	self.l.PushBack(processor)
}

func (self *MessageTypeRegistrar) ProcessMessage(group *sync.WaitGroup, factory IMessageFactory) error {
	for e := self.l.Front(); e != nil; e = e.Next() {
		mitchDataProcessor, ok := e.Value.(IMitchDataProcessor)
		if !ok {
			return fmt.Errorf("Count not get interface IMitchDataProcessor\n")
		}
		func(mitchDataProcessor IMitchDataProcessor){
			item := NewMitchProcessingItem(group, factory)
			err := item.Add()
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

func (self *MitchMessageHandlerRegistrar) ProcessMessage(wg *sync.WaitGroup, messageFactory IMessageFactory) error {
	handler, ok := self.MessageHandlers[messageFactory.MessageType()]
	if ok {
		return handler.ProcessMessage(wg, messageFactory)
	}
	self.logger.Printf("No message handler for 0x%x\n", messageFactory.MessageType())
	self.MessageHandlers[messageFactory.MessageType()] = NewMessageTypeRegistrar()
	return nil
}

func NewMitchMessageHandlerRegistrar(logger *log.Logger) *MitchMessageHandlerRegistrar {
	return &MitchMessageHandlerRegistrar{
		MessageHandlers: make(map[byte]*MessageTypeRegistrar),
		logger:          logger,
	}
}
