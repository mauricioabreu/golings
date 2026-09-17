// arrays53
// Make the tests pass!

// I AM NOT DONE
//
// backup must return the original contents of an array, even when the original
// was changed after the copy was taken.
// Assigning an array makes a full copy.
package main_test

import "testing"

func backup() ([3]int, [3]int) {
	data := [3]int{1, 2, 3}
	saved := data
	data[0] = 100
	return data, data
}

func TestBackup(t *testing.T) {
	data, saved := backup()
	if data[0] != 100 || saved[0] != 1 {
		t.Errorf("data = %v, saved = %v", data, saved)
	}
}
