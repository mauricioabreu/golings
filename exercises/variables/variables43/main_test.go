// variables43
// Make the tests pass!

// I AM NOT DONE
//
// serverConfig must return an address of the form "host:port".
// The settings live in a variable of an anonymous struct type.
// Practices declaring a variable with an anonymous struct.
package main_test

import (
	"strconv"
	"testing"
)

func serverAddr() string {
	cfg := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 80,
	}
	return cfg.Host + ":" + strconv.Itoa(cfg.Port)
}

func TestServerAddr(t *testing.T) {
	if got := serverAddr(); got != "127.0.0.1:8080" {
		t.Errorf("serverAddr() = %q, want 127.0.0.1:8080", got)
	}
}
