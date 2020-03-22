package reGex

import (
	"testing"
	"regexp"
	"fmt"
)

func TestMatchWholeWord(t *testing.T) {
	r := regexp.MustCompile("doom.+")
	text := "/doom/e1/m1/imp01"
	result := r.FindString(text)
	if "doom/e1/m1/imp01" != result {
		t.Errorf("failed: %s", result)
	}
}

func TestGetValueInsideCaptureGroups(t *testing.T) {
	r := regexp.MustCompile("p=(\\w+)")
	text := "p=n1, p=n2, p=n3,,,,"
	result := r.FindAllStringSubmatch(text, -1)
	matched := make([]string, len(result))
	for idx, tokens := range result {
		matched[idx] = tokens[1]
	}
	if "[n1 n2 n3]" != fmt.Sprint(matched) {
		t.Errorf("failed: %s", matched)
	}
}
