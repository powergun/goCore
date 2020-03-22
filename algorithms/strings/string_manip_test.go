package strings

import (
	"testing"
	"strings"
)

func TestIdentifyUppercase(t *testing.T) {
	if !IsUpper("S") {
		t.Error("not expected")
	}
	if IsUpper("s") {
		t.Error("not expected")
	}
}

func TestIdentifyNumericCharacter(t *testing.T) {
	if !IsDigit("1234") {
		t.Error("not expected")
	}
	if IsDigit("12ab34") {
		t.Error("not expected")
	}
}

// in Python: '*' * 8
func TestStringRepeater(t *testing.T) {
	if "********" != strings.Repeat("*", 8) {
		t.Fatal()
	}
}
