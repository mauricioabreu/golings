// functions55
// Make the tests pass!

// I AM NOT DONE
//
// reset must return the package settings to their initial state.
// The code does not compile: init is called explicitly.
// init functions are called by the runtime only and cannot be referred to by name.
package main_test

import "testing"

var settings map[string]string

func init() {
	settings = map[string]string{"theme": "dark", "lang": "en"}
}

func reset() {
	init()
}

func TestReset(t *testing.T) {
	if settings["theme"] != "dark" {
		t.Fatalf("settings not initialised: %v", settings)
	}
	settings["theme"] = "light"
	delete(settings, "lang")
	reset()
	if settings["theme"] != "dark" || settings["lang"] != "en" {
		t.Errorf("after reset settings = %v", settings)
	}
}
