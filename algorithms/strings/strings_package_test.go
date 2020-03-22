package strings

import (
	"fmt"
	"testing"
	"unicode"
	"strings"
)


func TestFields(t *testing.T) {
	var testCases = []struct {
		variables []string
		expected string
	} {
		{[]string{"", ""}, "[ ]"},
		{[]string{"1", "2"}, "[1 2]"},
	}
	for _, testCase := range testCases {
		result := fmt.Sprintf("%v", testCase.variables)
		if testCase.expected != result {
			t.Error(fmt.Sprintf("expect %v; got %v", testCase.expected, result))
		}
	}
}


func TestFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	result := strings.FieldsFunc(" foo1; bar2; foobar3", f)
	if "foo1" != result[0] {
		t.Error()
	}
}


func TestStringCompare(t *testing.T) {
	var testCases = []struct {
		lhs string
		rhs string
		expected int
	} {
		{"a11", "bzzz", -1},  // lhs < rhs
		{"a11", "a11", 0},  // lhs == rhs
		{"b<<", "a11", 1},  // lhs > rhs
		{"", "0", -1},  // empty string < any string
	}
	for _, testCase := range testCases {
		if testCase.expected != strings.Compare(testCase.lhs, testCase.rhs) {
			t.Error(testCase)
		}
	}
}


func TestStringContains(t *testing.T) {
	var testCases = []struct {
		sub string
		ss string
		expected bool
	} {
		{"", "asd", true},  // any string "contains" empty string
		{"123", "asd", false},
		{"asdasd", "asd", false},
		}
	for _, tc := range testCases {
		result := strings.Contains(tc.ss, tc.sub)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringContainsAny(t *testing.T) {
	var testCases = []struct {
		sub string
		ss string
		expected bool
	} {
		{"u & i", "ii", true},  // whether ss contains at least one rune in sub
		{"    ", "ii", false},
		}
	for _, tc := range testCases {
		result := strings.ContainsAny(tc.ss, tc.sub)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringContainsRune(t *testing.T) {
	var testCases = []struct {
		chr rune
		ss string
		expected bool
	} {
		{'c', "asdcasd", true},
		{97, "asdasds", true},  // the ascii code for 'a' is 97
		{97, "nvhujre", false},
	}
	for _, tc := range testCases {
		result := strings.ContainsRune(tc.ss, tc.chr)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestCountSubStringInString(t *testing.T) {
	var testCases = []struct {
		sub      string
		ss       string
		expected int
	} {
		{"er", "thereisacow1337", 1},
	}
	for _, tc := range testCases {
		result := strings.Count(tc.ss, tc.sub)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringEqualFold(t *testing.T) {
	var testCases = []struct {
		lhs string
		rhs string
		expected bool
	} {
		{"thereisacow", "ThereIsACow", true},  // case folding
	}
	for _, tc := range testCases {
		result := strings.EqualFold(tc.lhs, tc.rhs)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringStartswith(t *testing.T) {
	var testCases = []struct {
		prefix string
		ss string
		expected bool
	} {
		{"ther", "Thereisacow", false},
		{"the", "there", true},
	}
	for _, tc := range testCases {
		result := strings.HasPrefix(tc.ss, tc.prefix)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestFindSubStringInString(t *testing.T) {
	var testCases = []struct {
		sub string
		ss string
		expected int
	} {
		{"", "", 0},
		{"asdasd", "", -1},
		{"ni", "selenium", 4},
	}
	for _, tc := range testCases {
		result := strings.Index(tc.ss, tc.sub)
		if tc.expected != result {
			t.Error(fmt.Sprintf("result: %v, test case: %+v", result, tc))
		}
	}
}


func TestFindSubStringInStringPredicate(t *testing.T) {
	var testCases = []struct {
		ss string
		expected int
	} {
		{">>. as24j. a2", 6},
	}
	predicate := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	for _, tc := range testCases {
		result := strings.IndexFunc(tc.ss, predicate)
		if tc.expected != result {
			t.Error(fmt.Sprintf("%v, %+v", result, tc))
		}
	}
}


func TestIndexAny(t *testing.T) {
	var testCases = []struct {
		sub string
		ss string
		expected int
	} {
		{"sa ow", "thereisacow", 6},  // ow is omitted
	}
	for _, tc := range testCases {
		result := strings.IndexAny(tc.ss, tc.sub)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestIndexByte(t *testing.T) {
	var testCases = []struct {
		chr byte
		ss string
		expected int
	} {
		{'d', "ret", -1},
	}
	for _, tc := range testCases {
		result := strings.IndexByte(tc.ss, tc.chr)  // see also IndexRune()
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestLastIndex(t *testing.T) {
	var testCases = []struct {
		ss string
		sub string
		expected int
	} {
		{"123 2421 12421 12", "12", 15},
		{"123 2421 12421 12", "13", -1},
	}
	for _, tc := range testCases {
		result := strings.LastIndex(tc.ss, tc.sub)  // see also LastIndexAny()
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringJoin(t *testing.T) {
	var testCases = []struct {
		tokens []string
		expected string
	} {
		{[]string{" ", "1", " "}, " 1 "},
	}
	for _, tc := range testCases {
		result := strings.Join(tc.tokens, "")
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringRepeat(t *testing.T) {
	var testCases = []struct {
		ss string
		count int
		expected string
	} {
		{"doom",3, "doomdoomdoom"},
	}
	for _, tc := range testCases {
		result := strings.Repeat(tc.ss, tc.count)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringReplace(t *testing.T) {
	var testCases = []struct {
		ss string
		subOld string
		subNew string
		numOps int
		expected string
	} {
		{"doom", "oo", "OO", -1, "dOOm"},
		{"doom", "o", "O", 1, "dOom"},
	}
	for _, tc := range testCases {
		// note: in Go 1.8, this function has been factored to two: Replace() and ReplaceN()
		result := strings.Replace(tc.ss, tc.subOld, tc.subNew, tc.numOps)
		if tc.expected != result {
			t.Error()
		}
	}
}


func TestStringSplit(t *testing.T) {
	var testCases = []struct {
		ss string
		sep string
		expected []string
	} {
		{"/doom/e1/m1", "/", []string{"", "doom", "e1", "m1"}},
	}
	for _, tc := range testCases {
		result := strings.Split(tc.ss, tc.sep)
		if !(4 == len(result) && "" == result[0]) {
			t.Error()
		}
	}
}


func TestTitle(t *testing.T) {
	var testCases = []struct {
		orig string
		expected string
	} {
		{"doom", "Doom"},
		{"...asd", "...Asd"},  // this is a bug
		{"111asd", "111asd"},
	}
	for _, tc := range testCases {
		result := strings.Title(tc.orig)
		if tc.expected != result {
			t.Error(fmt.Sprintf("result: %v, %+v", result, tc))
		}
	}
}


func TestTrim(t *testing.T) {
	var testCases = []struct {
		ss string
		cutset string
		expected string
	} {
		{"* /doom/e1/m1/map: *", "* :","/doom/e1/m1/map"},
	}
	for _, tc := range testCases {
		result := strings.Trim(tc.ss, tc.cutset)
		if tc.expected != result {
			t.Error(fmt.Sprintf("result: %v, %+v", result, tc))
		}
	}
}


