package random

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestRandomNumGen(t *testing.T) {
	x := rand.Intn(100)
	assert.Greater(t, x, -1)
}
