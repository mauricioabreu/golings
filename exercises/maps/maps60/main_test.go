// maps60
// Make the tests pass!

// I AM NOT DONE
//
// twoSum returns the indexes of two distinct elements adding up to target.
// Practices a value to index map for an O(n) search.
package main_test

import "testing"

func twoSum(nums []int, target int) (int, int, bool) {
	seen := map[int]int{}
	for i, v := range nums {
		seen[v] = i
		if j, ok := seen[target-v]; ok {
			return i, j, true
		}
	}
	return 0, 0, false
}

func TestTwoSum(t *testing.T) {
	if i, j, ok := twoSum([]int{3, 2, 4}, 6); !ok || i != 1 || j != 2 {
		t.Errorf("twoSum = %d, %d, %v; want 1, 2", i, j, ok)
	}
	if i, j, ok := twoSum([]int{3, 3}, 6); !ok || i != 0 || j != 1 {
		t.Errorf("twoSum([3 3]) = %d, %d, %v", i, j, ok)
	}
	if _, _, ok := twoSum([]int{1, 2}, 10); ok {
		t.Errorf("twoSum should fail")
	}
}
