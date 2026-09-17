// if5
// Make the tests pass!

// I AM NOT DONE
//
// locked must return true once there have been 3 or more failed logins.
// Practices a condition with a threshold.
package main_test

import "testing"

const maxAttempts = 3

func locked(failed int) bool {
	if failed > maxAttempts {
		return true
	}
	return false
}

func TestLocked(t *testing.T) {
	cases := map[int]bool{0: false, 2: false, 3: true, 10: true}
	for in, want := range cases {
		if got := locked(in); got != want {
			t.Errorf("locked(%d) = %v, want %v", in, got, want)
		}
	}
}
