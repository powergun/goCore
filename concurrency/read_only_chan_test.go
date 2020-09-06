package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// google develops the concept of Context based on the read-only channel
// see context_test.go

type producer struct {
	c    chan int
	done chan struct{}
	i    int
}

// ctor
func newProducer() *producer {
	prod := new(producer)
	prod.c = make(chan int)
	prod.done = make(chan struct{})

	// spawn the goroutine that indefinitely produces values
	go func() {
		for {
			select {
			case <-prod.done:
				return
			case prod.c <- prod.i:
				prod.i++
			}
		}
	}()
	return prod
}

// return a read-only channel!!
// read := prod.getSource()
// read <- 123 won't compile
func (prod *producer) getSource() <-chan int {
	return prod.c
}

func (prod *producer) stop() {
	prod.done <- struct{}{}
}

func TestReadOnlyChan(t *testing.T) {
	prod := newProducer()
	read := prod.getSource()

	// won't compile!
	//read <- 213

	assert.Equal(t, 0, <-read)
	assert.Equal(t, 1, <-read)
	assert.Equal(t, 2, <-read)

	prod.stop()
	// <-read will deadlock
}
