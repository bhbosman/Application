package Messages

type IWaitGroup interface {
	AddOne() error
	Done() error
}
type NullWaitGroup struct {
	
}

func NewNullWaitGroup() *NullWaitGroup {
	return &NullWaitGroup{}
}

func (self *NullWaitGroup) AddOne() error {
	return nil
}

func (self *NullWaitGroup) Done() error {
	return nil
}

type IMessageFactory interface {
	Message() (interface{}, error)
	MessageType() int
	Length() uint16
}

type IMessageSource interface {
	Sequence() uint64
	Source() string
	FeedName() string
}

type IMessageServiceItem interface {
	IWaitGroup
	IMessageFactory
	MessageSource() IMessageSource
}
