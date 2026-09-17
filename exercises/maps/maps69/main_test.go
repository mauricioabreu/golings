// maps69
// Make the tests pass!

// I AM NOT DONE
//
// report formats the stock as the string "apple=3,kiwi=1,pear=5".
// The result comes out different every time.
// The iteration order of a map is undefined and deliberately randomized.
package main_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func report(stock map[string]int) string {
	var parts []string
	for k, v := range stock {
		parts = append(parts, k+"="+strconv.Itoa(v))
	}
	return strings.Join(parts, ",")
}

func TestReport(t *testing.T) {
	_ = sort.Strings
	stock := map[string]int{"pear": 5, "apple": 3, "kiwi": 1, "fig": 2, "lime": 4}
	for i := 0; i < 50; i++ {
		if got := report(stock); got != "apple=3,fig=2,kiwi=1,lime=4,pear=5" {
			t.Fatalf("report = %q", got)
		}
	}
}
