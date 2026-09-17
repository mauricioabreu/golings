// arrays64
// Make the tests pass!

// I AM NOT DONE
//
// zerosToEnd moves every zero to the end of the array, keeping the order of the rest.
// Practices writing through a second index and filling the tail.
package main_test

import "testing"

func zerosToEnd(a [6]int) [6]int {
	w := 0
	for _, v := range a {
		if v != 0 {
			a[w] = v
			w++
		}
	}
	return a
}

func TestZerosToEnd(t *testing.T) {
	if got := zerosToEnd([6]int{0, 1, 0, 3, 12, 0}); got != [6]int{1, 3, 12, 0, 0, 0} {
		t.Errorf("zerosToEnd = %v", got)
	}
	if got := zerosToEnd([6]int{4, 0, 0, 0, 0, 5}); got != [6]int{4, 5, 0, 0, 0, 0} {
		t.Errorf("zerosToEnd = %v", got)
	}
}
