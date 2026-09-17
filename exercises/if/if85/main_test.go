// if85
// Make the tests pass!

// I AM NOT DONE
//
// samePoint must compare points by their coordinates.
// For distinct points holding equal values it returns false.
// == on pointers compares addresses, not values.
package main_test

import "testing"

type Point struct{ X, Y int }

func samePoint(a, b *Point) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a == b {
		return true
	}
	return false
}

func TestSamePoint(t *testing.T) {
	if !samePoint(&Point{1, 2}, &Point{1, 2}) {
		t.Errorf("equal points should match")
	}
	if samePoint(&Point{1, 2}, &Point{2, 1}) {
		t.Errorf("different points should not match")
	}
	if !samePoint(nil, nil) || samePoint(nil, &Point{}) {
		t.Errorf("nil handling is wrong")
	}
}
