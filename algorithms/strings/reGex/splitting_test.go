package reGex

import (
	"testing"
	"regexp"
)

func groups(s string, re *regexp.Regexp) []string {
	result := re.FindAllStringSubmatch(s, -1)
	if len(result) == 1 {
		return result[0][1: ]
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
