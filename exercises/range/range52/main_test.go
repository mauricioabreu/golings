// range52
// Make the tests pass!

// I AM NOT DONE
//
// normalizeAll folds the strings to lower case through a pointer to a slice.
// The code does not compile: a range over a *[]string is not allowed.
// A range works on a pointer to an array, but not on a pointer to a slice.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func normalizeAll(p *[]string) {
	for i, s := range p {
		p[i] = strings.ToLower(s)
	}
}

func TestNormalizeAll(t *testing.T) {
	s := []string{"Go", "RUST"}
	normalizeAll(&s)
	if !reflect.DeepEqual(s, []string{"go", "rust"}) {
		t.Errorf("normalizeAll = %v", s)
	}
}
