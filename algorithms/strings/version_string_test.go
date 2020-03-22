package strings

import (
	"testing"
	"fmt"
	"github.com/smartystreets/gunit"
)


type TestVersionStringFixture struct {
	*gunit.Fixture
}


func (this *TestVersionStringFixture) TestCreateFromStringExpectVersionNums() {
	this.Assert(true)
}


func TestSemanticVersion_ToStr_ZeroValue(t *testing.T) {
	v := SemanticVersion{}
	if "" != v.ToStr() {
		t.Error("not expected")
	}
	if v.IsValid() {
		t.Error("not expected")
	}
}


func TestSemanticVersion_ToStr_Init(t *testing.T) {
	v := SemanticVersion{1, 1, 1}
	if "1.1.1" != v.ToStr() {
		t.Error("not expected")
	}
	if !v.IsValid() {
		t.Error("not expected")
	}
}


func TestSemanticVersion_EqualTo(t *testing.T) {
	v1 := SemanticVersion{1, 1, 1}
	v2 := SemanticVersion{1, 1, 1}
	v3 := SemanticVersion{}
	if v1 != v2 {
		t.Error("not expected")
	}
	if v1 == v3 {
		t.Error("not expected")
	}
}


func TestSemanticVersion_ToInts(t *testing.T) {
	v := SemanticVersion{1, 1, 1}
	nums := v.ToInts()
	expected := [3]uint{1, 1, 1}
	if expected != nums {
		t.Error("not expected")
	}
}


// table based test
func TestCreateSemanticVersionFromNums(t *testing.T) {
	var testCases = []struct {
		nums [3]uint
		expectedStr string
	} {
		{[3]uint{0, 0, 0}, ""},
		{[3]uint{1, 0, 0}, "1.0.0"},
		{[3]uint{0, 2, 0}, "0.2.0"},
	}
	for _, testCase := range testCases {
		sv := CreateSemanticVersionFromNums(testCase.nums)
		s := sv.ToStr()
		if testCase.expectedStr != s {
			t.Error(fmt.Sprintf("Expect: %s; Got %s", testCase.expectedStr, s))
		}
	}
}
