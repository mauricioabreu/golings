// if38
// Make the tests pass!

// I AM NOT DONE
//
// isLeap must implement the rule: divisible by 4 but not by 100, or divisible by 400.
// Practices compound boolean expressions with brackets.
package main_test

import "testing"

func isLeap(y int) bool {
	if (y%4 == 0 && y%100 != 0) || y%40 == 0 {
		return true
	}
	return false
}

func TestIsLeap(t *testing.T) {
	cases := map[int]bool{2024: true, 2023: false, 1900: false, 2000: true, 1960: true, 2040: true, 2100: false, 1800: false}
	for in, want := range cases {
		if got := isLeap(in); got != want {
			t.Errorf("isLeap(%d) = %v, want %v", in, got, want)
		}
	}
}
