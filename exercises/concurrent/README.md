# Concurrency

A goroutine is a function running at the same time as the rest of the program: `go f()`. Goroutines are cheap, and they talk to each other through channels (`ch <- value` to send, `<-ch` to receive) or share memory protected by the tools of the `sync` package.

A send on an unbuffered channel waits for a receiver, and a receive waits for a sender. If every goroutine is waiting, the program is deadlocked, and Go stops it with `fatal error: all goroutines are asleep - deadlock!`. A buffered channel lets that many sends go ahead without a receiver. `close` tells the receivers that nothing else is coming.

`select` waits on several channel operations at once, which is how you add a timeout or react to a `context` being cancelled. `sync.WaitGroup` waits for a group of goroutines to finish, and `sync.Mutex` keeps two goroutines from changing the same data at once. Run the tests with `-race` (golings does) to catch data races.

* [A Tour of Go: Concurrency](https://go.dev/tour/concurrency/1)
* [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
* [Go Concurrency Patterns: Context](https://go.dev/blog/context)
* [The Go memory model](https://go.dev/ref/mem)
* [Data Race Detector](https://go.dev/doc/articles/race_detector)
