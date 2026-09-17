// concurrent33
// Make the tests pass!

// I AM NOT DONE
//
// fetch returns a struct holding the data and an error through a channel.
// Practices channels of structs.
package main_test

import "testing"

type Result struct {
	Body string
	Err  error
}

func fetch(url string) Result {
	ch := make(chan Result, 1)
	go func() {
		ch <- Result{Err: nil, Body: ""}
	}()
	return <-ch
}

func TestFetch(t *testing.T) {
	if r := fetch("/a"); r.Err != nil || r.Body != "page /a" {
		t.Errorf("fetch = %+v", r)
	}
}
