package PubSub

import (
	"github.com/bhbosman/Application/Messages"
	"io"
)

type IInterConnector interface {
	io.Closer
	Key() string
	receiver() ISubKeyBucketReceiver
	receiverDescription() string
}

type ISubKeyBucketReceiver interface {
	io.Closer
	Handle(waitGroup Messages.IWaitGroup, data interface{}) error
}
