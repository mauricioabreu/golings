// variables45
// Make the tests pass!

// I AM NOT DONE
//
// capitalize must uppercase the first letter of an ASCII string.
// The code does not compile: a byte of a string cannot be assigned to.
// Practices string immutability and the conversion to []byte.
package main_test

import "testing"

func capitalize(s string) string {
	if s == "" {
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		s[0] -= 'a' - 'A'
	}
	return s
}

func TestCapitalize(t *testing.T) {
	cases := map[string]string{"gopher": "Gopher", "Go": "Go", "": "", "1st": "1st"}
	for in, want := range cases {
		if got := capitalize(in); got != want {
			t.Errorf("capitalize(%q) = %q, want %q", in, got, want)
		}
	}
}
