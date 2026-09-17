// switch53
// Make the tests pass!

// I AM NOT DONE
//
// handle deals with a command. The exit command has to stop the work
// even when it matches the previous command.
// A case may be a variable, and on a tie the first one wins.
package main_test

import "testing"

func handle(cmd, prev string) string {
	switch cmd {
	case prev:
		return "repeat"
	case "exit":
		return "bye"
	}
	return "run " + cmd
}

func TestHandle(t *testing.T) {
	cases := []struct{ cmd, prev, want string }{
		{"ls", "ls", "repeat"}, {"ls", "cd", "run ls"}, {"exit", "ls", "bye"}, {"exit", "exit", "bye"},
	}
	for _, c := range cases {
		if got := handle(c.cmd, c.prev); got != c.want {
			t.Errorf("handle(%s, %s) = %s, want %s", c.cmd, c.prev, got, c.want)
		}
	}
}
