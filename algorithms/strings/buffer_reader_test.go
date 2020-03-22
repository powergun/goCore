package strings

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func AssertEqual(lhs interface{}, rhs interface{}) {
	if lhs != rhs {
		panic(fmt.Sprintf("lhs %s != rhs %s", lhs, rhs))
	}
}

func TestReadFromStaticBuffer(t *testing.T) {
	r := bufio.NewReader(bytes.NewBufferString("\n1\n2\n3\n"))
	drain := func(r *bufio.Reader) string {
		if s, err := r.ReadString('\n'); err == nil {
			return strings.TrimSuffix(s, "\n")
		} else {
			return ""
		}
	}
	AssertEqual(drain(r), "")
	AssertEqual(drain(r), "1")
	AssertEqual(drain(r), "2")
	AssertEqual(drain(r), "3")
	AssertEqual(drain(r), "")
}

const poisonPill = "==POISONPILL=="

func drainStdin(output chan string) {
	defer close(output)
	reader := bufio.NewReader(os.Stdin)
	for {
		if s, err := reader.ReadString('\n'); err != nil {
			output <- poisonPill
			return
		} else {
			output <- s
		}
	}
}

func drainSimulatedStdin(output chan string, s string) {
	defer close(output)
	reader := bufio.NewReader(bytes.NewBufferString(s))
	for {
		if s, err := reader.ReadString('\n'); err != nil {
			output <- poisonPill
			return
		} else {
			output <- s
		}
	}
}

func drink(ch chan string, t *testing.T) {
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				return
			} else {
				t.Logf("read from stdin: %s", s)
			}
		case <-time.After(time.Second * 1):
		}
	}
}

func TestReadFromSimulatedStdin(t *testing.T) {
	const text = `
doom 1992
e1m1 factory
imp01
shotgun
ammo crate

`
	ch := make(chan string)
	go drainSimulatedStdin(ch, text)

	drink(ch, t)
	t.Log("Done")
}
