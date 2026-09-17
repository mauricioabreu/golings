// arrays85
// Make the tests pass!

// I AM NOT DONE
//
// rotate turns a 3 by 3 matrix 90 degrees clockwise.
// Practices working out the new indexes.
package main_test

import "testing"

func rotate(m [3][3]int) [3][3]int {
	r := m
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r[i][2-j] = m[i][j]
		}
	}
	return r
}

func TestRotate(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := [3][3]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	if got := rotate(m); got != want {
		t.Errorf("rotate = %v, want %v", got, want)
	}
}
