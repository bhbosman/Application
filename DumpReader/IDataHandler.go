package main

import (
	"github.com/bhbosman/Application/Messages"
)

type IDataHandler interface {
	CreateMessageFactory(messageType int, length uint16, stream IStreamData) (Messages.IMessageFactory, error)
}
