// range84
// Make the tests pass!

// I AM NOT DONE
//
// runeCount counts the characters of a string with a range.
// A range over a string walks the runes.
package main_test

import "testing"

func runeCount(s string) int {
	n := 0
	for range s {
		n += 2
	}
	return n
}

func TestRuneCount(t *testing.T) {
	cases := map[string]int{"go": 2, "γράφω": 5, "": 0}
	for in, want := range cases {
		if got := runeCount(in); got != want {
			t.Errorf("runeCount(%q) = %d, want %d", in, got, want)
		}
	}
}
