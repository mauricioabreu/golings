// if20
// Make the tests pass!

// I AM NOT DONE
//
// isDefaultRatio checks that a float32 ratio is 0.1.
// The comparison after a conversion to float64 never matches.
// 0.1 rounds differently in float32 and in float64.
package main_test

import "testing"

func isDefaultRatio(r float32) bool {
	if float64(r) == 0.1 {
		return true
	}
	return false
}

func TestIsDefaultRatio(t *testing.T) {
	var r float32 = 0.1
	if !isDefaultRatio(r) {
		t.Errorf("isDefaultRatio(0.1) = false, want true")
	}
	if isDefaultRatio(0.2) {
		t.Errorf("isDefaultRatio(0.2) = true, want false")
	}
}
