package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoopThroughCharacters(t *testing.T) {
	for i, r := range "iddqd" {
		assert.Greater(t, i, -1)
		assert.Greater(t, r, 'a')
	}
}
