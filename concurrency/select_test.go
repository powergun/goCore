package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func generate(wordChan chan string, done chan bool) {
	words := []string{"i", "d", "d", "q", "d"}
	for i := 0; ; {
		select {
		case wordChan <- words[i]:
			i = (i + 1) % len(words)
		case <-done:
			done <- false
			return
		}
	}
}

func TestSelect(t *testing.T) {
	wordChan := make(chan string)
	doneChan := make(chan bool)
	go generate(wordChan, doneChan)
	for i := 0; i < 10; i++ {
		fmt.Printf("%s", <-wordChan)
	}
	// channel is bidirectional
	doneChan <- true

	assert.False(t, <-doneChan)
}
