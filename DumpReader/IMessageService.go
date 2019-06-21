package main

type IMessageService interface {
	Push(message IMessageServiceItem)
}

