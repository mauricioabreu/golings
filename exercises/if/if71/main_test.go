// if71
// Make the tests pass!

// I AM NOT DONE
//
// fits must check that used + size does not go over limit.
// For huge values the check misses the overflow.
// Practices overflow inside an if condition.
package main_test

import (
	"math"
	"testing"
)

func fits(used, size, limit int) bool {
	if used+size <= limit {
		return true
	}
	return false
}

func TestFits(t *testing.T) {
	if !fits(10, 20, 30) {
		t.Errorf("fits(10, 20, 30) should be true")
	}
	if fits(10, 21, 30) {
		t.Errorf("fits(10, 21, 30) should be false")
	}
	if fits(100, math.MaxInt, 1000) {
		t.Errorf("fits with huge size should be false")
	}
}
