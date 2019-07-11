package PubSub

import (
	"github.com/bhbosman/Application/Messages"
	"io"
)

type ISubKeyBucket interface {
	io.Closer
	Publish(waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error
	Register(receiver ISubKeyBucketReceiver) (IInterConnector, error)
	UnRegister(key string) error
	UnRegisterReceiver(receiver ISubKeyBucketReceiver)
	Routes() []IRoute
	Count() int
}
