// if29
// Make the tests pass!

// I AM NOT DONE
//
// letter for scores of 90-100: 97 and up is "A+", 93-96 "A", 90-92 "A-";
// anything lower is "B or lower".
// Practices nested ifs inside a branch.
package main_test

import "testing"

func letter(score int) string {
	if score >= 90 {
		if score > 97 {
			return "A+"
		}
		return "A-"
	}
	return "B or lower"
}

func TestLetter(t *testing.T) {
	cases := map[int]string{100: "A+", 97: "A+", 96: "A", 93: "A", 92: "A-", 90: "A-", 89: "B or lower"}
	for in, want := range cases {
		if got := letter(in); got != want {
			t.Errorf("letter(%d) = %s, want %s", in, got, want)
		}
	}
}
