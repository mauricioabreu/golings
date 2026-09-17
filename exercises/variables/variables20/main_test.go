// variables20
// Make the tests pass!

// I AM NOT DONE
//
// stats must return the count, the sum and the average.
// The average of an empty slice must be 0.
// Practices several accumulator variables and a division that cannot panic.
package main_test

import "testing"

func stats(nums []int) (count, sum int, mean float64) {
	for _, v := range nums {
		count++
		sum += v
	}
	mean = float64(sum / count)
	return
}

func TestStats(t *testing.T) {
	c, s, m := stats([]int{1, 2, 4})
	if c != 3 || s != 7 || m < 2.33 || m > 2.34 {
		t.Errorf("stats(1,2,4) = %d, %d, %v", c, s, m)
	}
	c, s, m = stats(nil)
	if c != 0 || s != 0 || m != 0 {
		t.Errorf("stats(nil) = %d, %d, %v; want 0, 0, 0", c, s, m)
	}
}
