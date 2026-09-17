// generics43
// Make the tests pass!

// I AM NOT DONE
//
// BinarySearch looks an element up in a sorted slice of any ordered type.
// Practices generic algorithms.
package main_test

import "testing"

type Ordered interface{ ~int | ~float64 | ~string }

func BinarySearch[T Ordered](s []T, x T) (int, bool) {
	lo, hi := 0, len(s)
	for lo < hi {
		mid := (lo + hi) / 2
		if s[mid] < x {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo, lo < len(s) && s[lo] == x
}

func TestBinarySearch(t *testing.T) {
	if i, ok := BinarySearch([]string{"a", "c", "e"}, "c"); !ok || i != 1 {
		t.Errorf("search c = %d, %v", i, ok)
	}
	if i, ok := BinarySearch([]int{1, 3, 5}, 4); ok || i != 2 {
		t.Errorf("search 4 = %d, %v", i, ok)
	}
}
