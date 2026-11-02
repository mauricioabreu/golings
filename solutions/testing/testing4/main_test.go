// testing4
// Make the tests pass!

package main_test

import "fmt"

// Greet says hello.
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// An example function is documentation that go test runs: it fails when what
// it prints is not what its "Output:" comment says.
func ExampleGreet() {
	fmt.Println(Greet("Gopher"))
	// Output: Hello, Gopher!
}

func ExampleGreet_everyone() {
	for _, name := range []string{"Ada", "Grace"} {
		fmt.Println(Greet(name))
	}
	// Output:
	// Hello, Ada!
	// Hello, Grace!
}
