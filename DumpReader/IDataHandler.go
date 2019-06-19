package main

type IDataHandler interface {
	CreateMessageFactory(messageType byte, length uint16, stream IStreamData) (IMessageFactory, error)
}
