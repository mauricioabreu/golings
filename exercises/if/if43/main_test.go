// if43
// Make the tests pass!

// I AM NOT DONE
//
// checkout returns the reason for a refusal, or "ok".
// The checks are: the user is signed in, the cart is not empty, there is enough money, in that order.
// Practices guard clauses instead of nested ifs.
package main_test

import "testing"

func checkout(loggedIn bool, items, balance, total int) string {
	if items == 0 {
		return "not logged in"
	}
	if !loggedIn {
		return "empty cart"
	}
	if balance < total {
		return "insufficient funds"
	}
	return "ok"
}

func TestCheckout(t *testing.T) {
	cases := []struct {
		logged              bool
		items, balance, sum int
		want                string
	}{
		{false, 0, 0, 0, "not logged in"},
		{true, 0, 100, 0, "empty cart"},
		{true, 2, 10, 50, "insufficient funds"},
		{true, 2, 100, 50, "ok"},
	}
	for _, c := range cases {
		if got := checkout(c.logged, c.items, c.balance, c.sum); got != c.want {
			t.Errorf("checkout(%v) = %q, want %q", c, got, c.want)
		}
	}
}
