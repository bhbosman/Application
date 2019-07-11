package PubSub

import (
	"github.com/bhbosman/Application/Messages"
)

type ISubKeyBucketReceiver interface {
	FeedStopped() error
	Handle(waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error
}
