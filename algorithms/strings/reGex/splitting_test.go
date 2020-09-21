package reGex

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSplit(t *testing.T) {
	reg := regexp.MustCompile(`[\s\n\t,.]+`)
	tokens := reg.Split(`there  . is ,   

a

cow`, -1)
	assert.Equal(t, []string{"there", "is", "a", "cow"}, tokens)
}

func TestSplitHead(t *testing.T) {
	reg := regexp.MustCompile(`[\s\n\t,.]+`)
	tokens := reg.Split(`there  . is ,   

a

cow`, 2)
	// n: at most n substrings
	assert.Equal(t, "there", tokens[0])
}

func groups(s string, re *regexp.Regexp) []string {
	result := re.FindAllStringSubmatch(s, -1)
	if len(result) == 1 {
		return result[0][1:]
	}
	return nil
}

func TestSplitDeterministic(t *testing.T) {
	re := regexp.MustCompile(`^(.+):(\d+)$`)
	assertEmpty(t, groups("", re))
	assertEmpty(t, groups("....", re))
	assertEmpty(t, groups("::::", re))
	assertEmpty(t, groups(":0", re))
	assertStringsEqual(t, []string{" ", "0"}, groups(" :0", re))
	assertStringsEqual(t, []string{"asd", "10"}, groups("asd:10", re))
	assertStringsEqual(t, []string{"/doom/e1m1/imp.id1", "123"}, groups("/doom/e1m1/imp.id1:123", re))
}
