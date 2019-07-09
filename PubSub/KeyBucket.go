package PubSub

import (
	"fmt"
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/UniqueNumber"
	"log"
	"sync"
)

type KeyBucket struct {
	key           string
	subKeys       map[string]ISubKeyBucket
	logger        *log.Logger
	uniqueNumber  UniqueNumber.IGenerator
	nullSubBucket ISubKeyBucket
	mutex         sync.RWMutex
}

func NewKeyBucket(key string, logger *log.Logger, uniqueNumber UniqueNumber.IGenerator) *KeyBucket {
	return &KeyBucket{
		key:           key,
		subKeys:       make(map[string]ISubKeyBucket),
		logger:        logger,
		uniqueNumber:  uniqueNumber,
		nullSubBucket: newSubNullBucket(),
		mutex:         sync.RWMutex{},
	}
}

func (self *KeyBucket) Key() string {
	return self.key
}

func (self *KeyBucket) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {
	keys := make([]string, 0, 10)
	for key, value := range self.subKeys {
		value.UnRegisterReceiver(receiver)
		if value.Count() == 0 {
			keys = append(keys, key)
		}
	}
	if len(keys) > 0 {
		for _, key := range keys {
			if value, ok := self.subKeys[key]; ok {
				self.LockScope(func() {
					if value.Count() == 0 {
						delete(self.subKeys, key)
					}
				})
			}
		}
	}
}

func (self *KeyBucket) Count() int {
	return len(self.subKeys)
}

func (self *KeyBucket) UnRegister(subKey string, receiverKey string) error {
	ok := false
	var subKeyBucket ISubKeyBucket
	self.ReadLockScope(func() {
		subKeyBucket, ok = self.subKeys[subKey]
	})
	if ok {
		err := subKeyBucket.UnRegister(receiverKey)
		if err != nil {

		}
		if subKeyBucket.Count() == 0 {
			self.LockScope(func() {
				delete(self.subKeys, subKey)
			})
		}
	}
	return nil
}

func (self *KeyBucket) FindCreateAndAddSubKeyBucket(subKey string) (ISubKeyBucket, error) {
	result := self.FindSubKeyBucket(subKey, false)
	if result != nil {
		return result, nil
	}
	return self.CreateAndAddSubKeyBucket(subKey)
}

func (self *KeyBucket) FindSubKeyBucket(subKey string, produceNulSubKeyBucket bool) ISubKeyBucket {
	var result ISubKeyBucket
	ok := false
	self.ReadLockScope(func() {
		result, ok = self.subKeys[subKey]
	})
	if ok {
		return result
	}
	if produceNulSubKeyBucket {
		return self.nullSubBucket
	}
	return nil

}

func (self *KeyBucket) CreateAndAddSubKeyBucket(subKey string) (ISubKeyBucket, error) {
	if subKey == "" {
		return nil, fmt.Errorf("key can not be empty")
	}
	self.logger.Printf("Create SubKeyBucket(%v, %v\n)", self.key, subKey)
	result := NewSubKeyBucket(self.key, subKey, self.logger, self.uniqueNumber)
	self.LockScope(func() {
		self.subKeys[subKey] = result
	})
	return result, nil
}

func (self *KeyBucket) Register(subKey string, receiver ISubKeyBucketReceiver) (IInterConnector, error) {
	if receiver == nil {
		return nil, fmt.Errorf("no receiver")
	}

	subKeyBucket, err := self.FindCreateAndAddSubKeyBucket(subKey)
	if err != nil {
		return nil, err
	}
	return subKeyBucket.Register(receiver)
}

func (self *KeyBucket) Close() error {
	for _, value := range self.subKeys {
		err := value.Close()
		if err != nil {

		}
	}

	self.LockScope(func() {
		self.subKeys = make(map[string]ISubKeyBucket)
	})
	return nil
}

func (self *KeyBucket) Publish(subKey string, waitGroup Messages.IWaitGroup, data interface{}) error {
	subKeyBucket := self.FindSubKeyBucket(subKey, true)
	return subKeyBucket.Publish(waitGroup, data)
}

func (self *KeyBucket) LockScope(cb func()) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	if cb != nil {
		cb()
	}
}

func (self *KeyBucket) ReadLockScope(cb func()) {
	self.mutex.RLock()
	defer self.mutex.RUnlock()
	if cb != nil {
		cb()
	}
}
