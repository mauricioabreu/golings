// slices69
// Make the tests pass!

// I AM NOT DONE
//
// windowSums returns the sums of every window of size k.
// Practices a sliding window over a slice.
package main_test

import (
	"reflect"
	"testing"
)

func windowSums(s []int, k int) []int {
	if k > len(s) || k <= 0 {
		return nil
	}
	sum := 0
	for _, v := range s[:k] {
		sum += v
	}
	out := []int{sum}
	for i := k; i < len(s); i++ {
		sum += s[i]
		out = append(out, sum)
	}
	return out[1:]
}

func TestWindowSums(t *testing.T) {
	if got := windowSums([]int{1, 2, 3, 4, 5}, 3); !reflect.DeepEqual(got, []int{6, 9, 12}) {
		t.Errorf("windowSums = %v", got)
	}
	if got := windowSums([]int{1}, 2); got != nil {
		t.Errorf("windowSums(k>len) = %v", got)
	}
}
