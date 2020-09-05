package sets

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSet(t *testing.T) {
	ss := make(map[string]struct{})
	for _, word := range []string{
		"there", "is", "a", "cow", "there",
	} {
		ss[word] = struct{}{}
	}
	arr := make([]string, len(ss))
	idx := 0
	for word := range ss {
		arr[idx] = word
		idx++
	}
	fmt.Printf("%v\n", arr)
	assert.Equal(t, 4, len(arr))
}
