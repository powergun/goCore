package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// one can exploits close(chan) to coordinate works between different
// goroutines

// gate
func worker(msg string, goCh chan bool, output chan string) {
	// wait for the "go ahead" signal
	// if the scheduler only sends a bool value to this channel, only
	// one goroutine will receive it; however if the schedule closes
	// this channel, all the goroutines will "receive the signal"
	<-goCh

	output <- fmt.Sprintf("%s ", msg)
}

func TestGate(t *testing.T) {
	goCh := make(chan bool)
	outputCh := make(chan string)
	const sz = 10
	for i := 0; i < sz; i++ {
		go worker(fmt.Sprintf("..."), goCh, outputCh)
	}
	time.Sleep(2 * time.Second)
	close(goCh)
	for i := 10; i > 0; i-- {
		fmt.Print(<-outputCh)
	}
}

// stopper
func server(msg string, stopCh chan bool) {
	for { // serve indefinitely
		select {
		// will fall into this case if the scheduler close the channel
		case <-stopCh:
			return
		default:
			fmt.Printf("...\n")
		}
	}
}

func TestStopper(t *testing.T) {
	stopCh := make(chan bool)
	const sz = 3
	for i := sz; i > 0; i-- {
		go server("...", stopCh)
	}
	time.Sleep(1 * time.Millisecond)
	close(stopCh)
	time.Sleep(1 * time.Millisecond)
}
