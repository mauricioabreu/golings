// switch39
// Make the tests pass!

// I AM NOT DONE
//
// sumValid adds up the numbers in the lines, skipping "#..." comments.
// A break in the switch does not skip the iteration, so comments break the sum.
// break leaves the switch, while continue moves on to the next iteration of the loop.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func sumValid(lines []string) (int, int) {
	sum, bad := 0, 0
	for _, l := range lines {
		switch {
		case strings.HasPrefix(l, "#"):
			break
		case l == "":
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			bad++
			continue
		}
		sum += n
	}
	return sum, bad
}

func TestSumValid(t *testing.T) {
	sum, bad := sumValid([]string{"1", "# comment", "", "2", "x"})
	if sum != 3 || bad != 1 {
		t.Errorf("sumValid = %d, %d; want 3, 1", sum, bad)
	}
}
