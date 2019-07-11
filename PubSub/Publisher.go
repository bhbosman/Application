package PubSub

import (
	"fmt"
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/UniqueNumber"
	"go.uber.org/multierr"
	"log"
	"sync"
)

type Publisher struct {
	mutex                 sync.RWMutex
	logger                *log.Logger
	keys                  map[string]IKeyBucket
	nullBucket            IKeyBucket
	newKeyBucket          func(key string) IKeyBucket
	uniqueNumberGenerator UniqueNumber.IGenerator
}

func (self *Publisher) ExistKeySubKey(key string, subkey string) bool {
	var keyValue IKeyBucket
	ok1 := false
	self.ReadLockScope(func() {
		keyValue, ok1 = self.keys[key]
	})
	if !ok1 {
		return false
	}
	return keyValue.ExistSubKey(subkey)
}

func (self *Publisher) ExistKey(key string) bool {
	ok := false
	self.ReadLockScope(func() {
		_, ok = self.keys[key]
	})
	return ok
}

func (self *Publisher) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {
	keys := make([]string, 0, 8)
	for key, value := range self.keys {
		value.UnRegisterReceiver(receiver)
		if value.Count() == 0 {
			keys = append(keys, key)
		}
	}

	for _, key := range keys {
		if value, ok := self.keys[key]; ok {
			if value.Count() == 0 {
				self.LockScope(func() {
					if value.Count() == 0 {
						self.logger.Printf("Unregister Key \"%v\"", value.Key())
						delete(self.keys, key)
					}
				})
			}
		}
	}
}

func NewPublisher(logger *log.Logger, uniqueNumberGenerator UniqueNumber.IGenerator) (*Publisher, error) {
	return &Publisher{
		mutex:      sync.RWMutex{},
		logger:     logger,
		keys:       make(map[string]IKeyBucket),
		nullBucket: newNullBucket(),
		newKeyBucket: func(key string) IKeyBucket {
			return NewKeyBucket(key, logger, uniqueNumberGenerator)
		},
		uniqueNumberGenerator: uniqueNumberGenerator,
	}, nil
}

func (self *Publisher) Register(key string, subKey string, receiver ISubKeyBucketReceiver) (IInterConnector, error) {
	if receiver == nil {
		return nil, fmt.Errorf("no receiver")
	}

	subKeyBucket, err := self.FindCreateAndAddSubKeyBucket(key)
	if err != nil {
		return nil, err
	}
	return subKeyBucket.Register(subKey, receiver)
}

func (self *Publisher) UnRegister(key string, subKey string, receiverKey string) error {
	ok := false
	var keyBucket IKeyBucket
	self.ReadLockScope(func() {
		keyBucket, ok = self.keys[subKey]
	})
	if ok {
		return keyBucket.UnRegister(subKey, receiverKey)
	}
	return nil
}

func (self Publisher) FindCreateAndAddSubKeyBucket(key string) (IKeyBucket, error) {
	result := self.FindSubKeyBucket(key, false)
	if result != nil {
		return result, nil
	}
	return self.CreateAndAddSubKeyBucket(key)
}

func (self *Publisher) FindSubKeyBucket(subKey string, produceNulSubKeyBucket bool) IKeyBucket {
	var result IKeyBucket
	ok := false
	self.ReadLockScope(func() {
		result, ok = self.keys[subKey]
	})
	if ok {
		return result
	}
	if produceNulSubKeyBucket {
		return self.nullBucket
	}
	return nil

}

func (self *Publisher) CreateAndAddSubKeyBucket(key string) (IKeyBucket, error) {
	if key == "" {
		return nil, fmt.Errorf("key can not be empty")
	}
	result := NewKeyBucket(key, self.logger, self.uniqueNumberGenerator)
	self.LockScope(func() {
		self.keys[key] = result
	})
	return result, nil
}

func (self *Publisher) ReadLockScope(cb func()) {
	if cb != nil {
		self.mutex.RLock()
		defer self.mutex.RUnlock()
		cb()
	}
}
func (self *Publisher) LockScope(cb func()) {
	if cb != nil {
		self.mutex.Lock()
		defer self.mutex.Unlock()
		cb()
	}
}

func (self *Publisher) Publish(key string, subKey string, waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error {
	bucket := self.FindBucket(key, true)
	return bucket.Publish(subKey, waitGroup, messageSource, data)
}

func (self *Publisher) Close() error {
	var errResult error
	self.LockScope(func() {
		for key, value := range self.keys {
			err := value.Close()
			if err != nil {
				errResult = multierr.Append(errResult, err)
				self.logger.Printf("Error closing SubKeyBucket. Key %v. Error: %v", key, err)
			}
		}
	})
	return errResult
}

func (self *Publisher) FindBucket(key string, produceNullBucker bool) IKeyBucket {
	var result IKeyBucket
	ok := false
	self.ReadLockScope(func() {
		result, ok = self.keys[key]
	})
	if ok {
		return result
	}
	if produceNullBucker {
		return self.nullBucket
	}
	return nil
}

func (self *Publisher) CreateAndAddKeyBucket(key string) (IKeyBucket, error) {
	if key == "" {
		return nil, fmt.Errorf("key can not be empty")
	}
	result := self.newKeyBucket(key)
	self.LockScope(func() {
		self.keys[key] = result
	})
	return result, nil
}
