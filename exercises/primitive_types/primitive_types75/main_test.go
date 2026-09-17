// primitive_types75
// Make the tests pass!

// I AM NOT DONE
//
// parseAmount must reject numbers that do not fit in a float64 ("1e400").
// ParseFloat returns +Inf along with an error, and the code ignores the error.
// Practices strconv.ParseFloat and the ErrRange error.
package main_test

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func parseAmount(s string) (float64, error) {
	v, _ := strconv.ParseFloat(s, 64)
	return v, nil
}

func TestParseAmount(t *testing.T) {
	_, _ = errors.New, math.IsInf
	if v, err := parseAmount("12.5"); err != nil || v != 12.5 {
		t.Errorf("parseAmount(12.5) = %v, %v", v, err)
	}
	for _, bad := range []string{"1e400", "Inf", "NaN", "abc"} {
		if _, err := parseAmount(bad); err == nil {
			t.Errorf("parseAmount(%s) should fail", bad)
		}
	}
}
