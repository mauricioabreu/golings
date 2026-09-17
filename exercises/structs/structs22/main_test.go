// structs22
// Make the tests pass!

// I AM NOT DONE
//
// Range.Contains checks a number and Overlaps checks whether two [Lo, Hi] ranges meet.
// Practices predicate methods on a struct.
package main_test

import "testing"

type Range struct{ Lo, Hi int }

func (r Range) Contains(x int) bool {
	return x > r.Lo && x < r.Hi
}

func (r Range) Overlaps(o Range) bool {
	return r.Contains(o.Lo)
}

func TestRange(t *testing.T) {
	r := Range{1, 5}
	if !r.Contains(1) || !r.Contains(5) || r.Contains(6) {
		t.Errorf("Contains works incorrectly")
	}
	if !r.Overlaps(Range{0, 1}) || !r.Overlaps(Range{2, 3}) || r.Overlaps(Range{6, 9}) {
		t.Errorf("Overlaps works incorrectly")
	}
}
