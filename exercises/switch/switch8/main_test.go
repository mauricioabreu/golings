// switch8
// Make the tests pass!

// I AM NOT DONE
//
// heading works out the level of a Markdown heading ("# " is 1, "## " is 2, "### " is 3),
// and 0 for every other line.
// Practices counting before a switch and checking the details inside a branch.
package main_test

import "testing"

func heading(line string) int {
	n := 0
	for n < len(line) && line[n] == '#' {
		n++
	}
	switch n {
	case 1, 2, 3:
		return n
	}
	return 0
}

func TestHeading(t *testing.T) {
	cases := map[string]int{"# Title": 1, "### Sub": 3, "#### Deep": 0, "#tag": 0, "text": 0, "##": 0}
	for in, want := range cases {
		if got := heading(in); got != want {
			t.Errorf("heading(%q) = %d, want %d", in, got, want)
		}
	}
}
