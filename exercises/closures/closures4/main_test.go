// closures4
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// makeAdders returns one function per number: adders[i](x) adds numbers[i] to x.
//
// Since Go 1.22 the variables a for loop declares are new on every iteration,
// so a closure built inside the loop keeps the value of its own iteration.
// A variable declared outside the loop is not: all the closures share it.
func makeAdders(numbers []int) []func(int) int {
	adders := make([]func(int) int, len(numbers))

	// TODO: Each adder has to add its own number, not the one of the last iteration.
	var n int
	for i := range numbers {
		n = numbers[i]
		adders[i] = func(x int) int {
			return x + n
		}
	}

	return adders
}

func TestMakeAdders(t *testing.T) {
	adders := makeAdders([]int{1, 10, 100})

	tests := []struct{ adder, want int }{
		{adder: 0, want: 2},
		{adder: 1, want: 11},
		{adder: 2, want: 101},
	}
	for _, tc := range tests {
		if got := adders[tc.adder](1); got != tc.want {
			t.Errorf("adders[%d](1) = %d, want %d", tc.adder, got, tc.want)
		}
	}
}
