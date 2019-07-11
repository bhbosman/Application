package Messages

type IWaitGroup interface {
	AddOne() error
	Done() error
}

type IGetWaitGroup interface {
	GetWaitGroup() IWaitGroup
}
