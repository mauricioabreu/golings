// functions37
// Make the tests pass!

// I AM NOT DONE
//
// joinInts must turn numbers into a comma separated string: "1,2,3".
// Practices the higher order function mapInts and strings.Join.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func mapInts(nums []int, f func(int) string) []string {
	out := make([]string, len(nums))

	for _, n := range nums {
		out = append(out, f(n))
	}
	return out
}

func joinInts(nums []int) string {
	return strings.Join(mapInts(nums, strconv.Itoa), "")
}

func TestJoinInts(t *testing.T) {
	if got := joinInts([]int{1, 2, 3}); got != "1,2,3" {
		t.Errorf("joinInts = %q, want %q", got, "1,2,3")
	}
}
