package sequences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraySlice(t *testing.T) {
	words := []string{"i", "d"}
	assert.Equal(t, 2, length(words))
	assert.Equal(t, "iddqd", words[0])
}

// nasty mutator function... do not do that in real world
func length(xs []string) int {
	xs[0] = "iddqd"
	return len(xs)
}

func TestCreateSlice(t *testing.T) {
	// allocate mem for a 4-elems array and create a slice for it, xs
	xs := make([]string, 4)
	xs[0] = "id"
	// access to uninitialized element 1, 2, 3 - panic!
	for i, w := range xs {
		if i > 0 {
			break
		}
		assert.Equal(t, "i", w[0:1])
	}

	ys := xs[:]
	// both xs and ys hold references to the underlying array elements

	// ns is a slice pointing a newly allocated array;
	// copy() copies element-wise from xs to ns
	ns := make([]string, 4)
	copy(ns, xs)
	assert.Equal(t, xs, ns)
	assert.Equal(t, xs, ys)
}
