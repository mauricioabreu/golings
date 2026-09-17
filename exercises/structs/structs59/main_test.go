// structs59
// Make the tests pass!

// I AM NOT DONE
//
// config returns the host and the port from an anonymous struct.
// Practices anonymous struct types.
package main_test

import "testing"

func config() (string, int) {
	c := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 80,
	}
	return c.Host, c.Port
}

func TestConfig(t *testing.T) {
	if h, p := config(); h != "localhost" || p != 5432 {
		t.Errorf("config = %s:%d", h, p)
	}
}
