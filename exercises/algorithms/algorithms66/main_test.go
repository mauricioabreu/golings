// algorithms66
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: backtracking with swaps. Return all permutations of a set of
// distinct numbers (in any order).
// Expected asymptotics: O(n·n!) time, O(n) extra space.
package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func permutations(nums []int) [][]int {
	return nil
}

func TestPermutations(t *testing.T) {
	_, _ = fmt.Sprint, sort.Strings
	got := permutations([]int{1, 2, 3})
	var strs []string
	for _, p := range got {
		strs = append(strs, fmt.Sprint(p))
	}
	sort.Strings(strs)
	want := []string{"[1 2 3]", "[1 3 2]", "[2 1 3]", "[2 3 1]", "[3 1 2]", "[3 2 1]"}
	if !reflect.DeepEqual(strs, want) {
		t.Errorf("permutations = %v", strs)
	}
	if got := permutations([]int{7}); len(got) != 1 || got[0][0] != 7 {
		t.Errorf("single element = %v", got)
	}
	if got := permutations(nil); len(got) != 1 {
		t.Errorf("permutations(nil) = %v", got)
	}
}
