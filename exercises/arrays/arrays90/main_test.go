// arrays90
// Make the tests pass!

// I AM NOT DONE
//
// primes must return an array of the first five prime numbers, with the size
// inferred by the compiler from the literal.
// Practices the [...]T{...} literal.
package main_test

import "testing"

func primes() [5]int {
	p := [...]int{2, 3, 5, 7, 9}
	return p
}

func TestPrimes(t *testing.T) {
	if got := primes(); got != [5]int{2, 3, 5, 7, 11} {
		t.Errorf("primes = %v", got)
	}
}
