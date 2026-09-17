// arrays46
// Make the tests pass!

// I AM NOT DONE
//
// isIdentity checks that a 3 by 3 matrix is the identity matrix.
// Practices comparing elements against the condition i == j.
package main_test

import "testing"

func isIdentity(m [3][3]int) bool {
	for i := 0; i < 3; i++ {
		if m[i][i] != 1 {
			return false
		}
	}
	return true
}

func TestIsIdentity(t *testing.T) {
	if !isIdentity([3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}) {
		t.Errorf("identity not recognised")
	}
	if isIdentity([3][3]int{{1, 0, 0}, {0, 1, 5}, {0, 0, 1}}) {
		t.Errorf("non-identity matrix accepted")
	}
}
