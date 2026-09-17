// switch70
// Make the tests pass!

// I AM NOT DONE
//
// classify returns "neg", "zero" or "pos". The code does not compile:
// the cases hold boolean expressions while the switch runs on a number.
// switch x compares x against every case; conditions call for a tagless switch.
package main_test

import "testing"

func classify(n int) string {
	switch n {
	case n < 0:
		return "neg"
	case n == 0:
		return "zero"
	default:
		return "pos"
	}
}

func TestClassify(t *testing.T) {
	cases := map[int]string{-3: "neg", 0: "zero", 8: "pos"}
	for in, want := range cases {
		if got := classify(in); got != want {
			t.Errorf("classify(%d) = %s, want %s", in, got, want)
		}
	}
}
