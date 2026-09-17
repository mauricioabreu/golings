// switch91
// Make the tests pass!

// I AM NOT DONE
//
// sizeLabel returns a size label. The code does not compile: two identical cases.
// The constants in the cases have to be distinct.
package main_test

import "testing"

func sizeLabel(size string) string {
	switch size {
	case "S":
		return "small"
	case "M":
		return "medium"
	case "S":
		return "large"
	}
	return "custom"
}

func TestSizeLabel(t *testing.T) {
	cases := map[string]string{"S": "small", "M": "medium", "L": "large", "XL": "custom"}
	for in, want := range cases {
		if got := sizeLabel(in); got != want {
			t.Errorf("sizeLabel(%s) = %s, want %s", in, got, want)
		}
	}
}
