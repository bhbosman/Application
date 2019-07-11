package Messages

type IMessageServiceItem interface {
	IWaitGroup
	IGetWaitGroup
	Message() interface{}
	MessageSource() IMessageSource
}
