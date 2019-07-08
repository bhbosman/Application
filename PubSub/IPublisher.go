package PubSub

import "io"

type IPublisher interface {
	io.Closer
	Publish(key string, data interface{}) error
}

type IPublisherRegistration interface {
	Register(key string, receiver IKeyBucketReceiver)
	UnRegister(key string, receiverKey string)
}
