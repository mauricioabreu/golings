// functions69
// Make the tests pass!

// I AM NOT DONE
//
// paths(r, c) must count the routes from the top left corner of an r by c grid
// to the bottom right one, moving only right and down.
// Practices recursion with two branches and base cases.
package main_test

import "testing"

func paths(r, c int) int {
	if r == 0 || c == 0 {
		return 1
	}
	return paths(r-1, c) * paths(r, c-1)
}

func TestPaths(t *testing.T) {
	cases := []struct{ r, c, want int }{{1, 1, 1}, {2, 2, 2}, {3, 3, 6}, {3, 7, 28}}
	for _, cs := range cases {
		if got := paths(cs.r, cs.c); got != cs.want {
			t.Errorf("paths(%d, %d) = %d, want %d", cs.r, cs.c, got, cs.want)
		}
	}
}
