package Managers

import "io"

type IMessageSource interface {
	Sequence() int
	Source() string
	FeedName() string
}

type IWaitGroup interface {
	AddOne() error
	Done() error
}

type IMessageServiceItem interface {
	IWaitGroup
	MessageSource() IMessageSource
	Message() (interface{}, error)
}


type IMessageService interface {
	Push(message IMessageServiceItem) error
}


type IMitchDataProcessor interface {
	IMessageService
	io.Closer
	DeclareInterestInMessages() ([]byte, error)
}
