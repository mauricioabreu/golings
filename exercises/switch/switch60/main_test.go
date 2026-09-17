// switch60
// Make the tests pass!

// I AM NOT DONE
//
// tariff picks a tariff by a struct key. The code does not compile:
// the struct holds a slice and is therefore not comparable.
// Only comparable types may be used in a switch.
package main_test

import "testing"

type key struct {
	Country string
	Tags    []string
}

func tariff(k key) int {
	switch k {
	case key{"RU", "vip"}:
		return 50
	case key{"RU", ""}:
		return 100
	}
	return 200
}

func TestTariff(t *testing.T) {
	if got := tariff(key{"RU", "vip"}); got != 50 {
		t.Errorf("tariff(RU vip) = %d", got)
	}
	if got := tariff(key{"RU", ""}); got != 100 {
		t.Errorf("tariff(RU) = %d", got)
	}
	if got := tariff(key{"US", ""}); got != 200 {
		t.Errorf("tariff(US) = %d", got)
	}
}
