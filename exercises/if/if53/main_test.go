// if53
// Make the tests pass!

// I AM NOT DONE
//
// at must return the element at an index, or "" when the index is out of range.
// For an index equal to the length it panics.
// Practices the bounds of the valid indexes in a condition.
package main_test

import "testing"

func at(items []string, i int) string {
	if i >= 0 && i <= len(items) {
		return items[i]
	}
	return ""
}

func TestAt(t *testing.T) {
	items := []string{"a", "b", "c"}
	cases := map[int]string{0: "a", 2: "c", 3: "", -1: "", 10: ""}
	for in, want := range cases {
		if got := at(items, in); got != want {
			t.Errorf("at(%d) = %q, want %q", in, got, want)
		}
	}
}
