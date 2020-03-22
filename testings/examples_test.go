package testings

import "testing"
import "github.com/stretchr/testify/assert"


func generateOne() int {
	return 1
}


func generateTwo() int {
	return 2
}


/*
NOTE: return type matters - if generateOne() returns uint8 the following assertions will fail

	Error Trace:	examples_test.go:18
	Error:      	Not equal:
	            	expected: int(1)
	            	actual: uint8(0x1)
	Error Trace:	examples_test.go:20
	Error:      	Not equal:
	            	expected: 0x2
	            	actual: 0x1
*/
func TestUsingAssert(t *testing.T) {
	assert.Equal(t, 1, generateOne())
	assert.NotEqual(t, 1, generateTwo())
	/*
	Error:      	Not equal:
				expected: 2
				actual: 1
	*/
	//assert.Equal(t, generateTwo(), generateOne())
}


