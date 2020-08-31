package functions

import (
	"fmt"
	"testing"
)

func PrtTestCase(args ...interface{}) {
	fmt.Printf("expected: %v, got: %v, args: %v\n", args[0], args[1], args[2:])
}

// iterate over each element in the variadic parameters
func PrintEach(args ...interface{}) {
	for idx, arg := range args {
		fmt.Printf("%d %q, ", idx, arg)
	}
	fmt.Printf("\n")
}

func PrtTc(tc interface{}) {
	fmt.Printf("%+v\n", tc)
}

func TestPrtTestCase(t *testing.T) {
	var testCases = []struct {
		value    int
		ss       string
		sub      string
		expected bool
	}{
		{-1, "test", "testtest", false},
	}
	for _, tc := range testCases {
		PrtTestCase(tc.expected, true, tc.value, tc.ss, tc.sub)
		PrtTc(tc)
		PrintEach(tc.expected, true, tc.value, tc.ss, tc.sub)
	}
}
