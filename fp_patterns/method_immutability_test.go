package fp_patterns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type T int

// double() takes a copy of T; mutating the original value is not
// possible
func (x T) double() T {
	return x * 2
}

func (x T) tryMutate() {
	x = x + 1
}

func TestMethodMutability(t *testing.T) {
	x := T(12)
	x.tryMutate()
	// NOTE T by default is no comparable to int
	// this won't work: Equal(t, 12, x)
	assert.Equal(t, T(12), x)
}
