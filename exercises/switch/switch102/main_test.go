// switch102
// Make the tests pass!

// I AM NOT DONE
//
// command parses a line read from a file. Lines from Windows files
// end in "\r\n", and the commands are not recognized.
// Practices invisible characters in a switch tag.
package main_test

import (
	"strings"
	"testing"
)

func command(line string) string {
	switch strings.TrimRight(line, "\n") {
	case "start":
		return "starting"
	case "stop":
		return "stopping"
	}
	return "unknown"
}

func TestCommand(t *testing.T) {
	cases := map[string]string{"start\n": "starting", "stop\r\n": "stopping", "start\r\n": "starting", "reset\n": "unknown"}
	for in, want := range cases {
		if got := command(in); got != want {
			t.Errorf("command(%q) = %s, want %s", in, got, want)
		}
	}
}
