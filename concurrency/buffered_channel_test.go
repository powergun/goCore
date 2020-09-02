package concurrency

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

var (
	running int64 = 0
)

// one use of the buffered channel is to control concurrency, like
// a semaphore;
// unbuffered channels are easier to reason about

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d", running)
	d := time.Duration(rand.Intn(10))
	time.Sleep(d * time.Millisecond)
	atomic.AddInt64(&running, -1)
	fmt.Printf("]")
}

func do_work(sema chan bool) {
	// [[[[[[[[][[][[][][][][][...
	// [2[3[9[8[6[5[10[7][10[1][10[4]][8[9][10....
	<-sema
	work()
	sema <- true
}

func TestBufferedChannel(t *testing.T) {
	// once the channel reaches its capacity (100), other writers have
	// to wait
	sema := make(chan bool, 10) // up to 10 goroutines
	const num_workers = 1000
	const sz = 10
	for i := num_workers; i > 0; i-- {
		go do_work(sema)
	}
	for i := sz; i > 0; i-- {
		sema <- true
	}
	time.Sleep(10 * time.Millisecond)
}
