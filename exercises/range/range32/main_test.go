// range32
// Make the tests pass!

// I AM NOT DONE
//
// indexAfter returns the position of the first negative number, or len(s).
// The code does not compile: the loop variable is out of reach afterwards.
// Variables declared in a for header are only visible inside it.
package main_test

import "testing"

func indexAfter(s []int) int {
	for i, v := range s {
		if v < 0 {
			break
		}
	}
	return i
}

func TestIndexAfter(t *testing.T) {
	if indexAfter([]int{1, 2, -3, 4}) != 2 || indexAfter([]int{1, 2}) != 2 || indexAfter(nil) != 0 {
		t.Errorf("indexAfter works incorrectly")
	}
}
