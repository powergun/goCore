package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// a context can be shared among various producers, with each producer
// depending on the "Done()" mechanism to stop computing;
// these producers can form a data pipeline

type producerContext struct {
	done chan struct{}
}

func newProducerContext() *producerContext {
	ctx := new(producerContext)
	ctx.done = make(chan struct{})
	return ctx
}

// have a single place to  close the channel to ensure it won't be
// closed twice
func (ctx *producerContext) stop() {
	close(ctx.done)
}

func (ctx *producerContext) getDone() <-chan struct{} {
	return ctx.done
}

type betterProducer struct {
	ctx *producerContext
	c   chan int
	i   int
}

func newBetterProducer(ctx *producerContext, wg *sync.WaitGroup) *betterProducer {
	prod := new(betterProducer)
	prod.c = make(chan int)
	prod.ctx = ctx

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := prod.ctx.getDone()
		for {
			select {
			case prod.c <- prod.i:
				prod.i++
			case <-done:
				return
			}
		}
	}()
	return prod
}

func (prod *betterProducer) getSource() <-chan int {
	return prod.c
}

func TestContext(t *testing.T) {
	ctx := newProducerContext()
	wg := sync.WaitGroup{}
	prod := newBetterProducer(ctx, &wg)
	read := prod.getSource()

	assert.Equal(t, 0, <-read)
	assert.Equal(t, 1, <-read)
	assert.Equal(t, 2, <-read)
	assert.Equal(t, 3, <-read)
	assert.Equal(t, 4, <-read)
	assert.Equal(t, 5, <-read)

	ctx.stop()
	wg.Wait()

	fmt.Println("terminate cleanly")
}
