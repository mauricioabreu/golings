// variables93
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the Windows path C:\new\table exactly as written.
// In an interpreted string \n and \t turn into control characters.
// Practices raw string literals.
package main_test

import "testing"

func windowsPath() string {
	path := "C:\new\table"
	return path
}

func TestWindowsPath(t *testing.T) {
	want := "C:" + string(rune(92)) + "new" + string(rune(92)) + "table"
	if got := windowsPath(); got != want {
		t.Errorf("windowsPath() = %q, want %q", got, want)
	}
}
