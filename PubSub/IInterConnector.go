package PubSub

import "io"

type IInterConnector interface {
	io.Closer
	key() string
	receiver() IKeyBucketReceiver
}

type IKeyBucketReceiver interface {
	io.Closer
	Handle(data interface{}) error
}
