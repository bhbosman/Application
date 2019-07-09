package PubSub

import "github.com/bhbosman/Application/Messages"

type nullSubBucket struct {
}

func (self *nullSubBucket) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {

}

func (self *nullSubBucket) Count() int {
	return 1
}

func (self *nullSubBucket) Routes() []IRoute {
	return nil
}

func (self *nullSubBucket) UnRegister(key string) error {
	return nil
}

func (self *nullSubBucket) Register(receiver ISubKeyBucketReceiver) (IInterConnector, error) {
	return nil, nil
}

func (self *nullSubBucket) Close() error {
	return nil
}

func (self *nullSubBucket) Publish(waitGroup Messages.IWaitGroup,data interface{}) error {
	return nil
}

func newSubNullBucket() ISubKeyBucket {
	return &nullSubBucket{}
}
