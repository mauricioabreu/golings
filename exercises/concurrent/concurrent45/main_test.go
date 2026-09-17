// concurrent45
// Make the tests pass!

// I AM NOT DONE
//
// fetchAll handles the URLs in parallel, no more than limit at once,
// and returns the results in the original order.
// Practices a semaphore plus writing the results by index.
package main_test

import (
	"reflect"
	"strings"
	"sync"
	"testing"
)

func fetchAll(urls []string, limit int) []string {
	out := make([]string, len(urls))
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for i, u := range urls {
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()
			out = append(out, strings.ToUpper(u))
		}(i, u)
	}
	wg.Wait()
	return out
}

func TestFetchAll(t *testing.T) {
	got := fetchAll([]string{"a", "b", "c", "d"}, 2)
	if !reflect.DeepEqual(got, []string{"A", "B", "C", "D"}) {
		t.Errorf("fetchAll = %v", got)
	}
}
