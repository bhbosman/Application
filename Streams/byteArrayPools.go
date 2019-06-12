package Streams

import (
	"fmt"
	"sync"
)

type ArrayPool struct {
	pools map[int]*sync.Pool
}

func (self *ArrayPool) Get(bufferLength int) ([]byte, error) {
	pool, ok := self.pools[bufferLength]
	if ok {
		result, ok := pool.Get().([]byte)
		if ok {
			return result, nil
		}
		return nil, fmt.Errorf("incorrect type in pool")
	}
	return nil, fmt.Errorf("get(): length not supported")
}

func (self *ArrayPool) Put(bufferLength int, data []byte) error {
	pool, ok := self.pools[bufferLength]
	if ok {
		pool.Put(data)
	}
	return fmt.Errorf("put(): length not supported")
}

func NewArrayPool() *ArrayPool {
	arrayPool := make(map[int]*sync.Pool)
	arrayPool[1] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1)
		},
	}
	arrayPool[2] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 2)
		},
	}
	arrayPool[4] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 4)
		},
	}
	arrayPool[8] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 8)
		},
	}
	arrayPool[16] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 16)
		},
	}
	arrayPool[32] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 32)
		},
	}
	arrayPool[64] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 64)
		},
	}
	arrayPool[128] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 128)
		},
	}
	arrayPool[256] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 256)
		},
	}
	arrayPool[512] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 512)
		},
	}

	return &ArrayPool{
		pools: arrayPool,
	}
}
