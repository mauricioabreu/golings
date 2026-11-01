// if1
// Make me compile!

package main_test

import "testing"

func bigger(a int, b int) int {
	// Complete this function to return the bigger number
	// Use only if statements
	if a > b {
		return a
	}
	return b
}

func TestTwoIsBiggerThanOne(t *testing.T) {
	if bigger(2, 1) != 2 {
		t.Errorf("2 is bigger than 1")
	}
}

func TestTenIsBiggerThanFive(t *testing.T) {
	if bigger(5, 10) != 10 {
		t.Errorf("10 is bigger than 5")
	}
}
