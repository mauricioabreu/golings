// slices80
// Make the tests pass!

// I AM NOT DONE
//
// variants builds two paths that continue a shared prefix with different steps.
// The second variant overwrites the first one.
// Two appends to a slice with spare capacity use the same array.
package main_test

import (
	"reflect"
	"testing"
)

func extend(base []string, step string) []string {
	return append(base, step)
}

func variants(base []string) ([]string, []string) {
	return extend(base, "left"), extend(base, "right")
}

func TestVariants(t *testing.T) {
	base := make([]string, 0, 10)
	base = append(base, "start")
	l, r := variants(base)
	if !reflect.DeepEqual(l, []string{"start", "left"}) || !reflect.DeepEqual(r, []string{"start", "right"}) {
		t.Errorf("variants = %v, %v", l, r)
	}
}
