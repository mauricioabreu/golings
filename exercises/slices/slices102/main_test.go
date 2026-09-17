// slices102
// Make the tests pass!

// I AM NOT DONE
//
// withBonus returns the first two scores with a bonus added, leaving the input alone.
// After the call the third element of the original slice has unexpectedly changed.
// An append to a subslice writes into the shared array while there is capacity.
package main_test

import (
	"reflect"
	"testing"
)

func withBonus(s []int, bonus int) []int {
	return append(s[:2], bonus)
}

func TestWithBonus(t *testing.T) {
	scores := []int{10, 20, 30}
	got := withBonus(scores, 5)
	if !reflect.DeepEqual(got, []int{10, 20, 5}) {
		t.Errorf("withBonus = %v", got)
	}
	if !reflect.DeepEqual(scores, []int{10, 20, 30}) {
		t.Errorf("scores modified: %v", scores)
	}
}
