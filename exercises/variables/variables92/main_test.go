// variables92
// Make the tests pass!

// I AM NOT DONE
//
// This function must add up sensor readings, each within the range of an int8.
// For large totals the result suddenly turns negative.
// Practices overflow of fixed size integers.
package main_test

import "testing"

func totalReading(values []int8) int {
	var sum int8
	for _, v := range values {
		sum += v
	}
	return int(sum)
}

func TestTotalReading(t *testing.T) {
	if got := totalReading([]int8{100, 100, 100}); got != 300 {
		t.Errorf("totalReading(100,100,100) = %d, want 300", got)
	}
	if got := totalReading([]int8{-128, -128}); got != -256 {
		t.Errorf("totalReading(-128,-128) = %d, want -256", got)
	}
}
