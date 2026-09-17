// maps40
// Make the tests pass!

// I AM NOT DONE
//
// isEnabled: features are on by default, but may be switched off explicitly (false).
// Right now a feature switched off counts as on.
// Practices the difference between "the key holds false" and "the key is missing".
package main_test

import "testing"

func isEnabled(flags map[string]bool, name string) bool {
	if !flags[name] {
		return true
	}
	return flags[name]
}

func TestIsEnabled(t *testing.T) {
	flags := map[string]bool{"search": true, "chat": false}
	if !isEnabled(flags, "search") || !isEnabled(flags, "new") {
		t.Errorf("enabled and unknown features should be enabled")
	}
	if isEnabled(flags, "chat") {
		t.Errorf("explicitly disabled feature should be disabled")
	}
}
