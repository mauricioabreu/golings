// variables42
// Make the tests pass!

// I AM NOT DONE
//
// row must build a row of a table with the fields separated by a tab.
// Practices escape sequences in interpreted string literals.
package main_test

import "testing"

func row(name, city string) string {
	sep := `\t`
	return name + sep + city
}

func TestRow(t *testing.T) {
	got := row("Ann", "Oslo")
	if got != "Ann"+string(rune(9))+"Oslo" {
		t.Errorf("row(Ann, Oslo) = %q, want tab-separated", got)
	}
}
