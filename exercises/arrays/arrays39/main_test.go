// arrays39
// Make the tests pass!

// I AM NOT DONE
//
// parseIP turns "10.0.0.255" into a [4]byte, and formatIP builds the string back.
// Practices a fixed size array as a value type.
package main_test

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func parseIP(s string) ([4]byte, error) {
	var ip [4]byte
	parts := strings.Split(s, ".")
	for i, p := range parts {
		n, _ := strconv.Atoi(p)
		ip[i] = byte(n)
	}
	return ip, nil
}

func formatIP(ip [4]byte) string {
	parts := make([]string, 4)
	for i, b := range ip {
		parts[i] = strconv.Itoa(int(b))
	}
	return strings.Join(parts, ".")
}

func TestParseIP(t *testing.T) {
	_ = errors.New
	ip, err := parseIP("10.0.0.255")
	if err != nil || ip != [4]byte{10, 0, 0, 255} || formatIP(ip) != "10.0.0.255" {
		t.Errorf("parseIP = %v, %v", ip, err)
	}
	for _, bad := range []string{"1.2.3", "1.2.3.4.5", "1.2.3.256"} {
		if _, err := parseIP(bad); err == nil {
			t.Errorf("parseIP(%s) should fail", bad)
		}
	}
}
