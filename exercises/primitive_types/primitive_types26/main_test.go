// primitive_types26
// Make the tests pass!

// I AM NOT DONE
//
// plain must print a number with no exponent and no stray zeros: 1e21 -> "1000000000000000000000".
// %v switches to the exponent form for large float64 values.
package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func plain(x float64) string {
	return fmt.Sprintf("%v", x)
}

func TestPlain(t *testing.T) {
	_, _ = fmt.Sprint, strconv.FormatFloat
	cases := map[float64]string{1e21: "1000000000000000000000", 0.000001: "0.000001", 2.5: "2.5", 100: "100"}
	for in, want := range cases {
		if got := plain(in); got != want {
			t.Errorf("plain(%v) = %s, want %s", in, got, want)
		}
	}
}
