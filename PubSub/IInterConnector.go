package PubSub

import (
	"io"
)

type IInterConnector interface {
	io.Closer
	Key() string
	receiver() ISubKeyBucketReceiver
	receiverDescription() string
}

