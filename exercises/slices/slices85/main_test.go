// slices85
// Make the tests pass!

// I AM NOT DONE
//
// median computes the median, but the caller's data comes back in a different order.
// sort.Ints sorts a slice in place, the caller's one included.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func median(s []int) float64 {
	c := s
	sort.Ints(c)
	n := len(c)
	if n%2 == 1 {
		return float64(c[n/2])
	}
	return float64(c[n/2-1]+c[n/2]) / 2
}

func TestMedian(t *testing.T) {
	data := []int{5, 1, 4, 2}
	if got := median(data); got != 3 {
		t.Errorf("median = %v, want 3", got)
	}
	if !reflect.DeepEqual(data, []int{5, 1, 4, 2}) {
		t.Errorf("data reordered: %v", data)
	}
}
