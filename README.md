# golings

![gopher](misc/gopher-dance.gif)

> rustlings but for golang this time

You may know [rustlings](https://github.com/rust-lang/rustlings), right? If you don't, please go ahead and check out.
`rustlings` is awesome. It is a CLI app designed to teach the awesome Rust programming language through exercises.

## Installing

For now the only option avaible to install `golings` is using the `go install` command:

```sh
go install github.com/mauricioabreu/golings/golings@latest
```

## Learning

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
