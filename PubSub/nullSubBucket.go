package PubSub

import (
	"fmt"
	"github.com/bhbosman/Application/Messages"
	"math"
	"sync/atomic"
)

type nullSubBucket struct {
	count int32
}

func (self *nullSubBucket) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {

}

func (self *nullSubBucket) Count() int {
	return math.MaxInt32
}

func (self *nullSubBucket) Routes() []IRoute {
	return nil
}

func (self *nullSubBucket) UnRegister(key string) error {
	return fmt.Errorf("can not unregister against nullSubBucket")
}

func (self *nullSubBucket) Register(receiver ISubKeyBucketReceiver) (IInterConnector, error) {
	return nil, fmt.Errorf("can not register against nullSubBucket")
}

func (self *nullSubBucket) Close() error {
	return nil
}

func (self *nullSubBucket) Publish(waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error {
	atomic.AddInt32(&self.count, 1)
	return nil
}

func newSubNullBucket() ISubKeyBucket {
	return &nullSubBucket{}
}
