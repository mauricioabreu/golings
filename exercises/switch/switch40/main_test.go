// switch40
// Make the tests pass!

// I AM NOT DONE
//
// romanValue returns the value of a single roman numeral.
// Practices a switch on a rune.
package main_test

import "testing"

func romanValue(r rune) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 100
	case 'C':
		return 100
	}
	return 0
}

func TestRomanValue(t *testing.T) {
	cases := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'Z': 0}
	for in, want := range cases {
		if got := romanValue(in); got != want {
			t.Errorf("romanValue(%c) = %d, want %d", in, got, want)
		}
	}
}
