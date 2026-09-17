// if22
// Make the tests pass!

// I AM NOT DONE
//
// isNotFound must recognize the ErrNotFound error.
// The check builds a new error with the same text and compares against that.
// Every call of errors.New creates a distinct value.
package main_test

import (
	"errors"
	"testing"
)

var ErrNotFound = errors.New("not found")

func find(id int) error {
	if id != 1 {
		return ErrNotFound
	}
	return nil
}

func isNotFound(err error) bool {
	if err == errors.New("not found") {
		return true
	}
	return false
}

func TestIsNotFound(t *testing.T) {
	if !isNotFound(find(2)) {
		t.Errorf("find(2) should be not found")
	}
	if isNotFound(find(1)) || isNotFound(errors.New("boom")) {
		t.Errorf("only ErrNotFound should match")
	}
}
