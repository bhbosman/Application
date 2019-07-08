package PubSub

import (
	"log"
)

type InterConnector struct {
	logger     *log.Logger
	icKey      string
	icReceiver IKeyBucketReceiver
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
		logger:     logger,
		icKey:      key,
		icReceiver: receiver,
	}, nil
}
