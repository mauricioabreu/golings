// errors3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrNotFound = errors.New("not found")

func lookup(id int) (string, error) {
	if id != 1 {
		return "", ErrNotFound
	}
	return "Ada", nil
}

// loadUser looks the user up and adds context to the error: "loading user 7: not found".
// Whoever calls it must still be able to tell that the cause was ErrNotFound.
func loadUser(id int) (string, error) {
	name, err := lookup(id)
	if err != nil {
		// TODO: Wrap the error so that the original stays in the chain.
		return "", fmt.Errorf("loading user %d: %v", id, err)
	}
	return name, nil
}

func TestLoadUser(t *testing.T) {
	name, err := loadUser(1)

	if err != nil || name != "Ada" {
		t.Errorf("loadUser(1) = (%q, %v), want (Ada, nil)", name, err)
	}
}

func TestLoadUserAddsContext(t *testing.T) {
	_, err := loadUser(7)

	if err == nil || err.Error() != "loading user 7: not found" {
		t.Errorf("got %v, want %q", err, "loading user 7: not found")
	}
}

func TestLoadUserKeepsTheCause(t *testing.T) {
	_, err := loadUser(7)

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("errors.Is(%v, ErrNotFound) is false, want true", err)
	}
}
