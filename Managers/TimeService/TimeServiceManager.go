package TimeService

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
)

type TimeServiceManager struct {
	logger *log.Logger
}

func (self *TimeServiceManager) Close() error {
	return nil
}

func (self *TimeServiceManager) DeclareInterestInMessages() ([]byte, error) {
	return []byte{
		GeneratedFiles.TimeMessageMessageType,
	}, nil
}

func (self *TimeServiceManager) Push(message Managers.IMessageServiceItem) error {
	return self.processMessage(message)
}

func (self *TimeServiceManager) processMessage(item Managers.IMessageServiceItem) error {
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

func NewTimeServiceManager(logger *log.Logger) (Managers.IMitchDataProcessor, error) {
	return &TimeServiceManager{
		logger: logger,
	}, nil
}
