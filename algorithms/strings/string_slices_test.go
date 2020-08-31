package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSlices(t *testing.T) {
	s := "thereisacoe"
	assert.Equal(t, "th", s[:2])
}
