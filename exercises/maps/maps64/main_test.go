// maps64
// Make the tests pass!

// I AM NOT DONE
//
// hasOwner checks whether any file belongs to a given user.
// Practices walking the values of a map.
package main_test

import "testing"

func hasOwner(files map[string]string, user string) bool {
	for name, owner := range files {
		_ = name
		if name == user {
			return true
		}
	}
	return false
}

func TestHasOwner(t *testing.T) {
	files := map[string]string{"a.txt": "ann", "b.txt": "bob"}
	if !hasOwner(files, "bob") || hasOwner(files, "a.txt") {
		t.Errorf("hasOwner works incorrectly")
	}
}
