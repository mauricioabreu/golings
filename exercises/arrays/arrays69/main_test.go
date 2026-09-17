// arrays69
// Make the tests pass!

// I AM NOT DONE
//
// trace returns the sum of the main diagonal of a 3 by 3 matrix.
// Practices reaching the elements [i][i].
package main_test

import "testing"

func trace(m [3][3]int) int {
	s := 0
	for i := 0; i < 3; i++ {
		s += m[i][0]
	}
	return s
}

func TestTrace(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if got := trace(m); got != 15 {
		t.Errorf("trace = %d, want 15", got)
	}
}
