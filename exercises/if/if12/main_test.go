// if12
// Make the tests pass!

// I AM NOT DONE
//
// canEdit: an administrator or the owner may edit, but only
// while the document is not locked.
// && binds tighter than ||.
package main_test

import "testing"

func canEdit(admin, owner, locked bool) bool {
	if admin || owner && !locked {
		return true
	}
	return false
}

func TestCanEdit(t *testing.T) {
	cases := []struct {
		admin, owner, locked, want bool
	}{
		{true, false, false, true}, {false, true, false, true},
		{true, false, true, false}, {false, true, true, false}, {false, false, false, false},
	}
	for _, c := range cases {
		if got := canEdit(c.admin, c.owner, c.locked); got != c.want {
			t.Errorf("canEdit(%v, %v, %v) = %v, want %v", c.admin, c.owner, c.locked, got, c.want)
		}
	}
}
