package PubSub

import (
	"fmt"
	"github.com/bhbosman/Application/UniqueNumber"
	"log"
	"sync"
)

type KeyBucket struct {
	key               string                                                                                     `json:"key"`
	logger            *log.Logger                                                                                `json:"-"`
	mutex             sync.RWMutex                                                                               `json:"-"`
	routes            map[string]IInterConnector                                                                 `json:"routes"`
	uniqueNumber      UniqueNumber.IGenerator                                                                    `json:"-"`
	newInterConnector func(key string, receiver IKeyBucketReceiver, logger *log.Logger) (*InterConnector, error) `json:"-"`
}

func (self *KeyBucket) UnRegister(key string) error {
	ok := false
	self.ReadLockScope(func() {
		_, ok = self.routes[key]
	})
	if ok {
		self.LockScope(func() {
			delete(self.routes, key)
		})
	}
	return nil
}

func NewKeyBucket(key string, logger *log.Logger, uniqueNumber UniqueNumber.IGenerator) *KeyBucket {
	return &KeyBucket{
		key:          key,
		logger:       logger,
		mutex:        sync.RWMutex{},
		routes:       make(map[string]IInterConnector),
		uniqueNumber: uniqueNumber,
		newInterConnector: func(key string, receiver IKeyBucketReceiver, logger *log.Logger) (connector *InterConnector, e error) {
			return NewInterConnector(key, receiver, logger)
		},
	}
}

func (self *KeyBucket) Register(receiver IKeyBucketReceiver) (IInterConnector, error) {
	if receiver == nil {
		return nil, fmt.Errorf("no receiver")
	}

	receiverFound := false
	var receiverInterconnect IInterConnector
	self.ReadLockScope(func() {
		for _, value := range self.routes {
			if value.receiver() == receiver {
				receiverFound = true
				receiverInterconnect = value
			}
		}
	})

	if receiverFound {
		return receiverInterconnect, nil
	}

	key := self.uniqueNumber.Next()
	ic, err := self.newInterConnector(key, receiver, self.logger)
	if err != nil {
		return nil, err
	}
	self.LockScope(func() {
		self.routes[key] = ic
	})
	return ic, nil
}

func (self *KeyBucket) Close() error {
	for _, value := range self.routes {
		err := value.Close()
		if err != nil {

		}
	}

	self.LockScope(func() {
		self.routes = make(map[string]IInterConnector)
	})
	return nil
}

func (self *KeyBucket) Publish(data interface{}) error {
	var connectorList []IInterConnector = nil
	self.ReadLockScope(func() {
		n := len(self.routes)
		connectorList = make([]IInterConnector, n, n)
		i := 0
		for _, value := range self.routes {
			connectorList[i] = value
			i++
		}
	})

	if len(connectorList) > 0 {
		errKeys := make([]string, 0, 4)
		for _, connector := range connectorList {
			err := connector.receiver().Handle(data)
			if err != nil {
				self.logger.Println(err)
				errKeys = append(errKeys, connector.key())
			}
		}
		if len(errKeys) > 0 {
			for _, key := range errKeys {
				err := self.UnRegister(key)
				if err != nil {
					self.logger.Printf("Could not Unregister. Error: %v", err)
				}
			}
		}
	}
	return nil
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
