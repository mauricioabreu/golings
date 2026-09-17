// arrays10
// Make the tests pass!

// I AM NOT DONE
//
// The level names live in an array indexed by the level constant.
// Practices an array literal indexed by constants.
package main_test

import "testing"

const (
	Low = iota
	Mid
	High
)

var levelNames = [...]string{
	Low:  "low",
	Mid:  "high",
	High: "high",
}

func TestLevelNames(t *testing.T) {
	if levelNames[Low] != "low" || levelNames[Mid] != "mid" || levelNames[High] != "high" {
		t.Errorf("levelNames = %v", levelNames)
	}
}
