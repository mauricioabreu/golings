// variables81
// Make the tests pass!

// I AM NOT DONE
//
// rect must return the area and the perimeter of a 3 by 4 rectangle.
// Practices declaring several variables in one var statement.
package main_test

import "testing"

func rect() (area, perimeter int) {
	var w, h = 3, 3
	area = w + h
	perimeter = 2 * (w * h)
	return
}

func TestRect(t *testing.T) {
	area, perimeter := rect()
	if area != 12 || perimeter != 14 {
		t.Errorf("rect() = %d, %d; want 12, 14", area, perimeter)
	}
}
