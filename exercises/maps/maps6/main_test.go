// maps6
// Make the tests pass!

// I AM NOT DONE
//
// enabled reports whether a feature is on; features that are not listed are off.
// Practices the zero value of a bool in a map.
package main_test

import "testing"

func enabled(flags map[string]bool, name string) bool {
	v, ok := flags[name]
	return ok
}

func TestEnabled(t *testing.T) {
	flags := map[string]bool{"beta": true, "legacy": false}
	if !enabled(flags, "beta") || enabled(flags, "legacy") || enabled(flags, "dark") {
		t.Errorf("enabled works incorrectly")
	}
}
