// algorithms113
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: the sieve of Eratosthenes. Count the number of primes strictly
// less than n. The solution must be faster than trial division.
// Expected asymptotics: O(n·log log n) time, O(n) space.
package main_test

import "testing"

func countPrimes(n int) int {
	return 0
}

func TestCountPrimes(t *testing.T) {
	cases := map[int]int{10: 4, 0: 0, 1: 0, 2: 0, 3: 1, 100: 25, 1000000: 78498}
	for in, want := range cases {
		if got := countPrimes(in); got != want {
			t.Errorf("countPrimes(%d) = %d, want %d", in, got, want)
		}
	}
}
