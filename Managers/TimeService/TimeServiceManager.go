package TimeService

import (
	"fmt"
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"github.com/bhbosman/Application/PubSub"
	"log"
)

type Manager struct {
	logger    *log.Logger
	publisher PubSub.IPublisher
}

func NewTimeServiceManager(logger *log.Logger, publisher PubSub.IPublisher) (Managers.IMitchDataProcessor, error) {
	return &Manager{
		logger:    logger,
		publisher: publisher,
	}, nil
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
	msg := message.Message()
	return self.processMessage(0, msg)
}

func (self *Manager) processMessage(recursiveCount int, message interface{}) error {
	if recursiveCount > 1 {
		return fmt.Errorf("recursive error")
	}

	if message == nil {
		return fmt.Errorf("message can not be nil")
	}

	switch v := message.(type) {
	case Messages.IMessageFactory:
		msg, e := v.Message()
		if e != nil {
			return e
		}
		return self.processMessage(recursiveCount+1, msg)
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

