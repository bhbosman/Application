package TimeService

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
)

type Manager struct {
	logger *log.Logger
}

func (self *Manager) Close() error {
	return nil
}

func (self *Manager) DeclareInterestInMessages() []int {
	return []int{
		GeneratedFiles.TimeMessage_MessageType,
	}
}

func (self *Manager) Push(message Messages.IMessageServiceItem) error {
	return self.processMessage(message)
}

func (self *Manager) processMessage(item Messages.IMessageServiceItem) error {
	msg, err := item.Message()
	if err != nil {
		return err
	}

	switch v := msg.(type) {
	case *GeneratedFiles.TimeMessage:
		return self.HandleTimeMessageMessage(v)
	default:
		self.logger.Println("This should never happen01")
		return nil
	}
}

func (self *Manager) HandleTimeMessageMessage(message *GeneratedFiles.TimeMessage) error {
	return nil
}

func (self *Manager) Process() error {
	return nil
}

func NewTimeServiceManager(logger *log.Logger) (Managers.IMitchDataProcessor, error) {
	return &Manager{
		logger: logger,
	}, nil
}
