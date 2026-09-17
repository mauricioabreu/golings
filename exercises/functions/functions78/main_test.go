// functions78
// Make the tests pass!

// I AM NOT DONE
//
// hasPlugin must report whether a plugin of that name is registered.
// The code does not compile: structs are compared with ==.
// Functions are not comparable, and neither are structs with function fields.
package main_test

import "testing"

type Plugin struct {
	Name string
	Run  func() string
}

func hasPlugin(ps []Plugin, target Plugin) bool {
	for _, p := range ps {
		if p == target {
			return true
		}
	}
	return false
}

func TestHasPlugin(t *testing.T) {
	gzip := Plugin{"gzip", func() string { return "compressed" }}
	auth := Plugin{"auth", func() string { return "ok" }}
	if !hasPlugin([]Plugin{gzip, auth}, Plugin{Name: "auth"}) {
		t.Errorf("auth plugin should be found")
	}
	if hasPlugin([]Plugin{gzip}, auth) {
		t.Errorf("auth plugin should not be found")
	}
}
