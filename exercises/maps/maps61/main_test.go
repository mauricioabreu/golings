// maps61
// Make the tests pass!

// I AM NOT DONE
//
// register stores the email of a user; registering again overwrites the email.
// Map keys are unique.
package main_test

import "testing"

func register(m map[string]string, user, email string) {
	m[email] = user
}

func TestRegister(t *testing.T) {
	m := map[string]string{}
	register(m, "ann", "a@x")
	register(m, "ann", "b@x")
	if len(m) != 1 || m["ann"] != "b@x" {
		t.Errorf("m = %v", m)
	}
}
