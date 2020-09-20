package reGex

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetValueInsideCaptureGroups(t *testing.T) {
	r := regexp.MustCompile(`p=(\w+)`)
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

func TestMultipleCaptureGroups(t *testing.T) {
	r := regexp.MustCompile(`(\S+)\s+is\s+(\S+)`)
	text := ` asd is  asd `
	result := r.FindAllStringSubmatch(text, -1)
	for idx, tokens := range result {
		// tokens slice consists of 3 elements:
		// the whole match (asd is  asd), the first capture group, the second capture group
		fmt.Printf("%v |%v|%v|%v|\n", idx, tokens[0], tokens[1], tokens[2])
	}
}
