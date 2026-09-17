// variables60
// Make the tests pass!

// I AM NOT DONE
//
// The log levels must be numbered from one: Debug=1, Info=2, Warn=3.
// Practices constants and the iota generator.
package main_test

import "testing"

const (
	Debug = iota
	Info
	Warn
)

func TestLevels(t *testing.T) {
	if Debug != 1 || Info != 2 || Warn != 3 {
		t.Errorf("got Debug=%d Info=%d Warn=%d, want 1 2 3", Debug, Info, Warn)
	}
}
