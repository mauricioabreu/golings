// arrays81
// Make the tests pass!

// I AM NOT DONE
//
// friends stores friendships in a map keyed by [2]string. Friendship is symmetric,
// yet a lookup with the names the other way round finds nothing.
// Arrays as keys compare element by element, order included.
package main_test

import "testing"

type Friends map[[2]string]bool

func key(a, b string) [2]string {
	return [2]string{a, b}
}

func (f Friends) Add(a, b string)      { f[key(a, b)] = true }
func (f Friends) Are(a, b string) bool { return f[key(a, b)] }

func TestFriends(t *testing.T) {
	f := Friends{}
	f.Add("bob", "ann")
	if !f.Are("ann", "bob") || !f.Are("bob", "ann") {
		t.Errorf("friendship should be symmetric")
	}
	if f.Are("ann", "eve") {
		t.Errorf("ann and eve are not friends")
	}
}
