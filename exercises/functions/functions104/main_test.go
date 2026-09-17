// functions104
// Make the tests pass!

// I AM NOT DONE
//
// execute must look a command up by name and run it,
// returning "unknown command" for an unknown one.
// Practices map[string]func and checking whether a key is present.
package main_test

import "testing"

var commands = map[string]func() string{
	"ping":    func() string { return "pong" },
	"version": func() string { return "1.0" },
}

func execute(name string) string {
	cmd := commands[name]
	return cmd()
}

func TestExecute(t *testing.T) {
	cases := map[string]string{"ping": "pong", "version": "1.0", "rm": "unknown command"}
	for in, want := range cases {
		if got := execute(in); got != want {
			t.Errorf("execute(%q) = %q, want %q", in, got, want)
		}
	}
}
