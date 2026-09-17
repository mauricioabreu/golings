// arrays3
// Make the tests pass!

// I AM NOT DONE
//
// sparse must return an array of length 5 where only the third element (index 2) is 10.
// Practices an array literal with explicit indexes.
package main_test

import "testing"

func sparse() [5]int {
	return [5]int{3: 10}
}

func TestSparse(t *testing.T) {
	if got := sparse(); got != [5]int{0, 0, 10, 0, 0} {
		t.Errorf("sparse = %v", got)
	}
}
