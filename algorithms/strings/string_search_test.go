package strings

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// sub problems:
//
// needle and haystack
// word count

func TestWordCount(t *testing.T) {
	s := "there is a cow, there is a silence"
	subs := "is"
	assert.Equal(t, 2, strings.Count(s, subs))

	// if the substring is an empty string, the count is len(s) + 1
	assert.Equal(t, 6, strings.Count("iddqd", ""))
}
