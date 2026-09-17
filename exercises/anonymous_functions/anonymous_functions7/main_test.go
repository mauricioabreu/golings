// anonymous_functions7
// Make the tests pass!

// I AM NOT DONE
//
// on registers a handler literal for an event, and emit calls it.
// Practices keeping literals in a map.
package main_test

import "testing"

func register(handlers map[string]func() string) {
	on := func(name string, h func() string) { handlers[name] = h }
	on("start", func() string { return "stopped" })
}

func TestHandlers(t *testing.T) {
	handlers := map[string]func() string{}
	register(handlers)
	if handlers["start"]() != "started" {
		t.Errorf("start handler = %q", handlers["start"]())
	}
}
