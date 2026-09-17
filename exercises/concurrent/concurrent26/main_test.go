// concurrent26
// Make the tests pass!

// I AM NOT DONE
//
// wordCounts counts the words of several texts in parallel and merges them into one map.
// Practices a mutex around a shared map.
package main_test

import (
	"strings"
	"sync"
	"testing"
)

func wordCounts(texts []string) map[string]int {
	total := map[string]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, t := range texts {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			for _, w := range strings.Fields(t) {
				total[w]++
			}
		}(t)
	}
	wg.Wait()
	_ = &mu
	return total
}

func TestWordCounts(t *testing.T) {
	texts := make([]string, 20)
	for i := range texts {
		texts[i] = "go go gopher"
	}
	got := wordCounts(texts)
	if got["go"] != 40 || got["gopher"] != 20 {
		t.Errorf("counts = %v", got)
	}
}
