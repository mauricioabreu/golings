// arrays89
// Make the tests pass!

// I AM NOT DONE
//
// The size of the buffer has to be computed at compile time from the number of headers.
// The code does not compile: len of a slice is not a constant.
// len of an array is a constant expression, len of a slice is not.
package main_test

import "testing"

var headers = []string{"id", "name", "email"}

const numHeaders = len(headers)

var widths [numHeaders]int

func TestNumHeaders(t *testing.T) {
	if numHeaders != 3 || len(widths) != 3 {
		t.Errorf("numHeaders = %d", numHeaders)
	}
}
