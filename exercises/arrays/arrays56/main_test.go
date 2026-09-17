// arrays56
// Make the tests pass!

// I AM NOT DONE
//
// head returns a slice of the first two elements of an array. The caller appends
// to that slice and overwrites the third element of the array.
// The full slice expression a[low:high:max] caps the capacity.
package main_test

import "testing"

func head(a *[3]int) []int {
	return a[0:2]
}

func TestHead(t *testing.T) {
	a := [3]int{1, 2, 3}
	h := head(&a)
	h = append(h, 99)
	if a[2] != 3 {
		t.Errorf("array overwritten: %v", a)
	}
	if len(h) != 3 || h[2] != 99 {
		t.Errorf("h = %v", h)
	}
}
