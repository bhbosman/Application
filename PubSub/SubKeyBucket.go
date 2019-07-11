package PubSub

import (
	"fmt"
	"github.com/bhbosman/Application/Messages"
	"github.com/bhbosman/Application/UniqueNumber"
	"log"
	"sync"
)

type SubKeyBucket struct {
	key               string
	subKey            string
	logger            *log.Logger
	mutex             sync.RWMutex
	routes            map[string]IInterConnector
	uniqueNumber      UniqueNumber.IGenerator
	newInterConnector func(key, subKey, ickey string, receiver ISubKeyBucketReceiver, logger *log.Logger) (IInterConnector, error)
}

func (self *SubKeyBucket) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {
	keys := make([]string, 0, 10)
	for k, v := range self.routes {
		if v.receiver() == receiver {
			keys = append(keys, k)
		}
	}

	for _, k := range keys {
		self.LockScope(func() {
			delete(self.routes, k)
		})
	}

}

func (self *SubKeyBucket) Count() int {
	return len(self.routes)
}

func NewSubKeyBucket(key, subKey string, logger *log.Logger, uniqueNumber UniqueNumber.IGenerator) *SubKeyBucket {
	return &SubKeyBucket{
		key:          key,
		subKey:       subKey,
		logger:       logger,
		mutex:        sync.RWMutex{},
		routes:       make(map[string]IInterConnector),
		uniqueNumber: uniqueNumber,
		newInterConnector: func(key, subKey, icKey string, receiver ISubKeyBucketReceiver, logger *log.Logger) (connector IInterConnector, e error) {
			return NewInterConnector(key, subKey, icKey, receiver, logger)
		},
	}
}

func (self *SubKeyBucket) Routes() []IRoute {
	var result []IRoute
	self.ReadLockScope(func() {
		result = make([]IRoute, len(self.routes))
		for _, value := range self.routes {
			result = append(result, newRouteDescription(value.Key(), value.receiverDescription()))
		}
	})
	return result
}

func (self *SubKeyBucket) UnRegister(key string) error {
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

func (self *SubKeyBucket) Register(receiver ISubKeyBucketReceiver) (IInterConnector, error) {
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

	key, _ := self.uniqueNumber.Next()
	ic, err := self.newInterConnector(self.key, self.subKey, key, receiver, self.logger)
	if err != nil {
		return nil, err
	}
	self.LockScope(func() {
		self.routes[key] = ic
	})
	self.logger.Printf("SubKeyBucket(%v, %v) Register receiver(%v). Id: %v\n", self.key, self.subKey, ic.receiverDescription(), ic.Key())
	return ic, nil
}

func (self *SubKeyBucket) Close() error {
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

func (self *SubKeyBucket) Publish(waitGroup Messages.IWaitGroup, messageSource Messages.IMessageSource, data interface{}) error {
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
			err := connector.receiver().Handle(waitGroup, messageSource, data)
			if err != nil {
				self.logger.Println(err)
				errKeys = append(errKeys, connector.Key())
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

func (self *SubKeyBucket) LockScope(cb func()) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	if cb != nil {
		cb()
	}
}

func (self *SubKeyBucket) ReadLockScope(cb func()) {
	self.mutex.RLock()
	defer self.mutex.RUnlock()
	if cb != nil {
		cb()
	}
}
