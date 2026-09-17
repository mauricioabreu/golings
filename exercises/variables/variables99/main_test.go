// variables99
// Make the tests pass!

// I AM NOT DONE
//
// pages must return the number of pages for total items at perPage per page.
// A partial page counts as a page too.
// Practices integer division rounding up.
package main_test

import "testing"

func pages(total, perPage int) int {
	return total / perPage
}

func TestPages(t *testing.T) {
	cases := []struct{ total, per, want int }{
		{25, 10, 3}, {30, 10, 3}, {0, 10, 0}, {1, 10, 1}, {5, 0, 0},
	}
	for _, c := range cases {
		if got := pages(c.total, c.per); got != c.want {
			t.Errorf("pages(%d, %d) = %d, want %d", c.total, c.per, got, c.want)
		}
	}
}
