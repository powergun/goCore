package sequences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	// []string{"..", ...} defines a slice of strings
	words := [3]string{"iddqd", "idkfa", "idnoclip"}
	for _, w := range words {
		// w[0] results in uint8
		assert.Equal(t, uint8('i'), w[0])
		assert.Equal(t, "i", w[0:1])
	}
}
