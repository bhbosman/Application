package SymbolService

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
)

type SymbolDirectoryManager struct {
	logger *log.Logger
}

func NewSymbolDirectoryManager(logger *log.Logger) (Managers.IMitchDataProcessor, error) {
	return &SymbolDirectoryManager{
		logger: logger,
	}, nil
}

func (self *SymbolDirectoryManager) Close() error {
	return nil
}

func (self *SymbolDirectoryManager) DeclareInterestInMessages() ([]byte, error) {
	return []byte{
		GeneratedFiles.SymbolDirectoryMessageMessageType,
		GeneratedFiles.SymbolStatusMessageMessageType,
	}, nil
}

func (self *SymbolDirectoryManager) Push(message Managers.IMessageServiceItem) error {
	msg, err := message.Message()
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

func (self *SymbolDirectoryManager) HandleSymbolDirectoryMessage(item *GeneratedFiles.SymbolDirectoryMessage) error {
	return nil
}

func (self *SymbolDirectoryManager) HandleSymbolStatusMessage(message *GeneratedFiles.SymbolStatusMessage) error {
	return nil
}
