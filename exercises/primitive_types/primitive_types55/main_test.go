// primitive_types55
// Make the tests pass!

// I AM NOT DONE
//
// initial must return the first letter of a name as a string.
// Practices converting a rune to a string.
package main_test

import (
	"strconv"
	"testing"
)

func initial(name string) string {
	for _, r := range name {
		return strconv.Itoa(int(r))
	}
	return ""
}

func TestInitial(t *testing.T) {
	_ = strconv.Itoa
	cases := map[string]string{"Anna": "A", "Ωμέγα": "Ω", "": ""}
	for in, want := range cases {
		if got := initial(in); got != want {
			t.Errorf("initial(%q) = %q, want %q", in, got, want)
		}
	}
}
