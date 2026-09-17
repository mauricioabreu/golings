// functions17
// Make the tests pass!

// I AM NOT DONE
//
// countdown(n) must return [n, n-1, ..., 1] recursively.
// Right now a stray 0 shows up at the end.
// Practices the stopping condition of a recursion.
package main_test

import (
	"reflect"
	"testing"
)

func countdown(n int) []int {
	if n < 0 {
		return nil
	}
	return append([]int{n}, countdown(n-1)...)
}

func TestCountdown(t *testing.T) {
	if got := countdown(3); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("countdown(3) = %v, want [3 2 1]", got)
	}
	if got := countdown(0); len(got) != 0 {
		t.Errorf("countdown(0) = %v, want empty", got)
	}
}
