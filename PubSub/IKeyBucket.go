package PubSub

import "io"

type IKeyBucket interface {
	io.Closer
	Publish(data interface{}) error
	Register(receiver IKeyBucketReceiver) (IInterConnector, error)
	UnRegister(key string) error
}
