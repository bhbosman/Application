package PubSub

import (
	"github.com/bhbosman/Application/Messages"
	"io"
)

type IKeyBucket interface {
	io.Closer
	Publish(subKey string, waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error
	Register(subKey string, receiver ISubKeyBucketReceiver) (IInterConnector, error)
	UnRegister(subKey string, receiverKey string) error
	UnRegisterReceiver(receiver ISubKeyBucketReceiver)
	ExistSubKey(subKey string) bool
	Count() int
	Key() string
}
