package main

type IMessageFactory interface {
	Message() (interface{}, error)
	MessageType() byte
	Length() uint16
}
