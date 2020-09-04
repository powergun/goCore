package ordered

import (
	"fmt"
	"sort"
	"testing"
)

// implement the sort interface for int-slice
type IS []int

func (is IS) Len() int {
	return len(is) // is.Len() will cause infinite recursion
}

func (is IS) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IS) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func TestBuiltinSort(t *testing.T) {
	xs := IS{
		1, 2, 3, -1, -2, -3,
	}
	// in normal circumstance I should just use sort.Ints()
	sort.Ints(xs)
	fmt.Printf("%v\n", xs)
	sort.Sort(xs)
	fmt.Printf("%v\n", xs)
}
