# Contributing to golings

Thank you for the interest on contributing to `golings`

Below you can find some useful information if you want to contribute to the project, be it opening a new issue or adding code.

## Adding an exercise

Two steps are required to add a new exercise to the project: adding the exercise metadata to the `info.toml` and creating the respective exercise in the `exercises` folder.

Add the metadata for your exercise in the correct order. Exercises are run in the order they are defined in the `info.toml` file. If you are unsure about the order, ask help in the pull request.

Here is an example of an exercise being added:

```toml
[[exercises]]
name = "compile1"
path = "compile/compile1.go"
mode = "compile"
hint = "hints are cool"
```

The exercise mode is very important. It tells `golings` how to run the exercise. If you are adding an exercise that expects the user to only make it compilable, use `compile` mode. If it has a test suite and you need the user to actually have the tests passing, use `test`.

## Running the test suite

```sh
cd golings
go test -coverprofile=coverage.out -v $$(go list ./... | grep -v fixtures/error1)
```

If you have `make` installed it is easy as running `make test`

## Issues and pull requests

There are specific templates that will guide you through opening issues or pull requests.

Thank you! ❤️
