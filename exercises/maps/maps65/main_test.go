// maps65
// Make the tests pass!

// I AM NOT DONE
//
// isomorphic checks whether b can be obtained from a by replacing characters,
// distinct characters going to distinct ones.
// Practices two maps for a one to one correspondence.
package main_test

import "testing"

func isomorphic(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	ab, ba := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < len(a); i++ {
		x, y := a[i], b[i]
		ab[x] = y
		ba[y] = x
	}
	return true
}

func TestIsomorphic(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{{"egg", "add", true}, {"foo", "bar", false}, {"paper", "title", true}, {"ab", "aa", false}}
	for _, c := range cases {
		if got := isomorphic(c.a, c.b); got != c.want {
			t.Errorf("isomorphic(%s, %s) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
