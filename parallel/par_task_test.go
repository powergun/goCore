package parallel

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// NOTE:
// - find out a better parallel map library (if it exists)
// - or extend this mini framework to support parallel map and reduce
// - implement the source code line counter program
// - compare its perf with the c++ impl (tbb) and python's

type task interface {
	process()
	output()
}

type factory interface {
	create(line string) task
}

func run(f factory, src io.Reader) {
	var wg sync.WaitGroup

	in := make(chan task)
	wg.Add(1)

	go func() {
		s := bufio.NewScanner(src)
		for s.Scan() {
			in <- f.create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	// source (number of cpus)
	// https://stackoverflow.com/questions/24073697/how-to-find-out-the-number-of-cpus-in-go-lang
	for i := runtime.NumCPU(); i > 0; i-- {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.output()
	}
}

type http_task_t struct {
	url string
	ok  bool
}

func (h *http_task_t) process() {
	if resp, err := http.Get(h.url); err != nil {
		h.ok = false
	} else {
		if resp.StatusCode != http.StatusOK {
			h.ok = false
		}
	}
	h.ok = true
}

func (h *http_task_t) output() {
	fmt.Printf("%s %t\n", h.url, h.ok)
}

type http_task_factory struct{}

func (f http_task_factory) create(line string) task {
	t := &http_task_t{}
	t.url = line
	return t
}

func TestHttpTaskStatusOnly(t *testing.T) {
	src := strings.NewReader(`
http://www.google.com
http://www.baidu.com
http://www.qq.com
http://www.amazon.com
`)
	fac := http_task_factory{}
	run(fac, src)
}
