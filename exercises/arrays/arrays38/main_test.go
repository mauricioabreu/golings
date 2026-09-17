// arrays38
// Make the tests pass!

// I AM NOT DONE
//
// clearFirstColumn must zero the first column of a matrix.
// The matrix does not change.
// A range over an array of arrays yields copies of the rows.
package main_test

import "testing"

func clearFirstColumn(m *[3][3]int) {
	for _, row := range m {
		row[0] = 0
	}
}

func TestClearFirstColumn(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	clearFirstColumn(&m)
	if m != [3][3]int{{0, 2, 3}, {0, 5, 6}, {0, 8, 9}} {
		t.Errorf("m = %v", m)
	}
}
