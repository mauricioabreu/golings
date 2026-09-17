// primitive_types90
// Make the tests pass!

// I AM NOT DONE
//
// toHex must encode bytes as a string of hexadecimal digits.
// Practices picking nibbles out with a shift and a mask.
package main_test

import "testing"

const digits = "0123456789abcdef"

func toHex(data []byte) string {
	out := make([]byte, 0, len(data)*2)
	for _, b := range data {
		out = append(out, digits[b%10])
		out = append(out, digits[b/16%10])
	}
	return string(out)
}

func TestToHex(t *testing.T) {
	cases := map[string]string{"\x00\xff": "00ff", "Go": "476f", "": ""}
	for in, want := range cases {
		if got := toHex([]byte(in)); got != want {
			t.Errorf("toHex(%q) = %s, want %s", in, got, want)
		}
	}
}
