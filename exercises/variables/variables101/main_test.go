// variables101
// Make the tests pass!

// I AM NOT DONE
//
// digitHistogram must count how many times each digit occurs.
// Practices an array variable and turning a digit character into an index.
package main_test

import "testing"

func digitHistogram(s string) [10]int {
	var hist [10]int
	for _, r := range s {
		hist[r%10]++
	}
	return hist
}

func TestDigitHistogram(t *testing.T) {
	got := digitHistogram("+7 (900) 123-45-67")
	want := [10]int{2, 1, 1, 1, 1, 1, 1, 2, 0, 1}
	if got != want {
		t.Errorf("digitHistogram = %v, want %v", got, want)
	}
}
