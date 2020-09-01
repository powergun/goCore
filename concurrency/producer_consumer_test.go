package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func emit(c chan string) {
	words := []string{"i", "d", "d", "q", "d"}
	for _, w := range words {
		c <- w
	}
	close(c)
}

func TestProducerConsumer(t *testing.T) {
	wordChan := make(chan string)
	go emit(wordChan)
	words := []string{}
	for w := range wordChan {
		words = append(words, w)
	}
	assert.Equal(t, []string{"i", "d", "d", "q", "d"}, words)

	_, ok := <-wordChan
	assert.Equal(t, false, ok)

	for _, ok := <-wordChan; ok; {
		assert.Equal(t, false, true)
	}
}
