// maps72
// Make the tests pass!

// I AM NOT DONE
//
// online marks the users that are connected in a map[string]bool. On a logout
// the user is marked false, and the online count comes out wrong.
// len(map) counts every key, those holding false included.
package main_test

import "testing"

type Online map[string]bool

func (o Online) Login(u string) { o[u] = true }

func (o Online) Logout(u string) {
	o[u] = false
}

func (o Online) Count() int { return len(o) }

func TestOnline(t *testing.T) {
	o := Online{}
	o.Login("ann")
	o.Login("bob")
	o.Logout("ann")
	if o.Count() != 1 {
		t.Errorf("Count = %d, want 1", o.Count())
	}
}
