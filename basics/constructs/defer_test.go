package logic

import "testing"

func TestDeferAnonymousFunction(t *testing.T) {
	store := make(map[string]int)
	// initialize store
	func() {
		for _, s := range []string{"there", "is", "a", "cow"} {
			store[s] = len(s)
		}
	}()

	// teardown store
	defer func() {
		// see: https://stackoverflow.com/questions/13812121/how-to-clear-a-map-in-go
		// this cause the original data referenced by store to GC
		store = make(map[string]int)
	}()
}
