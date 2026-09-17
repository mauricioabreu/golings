// maps84
// Make the tests pass!

// I AM NOT DONE
//
// getOr returns the value for a key, or def when the key is missing.
// Practices comma-ok with a default value.
package main_test

import "testing"

func getOr(m map[string]string, k, def string) string {
	if v := m[k]; v != "" {
		return v
	}
	return def
}

func TestGetOr(t *testing.T) {
	m := map[string]string{"theme": "", "lang": "ru"}
	if getOr(m, "lang", "en") != "ru" || getOr(m, "font", "mono") != "mono" || getOr(m, "theme", "dark") != "" {
		t.Errorf("getOr works incorrectly")
	}
}
