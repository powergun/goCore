package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// nil channel can be useful because it does not block;
// can disable part of the select structure

func reader(ch chan int) {
	t := time.NewTimer(30 * time.Millisecond)
	for {
		select {
		case i := <-ch: // print out if receives something, otherwise
			// carry on
			fmt.Printf("%d\n", i)
		case <-t.C:
			ch = nil // disable `case i := <-ch`
		}
	}
}

func writer(ch chan int) {
	for {
		ch <- rand.Intn(1000)
	}
}

func TestNilChannel(t *testing.T) {
	ch := make(chan int)
	// consumer
	go reader(ch)
	// producer
	go writer(ch)

	time.Sleep(50 * time.Millisecond)
}
