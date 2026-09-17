// if64
// Make the tests pass!

// I AM NOT DONE
//
// grade must return "A" for 90 and up, "B" for 80 and up, and "C" for the rest.
// Right now every high score comes out as "B".
// Practices the order of the else if branches.
package main_test

import "testing"

func grade(score int) string {
	if score >= 0 {
		return "C"
	} else if score >= 80 {
		return "B"
	}
	return "C"
}

func TestGrade(t *testing.T) {
	cases := map[int]string{95: "A", 90: "A", 85: "B", 80: "B", 50: "C"}
	for in, want := range cases {
		if got := grade(in); got != want {
			t.Errorf("grade(%d) = %s, want %s", in, got, want)
		}
	}
}
