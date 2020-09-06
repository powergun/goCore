package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// this example demonstrates that unbuffered channel is easy to reason
// about

func TestReasoning(t *testing.T) {
	c := make(chan int)

	go func() {
		i := 0
		for {
			// this infinite loop "is going nowhere" (it pauses at here)
			// if no one reads from this channel
			c <- i
			i += 1
		}
	}()

	// reading data will "drive" the goroutine to further computing
	// the next value
	assert.Equal(t, 0, <-c)
	assert.Equal(t, 1, <-c)
	assert.Equal(t, 2, <-c)

	// the channel and goroutine are tore down automatically
}
