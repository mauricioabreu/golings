// anonymous_functions76
// Make the tests pass!

// I AM NOT DONE
//
// snapshotSum has to remember the sum of a set at the moment of the call, and the literal
// reads the slice later, once the data has changed.
// A literal sees the current contents of a slice, not a snapshot.
package main_test

import "testing"

func snapshotSum(data []int) func() int {
	return func() int {
		s := 0
		for _, v := range data {
			s += v
		}
		return s
	}
}

func TestSnapshotSum(t *testing.T) {
	data := []int{1, 2, 3}
	sum := snapshotSum(data)
	data[0] = 100
	if got := sum(); got != 6 {
		t.Errorf("sum = %d, want 6", got)
	}
}
