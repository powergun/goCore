package concurrency

import "testing"

func TestExplicitlyCloseChannel(t *testing.T) {
	channel := make(chan int,  1)
	channel <- 3
	t.Log(<- channel)
	close(channel)

	// because channel is closed, the value coming out of it is zero-value
	t.Log(<- channel)
	t.Log(<- channel)
}

