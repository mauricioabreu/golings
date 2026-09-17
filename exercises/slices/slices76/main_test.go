// slices76
// Make the tests pass!

// I AM NOT DONE
//
// preview returns at most the first n comments. With fewer comments than that
// it panics.
// s[:n] with n > len(s) and n > cap panics.
package main_test

import "testing"

func preview(comments []string, n int) []string {
	return comments[:n]
}

func TestPreview(t *testing.T) {
	if got := preview([]string{"a", "b", "c"}, 2); len(got) != 2 {
		t.Errorf("preview(3, 2) = %v", got)
	}
	if got := preview([]string{"a"}, 5); len(got) != 1 {
		t.Errorf("preview(1, 5) = %v", got)
	}
}
