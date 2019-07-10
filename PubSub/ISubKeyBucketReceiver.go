package PubSub

import (
	"github.com/bhbosman/Application/Messages"
	"io"
)

type ISubKeyBucketReceiver interface {
	io.Closer
	Handle(waitGroup Messages.IWaitGroup, data interface{}) error
}

