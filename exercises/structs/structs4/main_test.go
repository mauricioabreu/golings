// structs4
// Make the tests pass!

// I AM NOT DONE
//
// Inc is called on an addressable value variable, and Go takes the address itself.
// Practices the automatic address-of in a call to a pointer method.
package main_test

import "testing"

type Hits struct{ N int }

func (h *Hits) Inc() {
	h.N += 2
}

func TestInc(t *testing.T) {
	var h Hits
	h.Inc()
	h.Inc()
	if h.N != 2 {
		t.Errorf("N = %d", h.N)
	}
}
