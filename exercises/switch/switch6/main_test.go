// switch6
// Make the tests pass!

// I AM NOT DONE
//
// chord names a chord from a list of notes. The code does not compile: slices
// cannot be used in a switch.
// The tag of a switch has to be of a comparable type.
package main_test

import (
	"strings"
	"testing"
)

func chord(notes []string) string {
	switch notes {
	case []string{"C", "E", "G"}:
		return "C major"
	case []string{"A", "C", "E"}:
		return "A minor"
	}
	return "unknown"
}

func TestChord(t *testing.T) {
	_ = strings.Join
	if got := chord([]string{"C", "E", "G"}); got != "C major" {
		t.Errorf("chord(CEG) = %s", got)
	}
	if got := chord([]string{"A", "C", "E"}); got != "A minor" {
		t.Errorf("chord(ACE) = %s", got)
	}
	if got := chord([]string{"C", "G"}); got != "unknown" {
		t.Errorf("chord(CG) = %s", got)
	}
}
