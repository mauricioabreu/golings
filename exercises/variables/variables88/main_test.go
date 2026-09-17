// variables88
// Make the tests pass!

// I AM NOT DONE
//
// This function must return an error when any of the strings is not a number.
// The error inside the loop is lost, even though the parsing fails.
// Practices variable shadowing with the := operator.
package main_test

import (
	"strconv"
	"testing"
)

func sumAll(items []string) (int, error) {
	var err error
	total := 0
	for _, s := range items {
		n, err := strconv.Atoi(s)
		if err != nil {
			break
		}
		total += n
	}
	return total, err
}

func TestSumAll(t *testing.T) {
	if got, err := sumAll([]string{"1", "2", "3"}); err != nil || got != 6 {
		t.Errorf("sumAll(1,2,3) = %d, %v; want 6, nil", got, err)
	}
	if _, err := sumAll([]string{"1", "x", "3"}); err == nil {
		t.Errorf("sumAll(1,x,3) should return an error")
	}
}
