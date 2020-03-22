package functions

import (
	"testing"
	"fmt"
)


func PrtTestCase(args ...interface{}) {
	fmt.Printf("expected: %v, got: %v, args: %v\n", args[0], args[1], args[2:])
}


func PrtTc(tc interface{}) {
	fmt.Printf("%+v\n", tc)
}


func TestPrtTestCase(t *testing.T) {
	var testCases = []struct {
		value int
		ss string
		sub string
		expected bool
	} {
		{-1, "test", "testtest", false},
	}
	for _, tc := range testCases {
		result := true
		PrtTestCase(tc.expected, result, tc.value, tc.ss, tc.sub)
		PrtTc(tc)
	}
}
