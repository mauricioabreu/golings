// variables33
// Make the tests pass!

// I AM NOT DONE
//
// configure must set the global log level.
// The global variable is unchanged after the call.
// A local variable shadows the package variable of the same name.
package main_test

import "testing"

var logLevel = "info"

func configure(debug bool) {
	if debug {
		logLevel := "debug"
		_ = logLevel
	}
}

func TestConfigure(t *testing.T) {
	logLevel = "info"
	configure(false)
	if logLevel != "info" {
		t.Errorf("configure(false): logLevel = %q, want info", logLevel)
	}
	configure(true)
	if logLevel != "debug" {
		t.Errorf("configure(true): logLevel = %q, want debug", logLevel)
	}
}
