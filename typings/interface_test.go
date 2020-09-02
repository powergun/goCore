package typings

import (
	"fmt"
	"math/rand"
	"testing"
)

type shuffler interface {
	// anything can be a "shuffler" if it implements the following
	// methods:
	Len() int
	Swap(i, j int)
}

// compatibility checking is done by matching the method(s)
// (can not be a pointer)
func shuffle(s shuffler) {
	sz := s.Len()
	for i := 0; i < sz; i++ {
		j := rand.Intn(sz - i)
		s.Swap(i, j)
	}
}

// create a concrete type
type ISlice []int

func (is ISlice) Len() int {
	return len(is)
}

func (is ISlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// create another concrete type that also satisfies the Shuffler
// interface
type StrSlice []string

func (ss StrSlice) Len() int {
	return len(ss)
}

func (ss StrSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func TestShuffle(t *testing.T) {
	is := ISlice{1, 2, 3, 4, 5, 6}
	shuffle(is)
	fmt.Printf("%v\n", is)

	ss := StrSlice{"i", "d", "d", "q", "d"}
	shuffle(ss)
	fmt.Printf("%v\n", ss)
}
