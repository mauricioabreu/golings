// functions38
// Make the tests pass!

// I AM NOT DONE
//
// The help command lists every command of the very table it is stored in.
// The code does not compile: "initialization cycle".
// Practices the initialization order of package variables and the init function.
package main_test

import (
	"sort"
	"strings"
	"testing"
)

var commands = map[string]func() string{
	"help": help,
	"ping": func() string { return "pong" },
}

func help() string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}

func TestHelp(t *testing.T) {
	if got := commands["help"](); got != "help,ping" {
		t.Errorf("help() = %q, want %q", got, "help,ping")
	}
	if got := commands["ping"](); got != "pong" {
		t.Errorf("ping() = %q", got)
	}
}
