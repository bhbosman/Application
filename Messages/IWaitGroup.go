package Messages

type IWaitGroup interface {
	AddOne() error
	Done() error
}
