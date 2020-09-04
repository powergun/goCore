package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

// synchronize memory accesses among goroutines
// the sync package provides the synchronization primitives
// there is also a RWMutex

// build synchronisation into the data structure by incorporate
// the Mutex struct in the custom struct
type ProtectedCache struct {
	sync.Mutex
	c map[string]string
}

// note that now I need an explicit init function to initialize
// the map since it is not a map-slice
// also note that this custom cache can not be copied around, because
// there should only be one mutex
func (pc *ProtectedCache) init() {
	pc.c = make(map[string]string)
}

func TestSyncAccess(t *testing.T) {
	ns := []string{
		"e1m1", "e2m2", "there", "is", "a", "cow", "13",
		"there", "is", "a", "silence",
		"where",
	}
	cache := make(map[string]int)

	producer := func(n string) int {
		return len(n)
	}
	mutex := sync.Mutex{}
	for _, name := range ns {
		go func(n string) {
			mutex.Lock()
			defer mutex.Unlock()
			cache[n] = producer(n)
		}(name)
	}
}

func TestProtectedCache(t *testing.T) {
	ns := []string{
		"e1m1", "e2m2", "there", "is", "a", "cow", "13",
		"there", "is", "a", "silence",
		"where",
	}
	cache := ProtectedCache{}
	cache.init()

	producer := func(n string) string {
		return fmt.Sprintf("%d", len(n))
	}
	for _, name := range ns {
		go func(n string) {
			cache.Lock()
			defer cache.Unlock()
			cache.c[n] = producer(n)
		}(name)
	}
}
