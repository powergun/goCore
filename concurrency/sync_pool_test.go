package concurrency

import (
	"math/rand"
	"sync"
	"testing"
)

// sync pool: thread safe construct that pools a bunch of resources
// go will not GC these resources

func makeBuffer() interface{} {
	return make([]byte, 100)
}

func TestSyncPool(t *testing.T) {
	pool := make([][]byte, 50)
	var spare sync.Pool
	spare.New = makeBuffer

	for i := 0; i < 10; i++ {
		go func(offset int) {
			for x := 3; x > 0; x-- {
				b := spare.Get().([]byte)
				j := offset + rand.Intn(20)
				if pool[j] != nil {
					spare.Put(pool[j])
				}
				pool[j] = b
			}
		}(i)
	}
}
