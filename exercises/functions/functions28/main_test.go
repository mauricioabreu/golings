// functions28
// Make the tests pass!

// I AM NOT DONE
//
// split cuts a total in two parts: x = sum*4/9 and y = sum - x.
// Practices named results and a return with no expressions.
package main_test

import "testing"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum + x
	return
}

func TestSplit(t *testing.T) {
	x, y := split(18)
	if x != 8 || y != 10 {
		t.Errorf("split(18) = %d, %d; want 8, 10", x, y)
	}
}
