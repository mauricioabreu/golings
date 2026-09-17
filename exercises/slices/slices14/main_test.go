// slices14
// Make the tests pass!

// I AM NOT DONE
//
// copied must return the number of elements that were copied.
// copy returns the number of elements copied.
package main_test

import "testing"

func copied(dst, src []int) int {
	n := copy(dst, src)
	return len(src)
}

func TestCopied(t *testing.T) {
	if got := copied(make([]int, 2), []int{1, 2, 3}); got != 2 {
		t.Errorf("copied = %d, want 2", got)
	}
}
