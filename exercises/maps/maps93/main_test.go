// maps93
// Make the tests pass!

// I AM NOT DONE
//
// defaultRegion must pick the region with the smallest code from a code to region map.
// Right now a different region comes out every time.
// The first element of a map iteration is undefined.
package main_test

import "testing"

func defaultRegion(regions map[int]string) string {
	for _, name := range regions {
		return name
	}
	return ""
}

func TestDefaultRegion(t *testing.T) {
	regions := map[int]string{77: "Moscow", 16: "Kazan", 78: "Saint Petersburg", 54: "Novosibirsk", 23: "Krasnodar"}
	for i := 0; i < 50; i++ {
		if got := defaultRegion(regions); got != "Kazan" {
			t.Fatalf("defaultRegion = %q, want Kazan", got)
		}
	}
	if got := defaultRegion(nil); got != "" {
		t.Errorf("defaultRegion(nil) = %q", got)
	}
}
