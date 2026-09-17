// arrays34
// Make the tests pass!

// I AM NOT DONE
//
// bubbleSort sorts an array in ascending order in place.
// Practices an array pointer and nested loops.
package main_test

import "testing"

func bubbleSort(a [6]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	a := [6]int{5, 2, 9, 1, 5, 6}
	bubbleSort(&a)
	if a != [6]int{1, 2, 5, 5, 6, 9} {
		t.Errorf("bubbleSort = %v", a)
	}
}
