package UniqueNumber

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Generator struct {
	number uint64
	now    func() time.Time
}

func (self *Generator) NextNumber() uint64 {
	return atomic.AddUint64(&self.number, 1)
}

func NewUniqueNumberGenerator() *Generator {
	return &Generator{
		number: 0,
		now: func() time.Time {
			return time.Now()
		},
	}
}

func (self *Generator) Next() string {
	number := self.NextNumber()
	return fmt.Sprintf("%v.%010d", self.now().Format("20060102150405"), number)
}
