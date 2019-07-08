package PubSub

type nullBucket struct {
}

func (self *nullBucket) UnRegister(key string) error {
	return nil
}

func (self *nullBucket) Register(receiver IKeyBucketReceiver) (IInterConnector, error) {
	return nil, nil
}

func (self *nullBucket) Close() error {
	return nil
}

func (self *nullBucket) Publish(data interface{}) error {
	return nil
}

func newNullBucket() IKeyBucket {
	return &nullBucket{}
}
