// generics67
// Make the tests pass!

// I AM NOT DONE
//
// FindMax returns the maximum, and the zero value for an empty slice.
// The caller cannot tell an empty slice from a slice whose maximum is 0.
// The zero value of T does not work as a marker of absence.
package main_test

import "testing"

type Ordered interface{ ~int | ~float64 }

func FindMax[T Ordered](s []T) T {
	var m T
	for i, v := range s {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func summarize(s []int) string {
	if FindMax(s) == 0 {
		return "empty"
	}
	return "has data"
}

func TestSummarize(t *testing.T) {
	if summarize(nil) != "empty" || summarize([]int{-5, 0}) != "has data" {
		t.Errorf("summarize: %q %q", summarize(nil), summarize([]int{-5, 0}))
	}
}
