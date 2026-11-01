# Strings

A Go string is an immutable sequence of bytes, usually holding UTF-8 text. That one sentence explains most of the surprises: indexing gives a byte, `len` counts bytes, and a character such as `é` or `世` takes more than one byte. A `rune` is a single Unicode character, and `range` over a string yields runes.

The exercises also introduce `strings.Builder`, which builds a string without copying it on every step, and a few helpers from the `strings` package.

* [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
* [Package strings](https://pkg.go.dev/strings)
* [Package unicode/utf8](https://pkg.go.dev/unicode/utf8)
