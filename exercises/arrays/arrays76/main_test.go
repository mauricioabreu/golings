// arrays76
// Make the tests pass!

// I AM NOT DONE
//
// pascalRow returns the n-th row of Pascal's triangle in a [10]int,
// with the unused elements left at zero.
// Practices updating an array in place from right to left.
package main_test

import "testing"

func pascalRow(n int) [10]int {
	var row [10]int
	row[1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			row[j] += row[j-1]
		}
	}
	return row
}

func TestPascalRow(t *testing.T) {
	if got := pascalRow(4); got != [10]int{1, 4, 6, 4, 1} {
		t.Errorf("pascalRow(4) = %v", got)
	}
	if got := pascalRow(0); got != [10]int{1} {
		t.Errorf("pascalRow(0) = %v", got)
	}
}
