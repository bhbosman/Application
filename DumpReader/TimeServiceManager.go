package main

import (
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
)

type TimeServiceManager struct {
	ch     chan IMessageServiceItem
	logger *log.Logger
}

func NewTimeServiceManager(logger *log.Logger) *TimeServiceManager {
	return &TimeServiceManager{
		logger: logger,
	}
}

func (self *TimeServiceManager) DeclareInterestInMessages() ([]byte, error) {
	return []byte{
		GeneratedFiles.TimeMessageMessageType,
	}, nil
}

func (self *TimeServiceManager) Push(message IMessageServiceItem) {
	ch := self.ch
	if ch != nil {
		err := message.AddOne()
		if err != nil {
			return
		}
		ch <- message
	}
}


func (self *TimeServiceManager) processMessage(item IMessageServiceItem) error {
	msg, err := item.Message()
	if err != nil {
		return err
	}
	switch v := msg.(type) {
	case *GeneratedFiles.TimeMessage:
		return self.HandleTimeMessageMessage(v)
	}
	return nil
}

func (self *TimeServiceManager) HandleTimeMessageMessage(message *GeneratedFiles.TimeMessage) error {
	return nil
}
