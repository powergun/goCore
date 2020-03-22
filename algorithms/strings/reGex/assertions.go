package reGex

import "testing"

func assertTrue(t *testing.T, mustBeTrue bool) {
	if mustBeTrue == false {
		t.Fatalf("Not true")
	}
}

func assertFalse(t *testing.T, mustBeFalse bool) {
	if mustBeFalse == true {
		t.Fatalf("Not false")
	}
}

func assertStringEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatalf("Expected: %s; Got: %s", expected, actual)
	}
}

func assertStringsEqual(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Fatalf("Expected: %s; Got: %s", expected, actual)
	}
	for idx, lhs := range expected {
		rhs := actual[idx]
		if lhs != rhs {
			t.Fatalf("#%d members are different. Expected: %s; Got: %s", idx, lhs, rhs)
		}
	}
}

func assertEmpty(t *testing.T, v []string) {
	if len(v) != 0 {
		t.Fatalf("Not empty! %s", v)
	}
}
