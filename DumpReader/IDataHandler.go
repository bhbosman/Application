package main

import "github.com/bhbosman/Application/Managers"

type IDataHandler interface {
	CreateMessageFactory(messageType int, length uint16, stream IStreamData) (Managers.IMessageFactory, error)
}
