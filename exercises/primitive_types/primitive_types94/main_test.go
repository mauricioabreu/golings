// primitive_types94
// Make the tests pass!

// I AM NOT DONE
//
// humanBytes formats a size: "512 B", "1.5 KiB", "2.0 MiB".
// Practices float division and formatting with one decimal place.
package main_test

import (
	"fmt"
	"testing"
)

func humanBytes(n int64) string {
	if n < 1024 {
		return fmt.Sprintf("%d B", n)
	}
	units := []string{"KiB", "MiB", "GiB"}
	v := float64(n) / 1024
	i := 0
	for v > 1024 {
		v /= 1000
		i++
	}
	return fmt.Sprintf("%.0f %s", v, units[i])
}

func TestHumanBytes(t *testing.T) {
	cases := map[int64]string{512: "512 B", 1536: "1.5 KiB", 2 << 20: "2.0 MiB", 1024: "1.0 KiB", 3 << 30: "3.0 GiB"}
	for in, want := range cases {
		if got := humanBytes(in); got != want {
			t.Errorf("humanBytes(%d) = %q, want %q", in, got, want)
		}
	}
}
