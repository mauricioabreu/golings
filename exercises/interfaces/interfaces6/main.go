// interfaces6
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

type Logger struct {
	prefix string
}

func (l Logger) Log(message string) {
	fmt.Println(l.prefix + message)
}

// Service should be able to log. Embed a Logger in it: the methods of an
// embedded field are available on the outer struct.
type Service struct {
	name string
}

type Starter interface {
	Start()
}

type Stopper interface {
	Stop()
}

// Runner requires both methods. Build it by embedding the two interfaces
// above instead of repeating their methods.
type Runner interface {
	Starter
}

type Job struct{}

func (Job) Start() { fmt.Println("job started") }
func (Job) Stop()  { fmt.Println("job stopped") }

func run(r Runner) {
	r.Start()
	r.Stop()
}

func main() {
	service := Service{name: "api"}
	service.Log("service is up")

	run(Job{})
}
