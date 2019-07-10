package Messages


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

