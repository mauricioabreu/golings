// variables97
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the n-th Fibonacci number (fib(0)=0, fib(1)=1).
// Practices assigning several variables at once.
package main_test

import "testing"

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a = b
		b = a + b
	}
	return a
}

func TestFib(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, w := range want {
		if got := fib(n); got != w {
			t.Errorf("fib(%d) = %d, want %d", n, got, w)
		}
	}
}
