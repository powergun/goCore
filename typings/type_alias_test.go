package typings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Summable is an alias to int-slice
type Summable []int

// if the type (or type alias) itself is a slice or reference, the
// receiver should use value (use *pointer will not compile)
func (s Summable) doSum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func TestSummable(t *testing.T) {
	assert.Equal(t, 6, Summable{0, 1, 2, 3}.doSum())
}
