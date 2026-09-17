// functions61
// Make the tests pass!

// I AM NOT DONE
//
// parseAll must return either all the numbers, or nil and an error.
// On an error it currently returns the numbers parsed so far.
// After a recover the named results still hold their intermediate values.
package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parseAll(items []string) (nums []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("bad input: %v", r)
		}
	}()
	for _, s := range items {
		nums = append(nums, mustAtoi(s))
	}
	return nums, nil
}

func TestParseAll(t *testing.T) {
	if nums, err := parseAll([]string{"1", "2"}); err != nil || len(nums) != 2 {
		t.Errorf("parseAll(1,2) = %v, %v", nums, err)
	}
	nums, err := parseAll([]string{"1", "x", "3"})
	if err == nil || nums != nil {
		t.Errorf("parseAll(1,x,3) = %v, %v; want nil and error", nums, err)
	}
}
