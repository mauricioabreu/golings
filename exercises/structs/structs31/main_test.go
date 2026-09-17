// structs31
// Make the tests pass!

// I AM NOT DONE
//
// Money marshals to the string "12.34 USD" through MarshalJSON. But when
// a value rather than a pointer is marshalled, the method is not called.
// Pointer methods are not part of the method set of a value.
package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Money struct {
	Cents    int
	Currency string
}

func (m *Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%d.%02d %s", m.Cents/100, m.Cents%100, m.Currency))
}

type Invoice struct {
	Total Money `json:"total"`
}

func TestInvoiceJSON(t *testing.T) {
	b, _ := json.Marshal(Invoice{Money{1234, "RUB"}})
	if string(b) != `{"total":"12.34 RUB"}` {
		t.Errorf("json = %s", b)
	}
}
