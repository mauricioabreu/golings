// range31
// Make the tests pass!

// I AM NOT DONE
//
// parseConfig parses lines of the form "key = value", skipping empty lines
// and comments starting with #.
// Practices continue in a range over lines.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func parseConfig(text string) map[string]string {
	cfg := map[string]string{}
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		k, v, _ := strings.Cut(line, "=")
		cfg[k] = v
	}
	return cfg
}

func TestParseConfig(t *testing.T) {
	text := "# settings\nhost = localhost\n\nport=8080\n  # off = 1\nbroken line"
	want := map[string]string{"host": "localhost", "port": "8080"}
	if got := parseConfig(text); !reflect.DeepEqual(got, want) {
		t.Errorf("parseConfig = %v", got)
	}
}
