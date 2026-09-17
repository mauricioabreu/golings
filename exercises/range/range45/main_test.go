// range45
// Make the tests pass!

// I AM NOT DONE
//
// sumOrZero adds up the elements of an array through a pointer, and returns 0 for nil.
// A range with a value variable over a nil *[N]T panics.
package main_test

import "testing"

func sumOrZero(p *[4]int) int {
	s := 0
	for _, v := range p {
		s += v
	}
	return s
}

func TestSumOrZero(t *testing.T) {
	if got := sumOrZero(&[4]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("sumOrZero = %d", got)
	}
	if got := sumOrZero(nil); got != 0 {
		t.Errorf("sumOrZero(nil) = %d", got)
	}
}
