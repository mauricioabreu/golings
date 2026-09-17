// algorithms27
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: backtracking. Return all subsets of a set of distinct numbers.
// Order: by increasing size of the inclusion "mask"; elements within a
// subset follow input order.
// Expected asymptotics: O(n·2^n) time, O(n) extra space.
package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func subsets(nums []int) [][]int {
	return nil
}

func TestSubsets(t *testing.T) {
	_, _ = fmt.Sprint, sort.Strings
	got := subsets([]int{1, 2, 3})
	if len(got) != 8 {
		t.Fatalf("subsets count = %d, want 8", len(got))
	}
	var strs []string
	for _, s := range got {
		strs = append(strs, fmt.Sprint(s))
	}
	sort.Strings(strs)
	want := []string{"[1 2 3]", "[1 2]", "[1 3]", "[1]", "[2 3]", "[2]", "[3]", "[]"}
	if !reflect.DeepEqual(strs, want) {
		t.Errorf("subsets = %v", strs)
	}
	if got := subsets(nil); len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("subsets(nil) = %v", got)
	}
}
