// variables37
// Make the tests pass!

// I AM NOT DONE
//
// featureEnabled must understand "true", "1", "T" and the other accepted spellings.
// Anything else counts as disabled.
// Practices strconv.ParseBool and handling its error.
package main_test

import (
	"strconv"
	"testing"
)

func featureEnabled(v string) bool {
	_ = strconv.ParseBool
	return v == "true"
}

func TestFeatureEnabled(t *testing.T) {
	cases := map[string]bool{"true": true, "1": true, "T": true, "false": false, "yes": false, "": false}
	for in, want := range cases {
		if got := featureEnabled(in); got != want {
			t.Errorf("featureEnabled(%q) = %v, want %v", in, got, want)
		}
	}
}
