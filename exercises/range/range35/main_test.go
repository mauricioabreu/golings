// range35
// Make the tests pass!

// I AM NOT DONE
//
// scale multiplies every element of an array by k through a pointer.
// Practices a range over a pointer to an array.
package main_test

import "testing"

func scale(a *[3]int, k int) {
	for i := range a {
		a[i] += k
	}
}

func TestScale(t *testing.T) {
	a := [3]int{1, 2, 3}
	scale(&a, 10)
	if a != [3]int{10, 20, 30} {
		t.Errorf("scale = %v", a)
	}
}
