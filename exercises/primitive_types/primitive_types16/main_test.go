// primitive_types16
// Make the tests pass!

// I AM NOT DONE
//
// avg must return the average of two brightness values (0..255).
// For bright pixels the result is wrong.
// The intermediate result overflows a uint8.
package main_test

import "testing"

func avg(a, b uint8) uint8 {
	return (a + b) / 2
}

func TestAvg(t *testing.T) {
	cases := [][3]uint8{{10, 20, 15}, {200, 250, 225}, {255, 255, 255}, {0, 1, 0}}
	for _, c := range cases {
		if got := avg(c[0], c[1]); got != c[2] {
			t.Errorf("avg(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
