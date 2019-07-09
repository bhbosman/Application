package UniqueNumber

import (
	"fmt"
	"io"
	"log"
	"sync/atomic"
	"time"
)

type Closer struct {
	closed bool
}

func (self *Closer) IsClosed() bool {
	return self.closed
}

func (self *Closer) Close() error {
	self.closed = true
	return nil
}

var CloseError = fmt.Errorf("close")

func (self *Closer) CloseError() error {
	return CloseError
}

type IRefCount interface {
	io.Closer
	Duplicate() error
}
type RefCount struct {
	Closer
	refCount int32
}

func (self *RefCount) Duplicate() error {
	atomic.AddInt32(&self.refCount, 1)
	return nil
}

func (self *RefCount) Close() error {
	if atomic.AddInt32(&self.refCount, -1) == 0 {
		return self.Closer.Close()
	}
	return nil
}

type Generator struct {
	Closer
	logger *log.Logger
	number uint64
	now    func() time.Time
}

func (self *Generator) Close() error {
	return self.Closer.Close()
}

func (self *Generator) NextNumber() (uint64, error) {
	if self.Closer.IsClosed() {
		return 0, self.Closer.CloseError()
	}
	return atomic.AddUint64(&self.number, 1), nil
}

func NewUniqueNumberGenerator(logger *log.Logger) (*Generator, error) {
	return &Generator{
		Closer: Closer{},
		logger: logger,
		number: 0,
		now: func() time.Time {
			return time.Now()
		},
	}, nil
}

func (self *Generator) Next() (string, error) {
	number, err := self.NextNumber()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v.%010d", self.now().Format("20060102150405"), number), nil
}
