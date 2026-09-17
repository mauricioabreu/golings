// functions77
// Make the tests pass!

// I AM NOT DONE
//
// fib must use the cache it is given and stay fast for n=80.
// Right now the cache is neither filled nor read.
// Practices passing a map into a function and memoized recursion.
package main_test

import "testing"

func fib(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	return fib(n-1, memo) + fib(n-2, memo)
}

func TestFibMemo(t *testing.T) {
	memo := map[int]int{}
	if got := fib(80, memo); got != 23416728348467685 {
		t.Errorf("fib(80) = %d", got)
	}
	if len(memo) != 79 {
		t.Errorf("memo should contain 79 entries, got %d", len(memo))
	}
}
