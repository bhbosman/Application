package PubSub

import (
	"fmt"
	"log"
	"reflect"
)

type InterConnector struct {
	logger         *log.Logger
	icKey          string             `json:"key"`
	icReceiver     IKeyBucketReceiver `json:"receiver"`
	icReceiverType string             `json:"receiver_type"`
}

func (self *InterConnector) String() string {
	return fmt.Sprintf("InterConnector: (key: %v, receiver type %v)", self.icKey, self.icReceiverType)
}

func (self *InterConnector) receiver() IKeyBucketReceiver {
	return self.icReceiver
}

func (self *InterConnector) key() string {
	return self.icKey
}

func (self *InterConnector) Close() error {
	if self.icReceiver != nil {
		err := self.icReceiver.Close()
		if err != nil {
			self.logger.Printf("error closing receiver. Error: %v", err)
		}
	}

	return nil
}

func NewInterConnector(key string, receiver IKeyBucketReceiver, logger *log.Logger) (*InterConnector, error) {
	return &InterConnector{
		logger:         logger,
		icKey:          key,
		icReceiver:     receiver,
		icReceiverType: reflect.TypeOf(receiver).String(),
	}, nil
}
