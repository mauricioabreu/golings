# Methods

A method is a function with a receiver: `func (r Rectangle) Area() int` is called as `r.Area()`. Go has no classes, but you can define methods on any type you declare yourself, not only on structs: `type Celsius float64` can have methods too.

A receiver is either a value (`r Rectangle`) or a pointer (`r *Rectangle`). A value receiver works on a copy, like any other argument. A pointer receiver can change the original, and it avoids copying big structs. The set of methods a type has depends on that choice: the method set of `T` holds the methods with a value receiver, and the method set of `*T` holds both. That matters as soon as you use interfaces.

* [A Tour of Go: Methods](https://go.dev/tour/methods/1)
* [The Go specification: Method declarations](https://go.dev/ref/spec#Method_declarations)
* [The Go specification: Method sets](https://go.dev/ref/spec#Method_sets)
* [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
