// arrays24
// Make the tests pass!

// I AM NOT DONE
//
// merge combines two sorted arrays of 3 elements into one sorted array.
// Practices two indexes walking two arrays.
package main_test

import "testing"

func merge(a, b [3]int) [6]int {
	var out [6]int
	i, j := 0, 0
	for k := range out {
		if a[i] < b[j] {
			out[k] = a[i]
			i++
		} else {
			out[k] = b[j]
		}
	}
	return out
}

func TestMerge(t *testing.T) {
	if got := merge([3]int{1, 4, 9}, [3]int{2, 3, 10}); got != [6]int{1, 2, 3, 4, 9, 10} {
		t.Errorf("merge = %v", got)
	}
	if got := merge([3]int{1, 2, 3}, [3]int{4, 5, 6}); got != [6]int{1, 2, 3, 4, 5, 6} {
		t.Errorf("merge = %v", got)
	}
}
