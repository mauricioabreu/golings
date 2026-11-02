# Defer, panic and recover

`defer` schedules a function call to run when the surrounding function returns, however it returns. It is how Go code releases things it acquired (files, locks, connections) right next to the line that acquired them. Deferred calls run in last-in, first-out order, and the arguments of a deferred call are evaluated when the `defer` statement runs, not when the call does.

`panic` stops the normal flow of a function; the deferred calls still run as the stack unwinds. `recover`, called inside a deferred function, stops the panic and returns its value. A deferred function can also read and change the named results of the function it belongs to.

In Go, errors are returned, not thrown. Use `panic` for situations that should never happen, and `recover` sparingly, for example to keep one failing request from taking a whole server down.

* [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
* [A Tour of Go: Defer](https://go.dev/tour/flowcontrol/12)
* [The Go specification: Defer statements](https://go.dev/ref/spec#Defer_statements)
* [Effective Go: Defer](https://go.dev/doc/effective_go#defer)
