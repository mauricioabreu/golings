// maps21
// Make the tests pass!

// I AM NOT DONE
//
// snapshot saves the state of the settings so it can be rolled back to later.
// When the settings change, the snapshot changes too.
// Assigning a map copies the reference, not the contents.
package main_test

import "testing"

func snapshot(settings map[string]string) map[string]string {
	saved := settings
	return saved
}

func TestSnapshot(t *testing.T) {
	settings := map[string]string{"mode": "safe"}
	saved := snapshot(settings)
	settings["mode"] = "turbo"
	if saved["mode"] != "safe" {
		t.Errorf("snapshot changed: %v", saved)
	}
}
