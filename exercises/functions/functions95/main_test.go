// functions95
// Make the tests pass!

// I AM NOT DONE
//
// power(base, exp) raises base to the power of exp.
// The call has the arguments the wrong way round.
// Practices positional arguments.
package main_test

import "testing"

func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func cubeOfTwo() int {
	return power(3, 2)
}

func TestCubeOfTwo(t *testing.T) {
	if got := cubeOfTwo(); got != 8 {
		t.Errorf("cubeOfTwo() = %d, want 8", got)
	}
}
