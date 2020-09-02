package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// one use of the buffered channel is to control concurrency, like
// a semaphore

func work() {
	fmt.Printf("[")
	d := time.Duration(rand.Intn(10))
	time.Sleep(d * time.Millisecond)
	fmt.Printf("]")
}

func do_work(sema chan bool) {
	// [[[[[[[[][[][[][][][][][...
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
	time.Sleep(1 * time.Second)
}
