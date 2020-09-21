package reGex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

// nested capture group, source:
// https://regexone.com/lesson/nested_groups
func TestNestedCaptureGroup(t *testing.T) {
	reg := regexp.MustCompile(`(a(b(.+)))`)
	for _, match := range reg.FindAllStringSubmatch(`abcow`, -1) {
		// the whole match
		assert.Equal(t, "abcow", match[0])

		// 1st capture group (the outermost one)
		assert.Equal(t, "abcow", match[1])
		// 2nd capture group (one layer down)
		assert.Equal(t, "bcow", match[2])
		// 3nd capture group, innermost
		assert.Equal(t, "cow", match[3])
	}
}

func TestGetValueInsideCaptureGroups(t *testing.T) {
	r := regexp.MustCompile(`(\w+)=([^,=]+)`)
	text := `
p=n1, p=n2, 

p=n3,
,
,,
`
	// result looks like this
	//          __ 2nd group
	// [[p=n1 p n1] [p=n2 p n2] [p=n3 p n3]]
	//        ^ 1st group
	//   ^^^^ match string
	// each element consists of the whole match and the capture group
	result := r.FindAllStringSubmatch(text, -1)
	matched := make([]string, len(result))
	for idx, tokens := range result {
		matched[idx] = tokens[2]
	}
	assert.Equal(t, []string{"n1", "n2", "n3"}, matched)
}

func TestMultipleCaptureGroups(t *testing.T) {
	r := regexp.MustCompile(`(\S+)\s+is\s+(\S+)`)
	text := ` asd is  asd `
	result := r.FindAllStringSubmatch(text, -1)
	for idx, tokens := range result {
		// tokens slice consists of 3 elements:
		// - [0] the whole match (asd is  asd)
		// - [1] the first capture group
		// - [2] the second capture group
		fmt.Printf("%v |%v|%v|%v|\n", idx, tokens[0], tokens[1], tokens[2])
	}
}
