// variables46
// Make the tests pass!

// I AM NOT DONE
//
// clock must split a number of seconds into hours, minutes and seconds.
// Practices integer division and remainders across several variables.
package main_test

import "testing"

func clock(total int) (h, m, s int) {
	h = total / 3600
	m = total / 60
	s = total % 3600
	return
}

func TestClock(t *testing.T) {
	cases := []struct{ total, h, m, s int }{
		{3661, 1, 1, 1}, {59, 0, 0, 59}, {7322, 2, 2, 2}, {0, 0, 0, 0},
	}
	for _, c := range cases {
		h, m, s := clock(c.total)
		if h != c.h || m != c.m || s != c.s {
			t.Errorf("clock(%d) = %d:%d:%d, want %d:%d:%d", c.total, h, m, s, c.h, c.m, c.s)
		}
	}
}
