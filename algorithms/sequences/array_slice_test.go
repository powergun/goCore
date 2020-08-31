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
