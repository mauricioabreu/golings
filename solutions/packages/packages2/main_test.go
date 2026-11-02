// packages2
// Make the tests pass!

package main_test

import "testing"

var config map[string]string

func init() {
	config = map[string]string{"env": "production", "region": "eu"}
}

func TestConfigIsLoadedBeforeTheTests(t *testing.T) {
	if config["env"] != "production" {
		t.Errorf(`config["env"] = %q, want "production"`, config["env"])
	}
	if config["region"] != "eu" {
		t.Errorf(`config["region"] = %q, want "eu"`, config["region"])
	}
}
