// algorithms131
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: column addition. Add two binary numbers given as strings and
// return the result as a string without leading zeros (except "0" itself).
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

func addBinary(a, b string) string {
	return ""
}

func TestAddBinary(t *testing.T) {
	cases := []struct{ a, b, want string }{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"", "101", "101"},
		{"1111", "1111", "11110"},
	}
	for _, c := range cases {
		if got := addBinary(c.a, c.b); got != c.want {
			t.Errorf("addBinary(%q, %q) = %q, want %q", c.a, c.b, got, c.want)
		}
	}
}
