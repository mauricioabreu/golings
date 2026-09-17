// primitive_types45
// Make the tests pass!

// I AM NOT DONE
//
// upper turns lowercase greek letters into capitals by subtracting 32.
// For the accented letters and the final sigma the result is wrong.
// The offset between the cases is not the same for every character.
package main_test

import (
	"testing"
	"unicode"
)

func upper(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] >= 'α' && r[i] <= 'ω' {
			r[i] -= 32
		}
	}
	return string(r)
}

func TestUpper(t *testing.T) {
	_ = unicode.ToUpper
	cases := map[string]string{"μέλι": "ΜΈΛΙ", "φως": "ΦΩΣ", "ρόδο": "ΡΌΔΟ"}
	for in, want := range cases {
		if got := upper(in); got != want {
			t.Errorf("upper(%q) = %q, want %q", in, got, want)
		}
	}
}
