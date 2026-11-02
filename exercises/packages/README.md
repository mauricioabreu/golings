# Packages

Every Go file starts with `package name`, and the files of one folder form one package. Code from another package is reached through an import: `import "math"` lets you write `math.Pi`.

Go has no `public` or `private` keyword. A name that starts with an upper case letter, like `Pi` or `Println`, is exported: other packages can use it. A name that starts with a lower case letter stays inside its package.

When a package is loaded, its package-level variables are initialized first, in the order their dependencies need, and then its `init` functions run. A program can have many `init` functions, in any file. They run before `main` (or before the tests) and cannot be called by hand.

If two imports have the same package name, or you want a shorter one, give one of them an alias: `import cryptorand "crypto/rand"`.

* [A Tour of Go: Packages](https://go.dev/tour/basics/1) and [Exported names](https://go.dev/tour/basics/3)
* [The Go specification: Package initialization](https://go.dev/ref/spec#Package_initialization)
* [The Go specification: Import declarations](https://go.dev/ref/spec#Import_declarations)
* [Effective Go: Initialization](https://go.dev/doc/effective_go#initialization)
