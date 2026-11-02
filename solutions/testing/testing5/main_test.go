// testing5
// Make the tests pass!

package main_test

import (
	"flag"
	"testing"
)

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

var setups int

func BenchmarkSum(b *testing.B) {
	setups++
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	for b.Loop() {
		sum(numbers)
	}
}

func TestBenchmarkLoopsWithBLoop(t *testing.T) {
	if err := flag.Set("test.benchtime", "10ms"); err != nil {
		t.Skip("can't shorten the benchmark:", err)
	}
	setups = 0

	result := testing.Benchmark(BenchmarkSum)

	if result.N == 0 {
		t.Fatal("the benchmark did not run")
	}
	if setups != 1 {
		t.Errorf("the benchmark function ran %d times: with b.N it is called again with a bigger N, with b.Loop it runs once", setups)
	}
}
