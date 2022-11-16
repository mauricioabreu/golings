# golings

[![build and test](https://github.com/mauricioabreu/golings/actions/workflows/go.yml/badge.svg)](https://github.com/mauricioabreu/golings/actions/workflows/go.yml)

![gopher](misc/gopher-dance.gif)

> rustlings but for golang this time

You may know [rustlings](https://github.com/rust-lang/rustlings), right? If you don't, please go ahead and check out.
`rustlings` is awesome. It is a CLI app designed to teach the awesome Rust programming language through exercises.

`golings` has the very same idea, but for the [Go programming language](https://go.dev/)

## Installing

First, you need to have `go` installed. You can install it by visiting the [Go downloads page](https://go.dev/dl/)

There are two ways to install

### go install

```sh
go install github.com/mauricioabreu/golings/golings@latest
```

### Binaries

Go to the [releases page](https://github.com/mauricioabreu/golings/releases) and choose the option that best fits your environment.

## Doing exercises

First, clone this repository

```sh
git clone git@github.com:mauricioabreu/golings.git
```

All the exercises can be found in the directory `golings/exercises/<topic>`.

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
