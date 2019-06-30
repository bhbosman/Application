package main

import (
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
)

type SymbolDirectoryManager struct {
	ch     chan IMessageServiceItem
	logger *log.Logger
}

func NewSymbolDirectoryManager(logger *log.Logger) *SymbolDirectoryManager {
	return &SymbolDirectoryManager{
		logger: logger,
	}
}

func (self *SymbolDirectoryManager) DeclareInterestInMessages() ([]byte, error) {
	return []byte{
		GeneratedFiles.SymbolDirectoryMessageMessageType,
		GeneratedFiles.SymbolStatusMessageMessageType,
	}, nil
}

func (self *SymbolDirectoryManager) Push(message IMessageServiceItem) {
	ch := self.ch
	if ch != nil {
		err := message.AddOne()
		if err != nil {
			return
		}
		ch <- message
	}
}

func (self *SymbolDirectoryManager) HandleSymbolDirectoryMessage(item *GeneratedFiles.SymbolDirectoryMessage) error {
	return nil
}

func (self *SymbolDirectoryManager) processMessage(item IMessageServiceItem) error {
	msg, err := item.Message()
	if err != nil {
		return err
	}
	switch v := msg.(type) {
	case *GeneratedFiles.SymbolDirectoryMessage:
		return self.HandleSymbolDirectoryMessage(v)
	case *GeneratedFiles.SymbolStatusMessage:
		return self.HandleSymbolStatusMessage(v)
	}
	return nil
}

func (self *SymbolDirectoryManager) HandleSymbolStatusMessage(message *GeneratedFiles.SymbolStatusMessage) error {
	return nil
}
