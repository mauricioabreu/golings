// if40
// Make the tests pass!

// I AM NOT DONE
//
// batteryStatus: while charging it is "charging", or "full" at 100%;
// on battery it is "low" up to 15% and "ok" otherwise.
// Practices nested ifs over several signals.
package main_test

import "testing"

func batteryStatus(level int, charging bool) string {
	if level < 15 {
		return "low"
	}
	if charging {
		return "charging"
	}
	return "ok"
}

func TestBatteryStatus(t *testing.T) {
	cases := []struct {
		level    int
		charging bool
		want     string
	}{{100, true, "full"}, {10, true, "charging"}, {50, true, "charging"}, {10, false, "low"}, {15, false, "ok"}}
	for _, c := range cases {
		if got := batteryStatus(c.level, c.charging); got != c.want {
			t.Errorf("batteryStatus(%d, %v) = %s, want %s", c.level, c.charging, got, c.want)
		}
	}
}
