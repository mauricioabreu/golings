// range9
// Make the tests pass!

// I AM NOT DONE
//
// secondLargest returns the second largest distinct value, or false.
// Practices two accumulators in one range.
package main_test

import "testing"

func secondLargest(s []int) (int, bool) {
	var first, second int
	hasFirst, hasSecond := false, false
	for _, v := range s {
		switch {
		case !hasFirst || v > first:
			first = v
			hasFirst = true
		case v < first && (!hasSecond || v > second):
			second, hasSecond = v, true
		}
	}
	return second, hasSecond
}

func TestSecondLargest(t *testing.T) {
	if v, ok := secondLargest([]int{3, 9, 7, 9}); !ok || v != 7 {
		t.Errorf("secondLargest = %d, %v; want 7", v, ok)
	}
	if v, ok := secondLargest([]int{1, 2}); !ok || v != 1 {
		t.Errorf("secondLargest([1 2]) = %d, %v", v, ok)
	}
	if _, ok := secondLargest([]int{5, 5}); ok {
		t.Errorf("secondLargest([5 5]) should fail")
	}
}
