// maps55
// Make the tests pass!

// I AM NOT DONE
//
// userCount returns the number of users.
// Practices len on a map.
package main_test

import "testing"

func userCount(m map[int]string) int {
	return len(m) - 1
}

func TestUserCount(t *testing.T) {
	if got := userCount(map[int]string{1: "a", 2: "b", 3: "c"}); got != 3 {
		t.Errorf("userCount = %d", got)
	}
}
