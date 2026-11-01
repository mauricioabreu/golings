# Interfaces

An interface is a set of method signatures. A type satisfies an interface just by having those methods: there is no `implements` keyword, which is why Go interfaces are small and easy to add after the fact. Functions that accept an interface work with every type that has the methods.

The empty interface, `any`, has no methods, so every value satisfies it. To get a concrete value back out you use a type assertion (`v.(int)`) or a type switch. An interface value is a pair of a type and a value, and that has a famous consequence: an interface holding a nil pointer is not nil.

Interfaces and structs can be built from smaller ones by embedding.

* [A Tour of Go: Interfaces](https://go.dev/tour/methods/9)
* [The Go specification: Interface types](https://go.dev/ref/spec#Interface_types)
* [The Go specification: Type assertions](https://go.dev/ref/spec#Type_assertions)
* [Effective Go: Interfaces and other types](https://go.dev/doc/effective_go#interfaces_and_types)
* [Go FAQ: Why is my nil error value not equal to nil?](https://go.dev/doc/faq#nil_error)
