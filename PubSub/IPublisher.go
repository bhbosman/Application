package PubSub

import "io"

type IPublisher interface {
	io.Closer
	Publish(key string, data interface{}) error
}
