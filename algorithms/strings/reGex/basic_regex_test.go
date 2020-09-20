package reGex

import (
	"regexp"
	"testing"
)

func TestMatchWholeWord(t *testing.T) {
	r := regexp.MustCompile("doom.+")
	text := "/doom/e1/m1/imp01"
	result := r.FindString(text)
	if "doom/e1/m1/imp01" != result {
		t.Errorf("failed: %s", result)
	}
}
