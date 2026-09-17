// if76
// Make the tests pass!

// I AM NOT DONE
//
// discount: club members get 10%, orders of 5000 or more get 5%,
// and a member with an order of 5000 or more gets 15%.
// Practices the order of the checks for combined conditions.
package main_test

import "testing"

func discount(member bool, total int) int {
	if member {
		return 10
	} else if total >= 5000 {
		return 5
	}
	return 0
}

func TestDiscount(t *testing.T) {
	cases := []struct {
		member bool
		total  int
		want   int
	}{{true, 6000, 15}, {true, 100, 10}, {false, 5000, 5}, {false, 100, 0}}
	for _, c := range cases {
		if got := discount(c.member, c.total); got != c.want {
			t.Errorf("discount(%v, %d) = %d, want %d", c.member, c.total, got, c.want)
		}
	}
}
