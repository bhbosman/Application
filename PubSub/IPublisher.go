package PubSub

import (
	"github.com/bhbosman/Application/Messages"
	"io"
)

type IPublisher interface {
	io.Closer
	ExistKey(key string) bool
	ExistKeySubKey(key string, subkey string) bool
	Publish(key string, subKey string, waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error
	Register(key string, subKey string, receiver ISubKeyBucketReceiver) (IInterConnector, error)
	UnRegister(key string, subKey string, receiverKey string) error
	UnRegisterReceiver(receiver ISubKeyBucketReceiver)
}
