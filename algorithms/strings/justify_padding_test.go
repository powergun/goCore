package strings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// left/right justify and padding

// padding by printf():
// Pros: not difficult to remember
// Cons: hardcode the total length of the resulting string
func TestPadding(t *testing.T) {
	s := "iddqd"

	// padding on the right side
	assert.Equal(t, "|iddqd     |", fmt.Sprintf("|%-10s|", s))

	// padding on the left side
	assert.Equal(t, "|     iddqd|", fmt.Sprintf("|%10s|", s))
}

// define my own justify function
// Pros: support dynamic length
// Cons: boilerplate code
// source
// https://www.socketloop.com/tutorials/golang-aligning-strings-to-right-left-and-center-with-fill-example
func TestLeftJustify(t *testing.T) {
	s := "thereisacow"

	f := func(x string, n int, c string) string {
		if sz := n - len(s); sz <= 0 {
			return x
		} else {
			return strings.Repeat(c, sz) + x
		}
	}

	// with sufficient length
	assert.Equal(t, ".........thereisacow", f(s, 20, "."))

	// with insufficient length
	assert.Equal(t, "thereisacow", f(s, 9, "x"))
}
