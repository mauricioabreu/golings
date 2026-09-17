// switch25
// Make the tests pass!

// I AM NOT DONE
//
// suitSymbol returns the symbol of a suit.
// Practices a switch on a rune returning a string.
package main_test

import "testing"

func suitSymbol(s rune) string {
	switch s {
	case 'H':
		return "♥"
	case 'D':
		return "♦"
	case 'C':
		return "♣"
	case 'S':
		return "♣"
	}
	return "?"
}

func TestSuitSymbol(t *testing.T) {
	cases := map[rune]string{'H': "♥", 'D': "♦", 'C': "♣", 'S': "♠", 'X': "?"}
	for in, want := range cases {
		if got := suitSymbol(in); got != want {
			t.Errorf("suitSymbol(%c) = %s, want %s", in, got, want)
		}
	}
}
