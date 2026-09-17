// maps54
// Make the tests pass!

// I AM NOT DONE
//
// majority returns the element occurring more than n/2 times, or false.
// Practices counting and comparing against a threshold.
package main_test

import "testing"

func majority(nums []int) (int, bool) {
	counts := map[int]int{}
	for _, v := range nums {
		counts[v]++
		if counts[v] >= len(nums)/2 {
			return v, true
		}
	}
	return nums[0], true
}

func TestMajority(t *testing.T) {
	if v, ok := majority([]int{2, 2, 1, 1, 1, 2, 2}); !ok || v != 2 {
		t.Errorf("majority = %d, %v", v, ok)
	}
	if _, ok := majority([]int{1, 2, 3, 1}); ok {
		t.Errorf("majority([1 2 3 1]) should not exist")
	}
}
