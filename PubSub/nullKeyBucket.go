package PubSub

import "github.com/bhbosman/Application/Messages"

type nullKeyBucket struct {
}

func (self *nullKeyBucket) ExistSubKey(subKey string) bool {
	return false
}

func (self *nullKeyBucket) Key() string {
	return "nullKeyBucket"
}

func (self *nullKeyBucket) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {

}

func (self *nullKeyBucket) Count() int {
	return 0
}

func (self *nullKeyBucket) UnRegister(subKey string, receiverKey string) error {
	return nil
}

func (self *nullKeyBucket) Register(subKey string, receiver ISubKeyBucketReceiver) (IInterConnector, error) {
	return nil, nil
}

func newNullBucket() IKeyBucket {
	return &nullKeyBucket{}
}

func (self *nullKeyBucket) Close() error {

	return nil
}

func (self *nullKeyBucket) Publish(subKey string, waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error {
	return nil
}
