package strings

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestInt2String(t *testing.T) {
	testcases := map[int]string{
		10: "10",
		16: "a",
		//-1: "",  // cause panic
	}
	for base, expected := range testcases {
		x := strconv.FormatInt(10, base)
		assert.Equal(t, expected, x)
	}
}
