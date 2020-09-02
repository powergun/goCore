package foobar

import "fmt"

// this function is called when this package is imported
var (
	ID = 0
)

func init() {
	fmt.Printf("called package init")
	ID = 101
}
