// arrays6
// Make the tests pass!

// I AM NOT DONE
//
// spiral returns the elements of a 3 by 3 matrix in spiral order.
// Practices shrinking top/bottom/left/right bounds.
package main_test

import (
	"reflect"
	"testing"
)

func spiral(m [3][3]int) []int {
	var out []int
	top, bottom, left, right := 0, 2, 0, 2
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			out = append(out, m[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			out = append(out, m[r][right])
		}
		right--
		for c := left; c <= right; c++ {
			out = append(out, m[bottom][c])
		}
		bottom--
		if left <= right {
			for r := bottom; r >= top; r-- {
				out = append(out, m[r][left])
			}
			left++
		}
	}
	return out
}

func TestSpiral(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if got := spiral(m); !reflect.DeepEqual(got, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}) {
		t.Errorf("spiral = %v", got)
	}
}
