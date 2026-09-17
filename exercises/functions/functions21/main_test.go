// functions21
// Make the tests pass!

// I AM NOT DONE
//
// transform applies f to every element.
// Squares are wanted, but the wrong function is passed in.
// Practices passing a function as a value.
package main_test

import (
	"reflect"
	"testing"
)

func double(x int) int { return x * 2 }
func square(x int) int { return x * x }

func transform(nums []int, f func(int) int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = f(n)
	}
	return out
}

func squares(nums []int) []int {
	return transform(nums, double)
}

func TestSquares(t *testing.T) {
	if got := squares([]int{1, 3, 4}); !reflect.DeepEqual(got, []int{1, 9, 16}) {
		t.Errorf("squares(1,3,4) = %v, want [1 9 16]", got)
	}
}
