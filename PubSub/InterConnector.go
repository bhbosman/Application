package PubSub

import (
	"fmt"
	"log"
	"reflect"
)

type InterConnector struct {
	logger       *log.Logger
	IcKey        string                `json:"key"`
	Receiver     ISubKeyBucketReceiver `json:"receiver"`
	ReceiverType string                `json:"receiver_type"`
}

func (self *InterConnector) receiverDescription() string {
	return self.ReceiverType
}

func (self *InterConnector) String() string {
	return fmt.Sprintf("InterConnector: (key: %v, receiver type %v)", self.IcKey, self.ReceiverType)
}

func (self *InterConnector) receiver() ISubKeyBucketReceiver {
	return self.Receiver
}

func (self *InterConnector) Key() string {
	return self.IcKey
}

func (self *InterConnector) Close() error {
	if self.Receiver != nil {
		err := self.Receiver.FeedStopped()
		if err != nil {
			self.logger.Printf("error closing receiver. Error: %v", err)
		}
	}
	return nil
}

func NewInterConnector(key, subKey, ickey string, receiver ISubKeyBucketReceiver, logger *log.Logger) (*InterConnector, error) {
	return &InterConnector{
		logger:       logger,
		IcKey:        ickey,
		Receiver:     receiver,
		ReceiverType: reflect.TypeOf(receiver).String(),
	}, nil
}
