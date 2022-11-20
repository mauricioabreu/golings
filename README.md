# golings

[![build and test](https://github.com/mauricioabreu/golings/actions/workflows/test.yml/badge.svg)](https://github.com/mauricioabreu/golings/actions/workflows/test.yml)

![gopher](misc/gopher-dance.gif)

> rustlings but for golang this time

You may know [rustlings](https://github.com/rust-lang/rustlings), right? If you don't, please go ahead and check out.
`rustlings` is awesome. It is a CLI app designed to teach the awesome Rust programming language through exercises.

`golings` has the very same idea, but for the [Go programming language](https://go.dev/)

After setting up all the tools required to run `golings` you have the task to fix tiny go programs.

## Installing

First, you need to have `go` installed. You can install it by visiting the [Go downloads page](https://go.dev/dl/)

There are two ways to install `golings`

### go install

```sh
go install github.com/mauricioabreu/golings/golings@latest
```

### Binaries

Go to the [releases page](https://github.com/mauricioabreu/golings/releases) and choose the option that best fits your environment.

## Doing exercises

All the exercises can be found in the directory `golings/exercises/<topic>`. For every topic there is an additional README file with some resources to get you started on the topic. We really recommend that you have a look at them before you start.

Now you have the task to fix all the programs. Some of them don't compile, and you need to fix them. Some of them compile, but have tests and you need to write some code to have them all green (these are the `compile` and `test` modes).

Clone the repository:

```sh
git clone git@github.com:mauricioabreu/golings.git
```

To run the next pending exercise:

```sh
golings run next
```

If you want to run a single exercise:

```sh
golings run variables1
```

In case you are stuck and need a hint:

```sh
golings hint variables1
```

To list all exercise while checking your progress:

```sh
golings list
```

To compile and run all the exercises:

```sh
golings verify
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md)

## Learning resources

* [Golang official tutorial](https://go.dev/doc/tutorial/)
* [Go by example](https://gobyexample.com)
* [Aprenda Go](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)

## Other 'lings

* [rustlings](https://github.com/rust-lang/rustlings)
* [ziglings](https://github.com/ratfactor/ziglings)
