// algorithms1
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: backtracking. From distinct positive numbers, build all
// combinations that sum to target; each number may be used any number
// of times. Combinations must not repeat (numbers within each — in
// non-decreasing order).
// Expected asymptotics: O(n^(t/min)) time, O(t/min) space.
package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	return nil
}

func TestCombinationSum(t *testing.T) {
	_, _ = fmt.Sprint, sort.Strings
	got := combinationSum([]int{2, 3, 6, 7}, 7)
	var strs []string
	for _, c := range got {
		strs = append(strs, fmt.Sprint(c))
	}
	sort.Strings(strs)
	if !reflect.DeepEqual(strs, []string{"[2 2 3]", "[7]"}) {
		t.Errorf("combinationSum = %v", strs)
	}
	if got := combinationSum([]int{2}, 1); len(got) != 0 {
		t.Errorf("no combinations expected, got %v", got)
	}
	if got := combinationSum([]int{1}, 0); len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("zero target = %v", got)
	}
}
