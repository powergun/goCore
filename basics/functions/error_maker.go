package functions

import (
	"errors"
	"fmt"
)


func MakeErrorIf(flag bool) error {
	if !flag {
		return nil
	}
	return errors.New("made error")
}


func MakeErrorIfLarger(v int, threshold int) error {
	if v <= threshold {
		return nil
	}
	return &CustomNumericError{v, "larger"}
}


type CustomNumericError struct {
	value int
	msg string
}


func (e *CustomNumericError) Error() string {
	return fmt.Sprintf("value: %d, reason: %s", e.value, e.msg)
}

