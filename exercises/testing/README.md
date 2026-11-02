# Testing

Go ships with a testing framework in the standard library. Tests live in files that end in `_test.go`, and a test is a function `func TestXxx(t *testing.T)`. `go test` runs them.

A few habits make tests easy to live with: table-driven tests run the same checks over a list of cases, and `t.Run` gives each case its own name, so a failure says which one broke. Helpers that fail a test should call `t.Helper()`, so the failure points at the line that called the helper. `t.TempDir()` and `t.Cleanup()` put resources away when a test ends. Example functions are documentation that `go test` checks against their `// Output:` comment, and benchmarks measure how fast code is.

These exercises work a little differently from the rest: you are fixing testing code, and the checks around it run the code the way `go test` would.

* [Package testing](https://pkg.go.dev/testing)
* [Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests)
* [Testable Examples in Go](https://go.dev/blog/examples)
* [More predictable benchmarking with testing.B.Loop](https://go.dev/blog/testing-b-loop)
* [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)
