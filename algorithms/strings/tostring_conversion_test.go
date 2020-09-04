package strings

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestString2Int(t *testing.T) {
	tests := map[string]int{
		"2vv": 0, // fail
		"2":   2,
		".2":  0,
		"1.2": 1,
		"1.9": 1, // 1.2 or 1.9 is truncated to 1
	}
	for s, n := range tests {
		if x, err := strconv.Atoi(s); err == nil {
			assert.Equal(t, n, x)
		}
	}
}
