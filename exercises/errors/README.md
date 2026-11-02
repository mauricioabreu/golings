# Errors

In Go an error is a value: the built-in `error` interface has a single method, `Error() string`. Functions that can fail return an `error` as their last result, and the caller checks it with `if err != nil`. There are no exceptions to catch.

Because errors are values, you can compare them, wrap them with more context (`fmt.Errorf("...: %w", err)`), and look inside the wrapped chain with `errors.Is` (is this particular error in there?) and `errors.As` (is there an error of this type in there, and give it to me). `errors.Join` combines several errors into one.

A *sentinel error* is a package-level variable such as `var ErrNotFound = errors.New("not found")`: callers compare against that exact value.

* [Error handling and Go](https://go.dev/blog/error-handling-and-go)
* [Working with errors in Go 1.13](https://go.dev/blog/go1.13-errors)
* [Package errors](https://pkg.go.dev/errors)
* [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
