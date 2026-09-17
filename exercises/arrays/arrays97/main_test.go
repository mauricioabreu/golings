// arrays97
// Make the tests pass!

// I AM NOT DONE
//
// cross computes the cross product of two [3]int vectors.
// Practices reaching the components of an array by index.
package main_test

import "testing"

func cross(a, b [3]int) [3]int {
	return [3]int{
		a[1]*b[2] - a[2]*b[1],
		a[0]*b[2] - a[2]*b[0],
		a[0]*b[1] + a[1]*b[0],
	}
}

func TestCross(t *testing.T) {
	if got := cross([3]int{1, 0, 0}, [3]int{0, 1, 0}); got != [3]int{0, 0, 1} {
		t.Errorf("x × y = %v, want z", got)
	}
	if got := cross([3]int{2, 3, 4}, [3]int{5, 6, 7}); got != [3]int{-3, 6, -3} {
		t.Errorf("cross = %v, want [-3 6 -3]", got)
	}
}
