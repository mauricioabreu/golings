// primitive_types79
// Make the tests pass!

// I AM NOT DONE
//
// bucket returns the number of the interval of width size that x falls into:
// [0, size) is 0 and [-size, 0) is -1.
// int(x) truncates towards zero rather than rounding down.
package main_test

import (
	"math"
	"testing"
)

func bucket(x, size float64) int {
	return int(x / size)
}

func TestBucket(t *testing.T) {
	_ = math.Floor
	cases := map[float64]int{0: 0, 9.9: 0, 10: 1, -0.5: -1, -10: -1, -10.1: -2}
	for in, want := range cases {
		if got := bucket(in, 10); got != want {
			t.Errorf("bucket(%v, 10) = %d, want %d", in, got, want)
		}
	}
}
