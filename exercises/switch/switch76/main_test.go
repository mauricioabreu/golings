// switch76
// Make the tests pass!

// I AM NOT DONE
//
// statusText returns the text of an HTTP status code.
// Practices a switch on an integer.
package main_test

import "testing"

func statusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 403:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	}
	return ""
}

func TestStatusText(t *testing.T) {
	cases := map[int]string{200: "OK", 404: "Not Found", 500: "Internal Server Error", 403: ""}
	for in, want := range cases {
		if got := statusText(in); got != want {
			t.Errorf("statusText(%d) = %q, want %q", in, got, want)
		}
	}
}
