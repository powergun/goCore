package typings

import (
	"fmt"
	"testing"
)

// empty interface can be useful to model an "unknown type" - a type
// to be determined at runtime
// however in production I should avoid this pattern, or at least
// cast it to a known type as early as possible

func whatIs(i interface{}) {
	// type to string (aka type-id)
	fmt.Printf("%T\n", i)
}

// type dispatch
// note do not exploit it as a sum type - there is no compile time
// checking
func typeSwitch(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("%s\n", i)
	case int:
		fmt.Printf("<%d>\n", i)
	default:
		fmt.Print("unknown\n")
	}
}

func TestWhatIs(t *testing.T) {
	whatIs([]int{1, 2, 3}) // showing `[]int`

	typeSwitch(1)
	typeSwitch("asd")
	typeSwitch(1.23)
}
