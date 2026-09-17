// arrays65
// Make the tests pass!

// I AM NOT DONE
//
// reverse reverses an array through a pointer.
// Practices swapping symmetric elements.
package main_test

import "testing"

func reverse(a *[5]int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i] = a[len(a)-i], a[i]
	}
}

func TestReverse(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	reverse(&a)
	if a != [5]int{5, 4, 3, 2, 1} {
		t.Errorf("reverse = %v", a)
	}
}
