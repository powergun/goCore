package functions

import "testing"


func TestModifyValue(t *testing.T) {
	f := func(pValue *int) {
		*pValue += 100
	}
	val := 0
	f(&val)
	if 100 != val {
		t.Error(val)
	}
}
