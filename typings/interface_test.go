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

// interface composition
type weightedShuffler interface {
	shuffler // compose anything that is part of the shuffler interface

	Weight(i int) int
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

func weightedShuffle(ws weightedShuffler) {
	sz := ws.Len()
	totalWeight := 0
	for i := 0; i < sz; i++ {
		totalWeight += ws.Weight(i)
	}
	for i := 0; i < sz; i++ {
		pos := rand.Intn(totalWeight)
		cum := 0
		for j := i; j < sz; j++ {
			if pos >= cum {
				totalWeight -= ws.Weight(j)
				ws.Swap(i, j)
				break
			}
		}
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

func (is StrSlice) Weight(i int) int {
	return 100 - len(is[i])
}

func TestShuffle(t *testing.T) {
	is := ISlice{1, 2, 3, 4, 5, 6}
	shuffle(is)
	fmt.Printf("%v\n", is)

	ss := StrSlice{"i", "d", "d", "q", "d"}
	shuffle(ss)
	fmt.Printf("%v\n", ss)
}

func TestWeightedShuffle(t *testing.T) {
	ss := StrSlice{
		"thereisasilence",
		"idnoclip", "thereisa",
		"id", "iddqd",
	}
	// is compatible with shuffle() function
	//shuffle(ss)
	//fmt.Printf("%v\n", ss)

	weightedShuffle(ss)
	fmt.Printf("%v\n", ss)
}
