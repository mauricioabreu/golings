// functions70
// Make the tests pass!

// I AM NOT DONE
//
// gcd must compute the greatest common divisor with Euclid's algorithm.
// Practices recursive calls with changed arguments.
package main_test

import "testing"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(a%b, b)
}

func TestGCD(t *testing.T) {
	cases := []struct{ a, b, want int }{{12, 18, 6}, {17, 5, 1}, {10, 0, 10}, {0, 7, 7}}
	for _, c := range cases {
		if got := gcd(c.a, c.b); got != c.want {
			t.Errorf("gcd(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
