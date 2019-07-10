package Messages

type IMessageServiceItem interface {
	IWaitGroup
	IMessageFactory
	MessageSource() IMessageSource
}
