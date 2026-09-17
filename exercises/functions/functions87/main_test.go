// functions87
// Make the tests pass!

// I AM NOT DONE
//
// loadUser must wrap the error so that errors.Is can find it.
// Practices fmt.Errorf with the %w verb.
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrNotFound = errors.New("not found")

func loadUser(id int) error {
	return fmt.Errorf("load user %d: %v", id, ErrNotFound)
}

func TestLoadUser(t *testing.T) {
	err := loadUser(7)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("errors.Is(loadUser(7), ErrNotFound) = false, want true")
	}
	if err.Error() != "load user 7: not found" {
		t.Errorf("message = %q", err.Error())
	}
}
