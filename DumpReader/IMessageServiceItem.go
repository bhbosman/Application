package main

import "sync"

type IWaitGroup interface {
	AddOne() error
	Done()error
}


type WaitGroupDefaultImpl struct {
	wg *sync.WaitGroup
}

func NewWaitGroupDefaultImpl(wg *sync.WaitGroup) *WaitGroupDefaultImpl {
	return &WaitGroupDefaultImpl{
		wg: wg,
	}
}

func (self *WaitGroupDefaultImpl) AddOne() error{
	self.wg.Add(1)
	return nil
}

func (self *WaitGroupDefaultImpl) Done() error{
	self.wg.Done()
	return nil
}

type IMessageServiceItem interface {
	IWaitGroup
	Message() (interface{}, error)
}
