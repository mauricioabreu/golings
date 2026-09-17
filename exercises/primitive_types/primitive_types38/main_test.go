// primitive_types38
// Make the tests pass!

// I AM NOT DONE
//
// areaCode must return the three digit code from a number such as "(495)1234567".
// Practices the string slice s[i:j], where j is excluded.
package main_test

import "testing"

func areaCode(phone string) string {
	return phone[0:3]
}

func TestAreaCode(t *testing.T) {
	if got := areaCode("(495)1234567"); got != "495" {
		t.Errorf("areaCode = %q, want 495", got)
	}
}
