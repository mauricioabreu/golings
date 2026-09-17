// variables41
// Make the tests pass!

// I AM NOT DONE
//
// parsePort returns a port number and an error through named results.
// The code does not compile: the compiler complains about shadowed results.
// Practices named return values and shadowing with a naked return.
package main_test

import (
	"errors"
	"strconv"
	"testing"
)

func parsePort(s string) (n int, err error) {
	if s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		if n <= 0 || n > 65535 {
			err = errors.New("port out of range")
		}
		return
	}
	return 80, nil
}

func TestParsePort(t *testing.T) {
	if n, err := parsePort("8080"); err != nil || n != 8080 {
		t.Errorf("parsePort(8080) = %d, %v", n, err)
	}
	if n, err := parsePort(""); err != nil || n != 80 {
		t.Errorf("parsePort(\"\") = %d, %v; want 80", n, err)
	}
	if _, err := parsePort("70000"); err == nil {
		t.Errorf("parsePort(70000) should fail")
	}
}
