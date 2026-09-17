// structs39
// Make the tests pass!

// I AM NOT DONE
//
// An IsEven method is wanted on integers. The code does not compile:
// methods may only be declared on types of the same package.
// Methods need a named type of your own.
package main_test

import "testing"

func (n int) IsEven() bool { return n%2 == 0 }

func countEven(nums []int) int {
	c := 0
	for _, n := range nums {
		if n.IsEven() {
			c++
		}
	}
	return c
}

func TestCountEven(t *testing.T) {
	if got := countEven([]int{1, 2, 4, 7}); got != 2 {
		t.Errorf("countEven = %d", got)
	}
}
