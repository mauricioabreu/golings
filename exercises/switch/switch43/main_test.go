// switch43
// Make the tests pass!

// I AM NOT DONE
//
// priority turns a label into a number: "high" is 1, "medium" 2, "low" 3, else 0.
// Practices a switch on a string.
package main_test

import "testing"

func priority(label string) int {
	switch label {
	case "high":
		return 1
	case "mid":
		return 2
	case "low":
		return 3
	}
	return 0
}

func TestPriority(t *testing.T) {
	cases := map[string]int{"high": 1, "medium": 2, "low": 3, "mid": 0}
	for in, want := range cases {
		if got := priority(in); got != want {
			t.Errorf("priority(%s) = %d, want %d", in, got, want)
		}
	}
}
