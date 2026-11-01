# Pointers

A pointer holds the address of a value. `&x` gives the address of `x`, and `*p` follows a pointer `p` to the value it points to. A pointer to an `int` has the type `*int`.

Go passes everything by value: a function gets a copy of its arguments. To let a function change the caller's variable, pass a pointer to it. Pointers to structs are everywhere in Go, and `p.Field` works on a pointer without writing `(*p).Field`.

The zero value of a pointer is `nil`, and following a nil pointer panics, so code that may receive one has to check.

* [A Tour of Go: Pointers](https://go.dev/tour/moretypes/1)
* [The Go specification: Address operators](https://go.dev/ref/spec#Address_operators)
* [The Go specification: Pointer types](https://go.dev/ref/spec#Pointer_types)
* [Effective Go: Allocation with new](https://go.dev/doc/effective_go#allocation_new)
