package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringLiterals(t *testing.T) {
	assert.Equal(t, `\n\t`, "\\n\\t")
}
