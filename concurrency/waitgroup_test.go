package concurrency

import (
	"sync"
	"testing"
)

type Counter struct {
	Num int
}

func Count(c *Counter, wg *sync.WaitGroup) {
	c.Num = 10
	wg.Done()
}

func TestExpectAllGoroutinesComplete(t *testing.T) {
	var wg sync.WaitGroup
	counter1 := Counter{}
	counter2 := Counter{}
	wg.Add(2)
	go Count(&counter1, &wg)
	go Count(&counter2, &wg)
	wg.Wait()
	t.Log(counter1.Num)
	t.Log(counter2.Num)
}
