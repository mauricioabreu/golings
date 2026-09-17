// functions32
// Make the tests pass!

// I AM NOT DONE
//
// reverse must reverse a string of runes recursively.
// Practices recursion over a rune slice and building the result.
package main_test

import "testing"

func reverse(s string) string {
	r := []rune(s)
	if len(r) <= 1 {
		return s
	}
	return string(r[0]) + reverse(string(r[1:]))
}

func reverseAll(items []string) []string {
	out := make([]string, len(items))
	for i, s := range items {
		out[i] = s
	}
	return out
}

func TestReverseAll(t *testing.T) {
	got := reverseAll([]string{"abc", "μία", ""})
	if got[0] != "cba" || got[1] != "αίμ" || got[2] != "" {
		t.Errorf("reverseAll = %q", got)
	}
}
