// algorithms7
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: DP over bits. Return a slice where element i is the number of
// set bits of the number i (for i from 0 to n).
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import (
	"reflect"
	"testing"
)

func countBits(n int) []int {
	return nil
}

func TestCountBits(t *testing.T) {
	if got := countBits(5); !reflect.DeepEqual(got, []int{0, 1, 1, 2, 1, 2}) {
		t.Errorf("countBits(5) = %v", got)
	}
	if got := countBits(0); !reflect.DeepEqual(got, []int{0}) {
		t.Errorf("countBits(0) = %v", got)
	}
	got := countBits(1000)
	if len(got) != 1001 || got[1000] != 6 {
		t.Errorf("countBits(1000) tail = %d", got[1000])
	}
}
