package CheckOrderNumbers

import (
	"github.com/bhbosman/Application/Managers"
	"github.com/bhbosman/Application/MitchFiles/GeneratedFiles"
	"log"
)

type CheckOrderNumbers struct {
	logger *log.Logger
}

func NewCheckOrderNumbers(logger *log.Logger) (*CheckOrderNumbers, error) {
	return &CheckOrderNumbers{
		logger: logger,
	}, nil
}

func (self *CheckOrderNumbers) Process() error {
	return nil
}

func (self *CheckOrderNumbers) Push(message Managers.IMessageServiceItem) error {
	msg, err := message.Message()
	if err != nil {
		return err
	}
	switch v := msg.(type) {
	case *GeneratedFiles.TimeMessage:
		return self.HandleTimeMessage(v, message.MessageSource())
	case *GeneratedFiles.AddOrderMessage:
		return nil
	case *GeneratedFiles.OrderDeletedMessage:
		return nil
	case *GeneratedFiles.OrderModifiedMessage:
		return nil
	case *GeneratedFiles.OrderBookClearMessage:
		return nil
	case *GeneratedFiles.OrderExecutedMessage:
		return nil
	case *GeneratedFiles.OrderExecutedWithPriceSizeMessage:
		return nil
	default:
		self.logger.Println("This should never happen02")
		return nil
	}
}

func (self *CheckOrderNumbers) Close() error {
	return nil
}

func (self *CheckOrderNumbers) DeclareInterestInMessages() []int {
	return []int{
		GeneratedFiles.TimeMessage_MessageType,
		GeneratedFiles.AddOrderMessage_MessageType,
		GeneratedFiles.OrderDeletedMessage_MessageType,
		GeneratedFiles.OrderModifiedMessage_MessageType,
		GeneratedFiles.OrderBookClearMessage_MessageType,
		GeneratedFiles.OrderExecutedMessage_MessageType,
		GeneratedFiles.OrderExecutedWithPriceSizeMessage_MessageType,
	}
}

func (self *CheckOrderNumbers) HandleTimeMessage(message *GeneratedFiles.TimeMessage, source Managers.IMessageSource) error {
	return nil
}
