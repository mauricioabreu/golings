# Closures

A function literal can use the variables of the function around it, even after that function has returned. The function together with the variables it uses is a closure. The variables are not copied: the closure and the code that declared them share the same variables.

That is what makes a counter generator work, and it is also the source of some classic bugs: a variable declared outside a loop is shared by every closure built in the loop, and `x := x + 1` inside a closure declares a new `x` instead of changing the captured one. Since Go 1.22 the variables a `for` loop declares (`for i := ...`, `for _, v := range ...`) are new on every iteration, so closures that capture them each see their own.

* [A Tour of Go: Function closures](https://go.dev/tour/moretypes/25)
* [The Go specification: Function literals](https://go.dev/ref/spec#Function_literals)
* [Fixing for loops in Go 1.22](https://go.dev/blog/loopvar-preview)
* [Go FAQ: What happens with closures running as goroutines?](https://go.dev/doc/faq#closures_and_goroutines)
