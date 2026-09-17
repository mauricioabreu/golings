// functions9
// Make the tests pass!

// I AM NOT DONE
//
// The variable config must hold the result of the anonymous function, not the function itself.
// The code does not compile.
// Practices immediately invoked function literals.
package main_test

import "testing"

func buildConfig() map[string]int {
	config := func() map[string]int {
		m := map[string]int{"workers": 4}
		m["retries"] = 3
		return m
	}
	return config
}

func TestBuildConfig(t *testing.T) {
	cfg := buildConfig()
	if cfg["workers"] != 4 || cfg["retries"] != 3 {
		t.Errorf("buildConfig() = %v", cfg)
	}
}
