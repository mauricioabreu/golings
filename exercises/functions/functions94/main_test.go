// functions94
// Make the tests pass!

// I AM NOT DONE
//
// subsets must return every subset of the numbers, using backtracking.
// The result holds damaged subsets.
// A stored slice shares its array with the buffer that keeps changing.
package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	var gen func(i int)
	gen = func(i int) {
		if i == len(nums) {
			res = append(res, cur)
			return
		}
		gen(i + 1)
		cur = append(cur, nums[i])
		gen(i + 1)
		cur = cur[:len(cur)-1]
	}
	gen(0)
	return res
}

func TestSubsets(t *testing.T) {
	var got []string
	for _, s := range subsets([]int{1, 2, 3}) {
		got = append(got, fmt.Sprint(s))
	}
	sort.Strings(got)
	want := []string{"[1 2 3]", "[1 2]", "[1 3]", "[1]", "[2 3]", "[2]", "[3]", "[]"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("subsets = %v, want %v", got, want)
	}
}
