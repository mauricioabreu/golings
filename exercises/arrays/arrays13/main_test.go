// arrays13
// Make the tests pass!

// I AM NOT DONE
//
// find looks a value up in a sorted array and returns its index, or -1.
// Practices the lo/hi bounds of a binary search.
package main_test

import "testing"

func find(a [8]int, x int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		switch {
		case a[mid] == x:
			return mid
		case a[mid] < x:
			hi = mid
		default:
			lo = mid
		}
	}
	return -1
}

func TestFind(t *testing.T) {
	a := [8]int{1, 3, 5, 7, 9, 11, 13, 15}
	for i, v := range a {
		if got := find(a, v); got != i {
			t.Errorf("find(%d) = %d, want %d", v, got, i)
		}
	}
	if find(a, 4) != -1 || find(a, 16) != -1 {
		t.Errorf("find should return -1 for missing values")
	}
}
