package functions

import "testing"


func TestReturnErrorObject(t *testing.T) {
	largerThan10 := func (v int) bool {
		return v > 10
	}
	err := MakeErrorIf(largerThan10(3))
	if err != nil {
		t.Error(err)
	}
	err = MakeErrorIf(largerThan10(11))
	if err == nil {
		t.Error(err)
	}
}


func TestReturnCustomErrorObject(t *testing.T) {
	err := MakeErrorIfLarger(10, 2)
	if err == nil {
		t.Error(err)
	}
	//a more violent version:
	//panic(err)
}

