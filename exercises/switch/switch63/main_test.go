// switch63
// Make the tests pass!

// I AM NOT DONE
//
// vat returns the VAT rate: the "food" category is 10, except for the "delicacy"
// subcategory which is 20; "books" is 0; everything else 20.
// Practices a switch inside a branch of a switch.
package main_test

import "testing"

func vat(category, sub string) int {
	switch category {
	case "food":
		return 10
	case "books":
		return 10
	}
	return 20
}

func TestVAT(t *testing.T) {
	cases := []struct {
		cat, sub string
		want     int
	}{{"food", "bread", 10}, {"food", "delicacy", 20}, {"books", "", 0}, {"toys", "", 20}}
	for _, c := range cases {
		if got := vat(c.cat, c.sub); got != c.want {
			t.Errorf("vat(%s, %s) = %d, want %d", c.cat, c.sub, got, c.want)
		}
	}
}
