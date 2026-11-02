# CLAUDE.md — golings modernization

## What this project is
A fork of https://github.com/mauricioabreu/golings ("rustlings but for Go"), an interactive CLI that teaches Go through small broken exercises the learner fixes. Upstream is dormant (last commit 2025-03-12, Go 1.19, 2022-era deps). Goal: modernize it for current Go and fill the missing core topics. See ROADMAP.md for the ordered plan.

Upstream is Apache-2.0: keep LICENSE, keep attribution to the original author in README.

## Setup
- Fork upstream on GitHub, clone the fork into this folder (or move these two files into the clone). Add upstream as a remote: `git remote add upstream https://github.com/mauricioabreu/golings`.
- Work on a branch per roadmap phase, e.g. `phase-1-housekeeping`. Never commit straight to main.

## Repo layout (as of upstream HEAD)
- `golings/` — the CLI (Go module `github.com/cleonte/golings`): `cmd/`, `exercises/` (loader/runner code), `ui/`, `fixtures/`, `golings.go`.
- `exercises/<topic>/<name>/` — the learner-facing exercises (21 topics, 86 exercises).
- `info.toml` — the exercise manifest (name, dir, mode, hint). Every exercise must be listed here, in learning order.
- `Makefile` — `make test`, `make watch`, `make run`.
- `.github/workflows/` — `test.yml` (build and test) and `lint.yml`.
- Inspect one existing exercise and its `info.toml` entry before writing new ones, and copy that format exactly.

## Commands
- Test: `make test` (runs `go test` in `golings/`, excluding `fixtures/error1`, which fails on purpose).
- Run the tool: `make run` / `make watch`, or `go run golings/golings.go <command>`.
- Update deps: `cd golings && go get -u ./... && go mod tidy`.
- Modernize code: `go fix ./...` (Go 1.26+).
- Check formatting and vet: `gofmt -l .` and `go vet` on the CLI package.

## Rules for exercises
- Each exercise is a small, single-concept, broken program. The learner fixes it until it compiles and, where applicable, passes its checks or prints the expected output.
- Keep the Rustlings style: a short comment at the top says what to learn, an `// I AM NOT DONE`-style marker only if the existing format uses one, and a `hint` in `info.toml`.
- Use only the standard library in exercises.
- Idiomatic modern Go only: `any` not `interface{}`, `slices`/`maps` packages, `for i := range n`, `errors.Is/As/Join`, `log/slog`. No `ioutil`, no `rand.Seed`.
- For every new exercise, also create a working solution (outside the learner path, in a `solutions/` tree or whatever mechanism Phase 5 settles on) and verify the solution compiles and passes, and that the unsolved version fails in the intended way.
- Exercises that need a newer Go than the project floor must declare it (see ROADMAP Phase 2 and 5) so the tool can skip or warn instead of failing confusingly.
- Order matters: add new exercises to `info.toml` where they belong in the learning progression, not at the end.

## Working style
- Small commits, one logical change each. Run `make test` before every commit.
- Don't rewrite existing exercises' intent; modernize idioms and fix wording only.
- Don't add dependencies to the CLI without a clear reason; prefer stdlib.
- When unsure about the exercise format or manifest schema, read the code in `golings/exercises/` and an existing example, then follow it.
