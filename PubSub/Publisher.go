package PubSub

import (
	"fmt"
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
	newKeyBucket          func(key string) (IKeyBucket, error)
	uniqueNumberGenerator UniqueNumber.IGenerator
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

func (self *Publisher) Publish(key string, data interface{}) error {
	bucket := self.FindBucket(key)
	return bucket.Publish(data)
}

func (self *Publisher) Close() error {
	var errResult error
	self.LockScope(func() {
		for key, value := range self.keys {
			err := value.Close()
			if err != nil {
				errResult = multierr.Append(errResult, err)
				self.logger.Printf("Error closing KeyBucket. Key %v. Error: %v", key, err)
			}
		}
	})
	return errResult
}

func (self *Publisher) FindBucket(key string) IKeyBucket {
	var result IKeyBucket = nil
	ok := false
	self.ReadLockScope(func() {
		result, ok = self.keys[key]
	})
	if ok {
		return result
	}
	return self.nullBucket
}

func (self *Publisher) CreateBucket(key string) (IKeyBucket, error) {
	if key == "" {
		return nil, fmt.Errorf("key can not be empty")
	}
	result, err := self.newKeyBucket(key)
	if err != nil {
		return nil, err
	}
	self.LockScope(func() {
		self.keys[key] = result
	})
	return result, nil
}

//func NewBucket(key string) IKeyBucket {
//	return &KeyBucket{
//		routes: list.New(),
//	}
//}

func NewPublisher(logger *log.Logger, uniqueNumberGenerator UniqueNumber.IGenerator) (*Publisher, error) {
	return &Publisher{
		mutex:      sync.RWMutex{},
		logger:     logger,
		keys:       make(map[string]IKeyBucket),
		nullBucket: newNullBucket(),
		newKeyBucket: func(key string) (IKeyBucket, error) {
			return NewKeyBucket(key, logger, uniqueNumberGenerator)
		},
		uniqueNumberGenerator: uniqueNumberGenerator,
	}, nil
}
