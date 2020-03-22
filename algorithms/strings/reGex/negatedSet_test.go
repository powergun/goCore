package reGex

import (
	"testing"
	"regexp"
)

func TestNegatedSetContainsMultipleCharacters(t *testing.T) {
	re := regexp.MustCompile(`[^ (I]+`)
	assertStringEqual(t, "", re.FindString(""))
	assertStringEqual(t, "something", re.FindString("something"))
	assertStringEqual(t, "something", re.FindString("something "))
	assertStringEqual(t, "something", re.FindString("something( "))
	assertStringEqual(t, "something", re.FindString("somethingI( "))

	re = regexp.MustCompile("^[^ (I]+$")
	assertFalse(t, re.MatchString("somethingI( "))
}

func TestOptionalNegatedSet(t *testing.T) {
	re := regexp.MustCompile(`^\?\?$|^[^ (I]+$`)
	t.Log(re.FindString("SomeModule::SomeClass::someMeth(int,"))
	assertFalse(t, re.MatchString("SomeModule::SomeClass::someMeth(int,"))
}


