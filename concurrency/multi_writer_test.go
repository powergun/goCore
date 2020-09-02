package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

// this is similar to the real world app I developed at Wt, call
// flamebird, which wrangled some data over the network, parsed it
// and sent it to elasticsearch

// I can also introduce a simple load-balancing mechanism here:
// instead of spinning up a goroutine for each payload, I can keep a
// pool of workers (serving indefinitely), with each taking inputs
// from a shared channel

func getPage(url string) (int, error) {
	if resp, err := http.Get(url); err != nil {
		return -1, err
	} else {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return -1, err
		} else {
			return len(body), nil
		}
	}
}

func getPage2(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	return len(body), nil
}

func getPage3(url string) (int, error) {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			return len(body), nil
		} else {
			return -1, err
		}
	} else {
		return -1, err
	}
}

func TestGetPage(t *testing.T) {
	_, err := getPage3("http://www.google.com")
	assert.Equal(t, nil, err)

	_, err = getPage3("iddqd")
	assert.NotEqual(t, nil, err)
}

func getter(url string, sz chan int) {
	length, err := getPage3(url)
	if err == nil {
		sz <- length
	}
}

func TestMultiWriter(t *testing.T) {
	ch := make(chan int)
	urls := []string{
		"http://www.google.com", "http://www.linkedin.com",
		"http://www.amazon.com", "http://www.apple.com"}
	for _, url := range urls {
		go getter(url, ch)
	}
	for i := len(urls); i > 0; i-- {
		fmt.Printf("%d ", <-ch)
	}
}
