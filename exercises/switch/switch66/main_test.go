// switch66
// Make the tests pass!

// I AM NOT DONE
//
// bucket sorts a number into a bucket: "neg" below 0, "small" for 0..9, "big" for 10 and up.
// Practices a tagless switch in place of an if-else chain.
package main_test

import "testing"

func bucket(n int) string {
	switch {
	case n < 0:
		return "neg"
	case n <= 10:
		return "small"
	default:
		return "big"
	}
}

func TestBucket(t *testing.T) {
	cases := map[int]string{-1: "neg", 0: "small", 9: "small", 10: "big", 99: "big"}
	for in, want := range cases {
		if got := bucket(in); got != want {
			t.Errorf("bucket(%d) = %s, want %s", in, got, want)
		}
	}
}
