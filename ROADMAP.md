# ROADMAP — modernizing golings

Baseline (upstream HEAD): last commit 2025-03-12, `go 1.19` in go.mod and CI, `actions/checkout@v3` / `setup-go@v3`, deps from 2022 (cobra 1.6, fsnotify 1.6, go-pretty 6.4, ginkgo 2.5, progressbar 3.12, go-toml 2.0.5, fatih/color 1.13). 43 exercises in 13 topics: variables, primitive_types, functions, anonymous_functions, if, switch, arrays, slices, maps, range, structs, generics (2), concurrent (3).

Go version policy: set the project floor at **Go 1.26** (latest deps such as x/sys, x/term, x/text require it; raised from the original 1.24 plan). Exercises needing a newer Go than the floor are tagged with a minimum Go version. CI tests a matrix: the floor and the latest stable.

## Phase 1 — Housekeeping
- [x] Fork, add `upstream` remote, create branch.
- [x] `go.mod`: `go 1.26`; update every dependency (`go get -u ./... && go mod tidy`); fix any breaking API changes (cobra, go-pretty, progressbar).
- [x] CI: `actions/checkout@v4+`, `actions/setup-go@v5+`, version matrix (1.26, stable), cache on; keep the `fixtures/error1` exclusion.
- [x] Run `go fix ./...` and `gofmt`; replace `interface{}` with `any`; remove deprecated APIs.
- [x] Verify all 43 existing exercises still behave: unsolved fails correctly (done); solved passes (deferred to Phase 5, no solutions yet).
- [x] README: modern install (`go install github.com/cleonte/golings/golings@latest`), fork notice, supported Go versions.
- [x] Release config (GoReleaser) check; tag a release so binaries exist (v0.9.1, release workflow on `v*` tags).
Exit: `make test` green on the floor and latest Go.

## Phase 2 — Tooling and learner experience
- [x] `golings doctor`: check Go version against the floor and per-exercise minimums, check the workspace.
- [x] Per-exercise `go_version` (min) field in `info.toml`; runner skips or warns when the local Go is too old.
- [x] `golings reset <exercise>` (and `golings init`, which creates the exercises folder from the embedded copy): restore an exercise to its original broken state.
- [x] Solutions behind an explicit flag (`golings solution <exercise>`), not shown by default. (All 43 exercises have a verified solution.)
- [x] Clearer compile-error output (show file:line, trim noise).
- [x] Progress shown by topic; fix the completion-state edge cases.

## Phase 2.5 — Terminal UI (rustlings-style `watch`)
Prerequisites (each its own small PR):
- [ ] Runner: build first, then run the compiled binary under a timeout; `RunContext(ctx)` so a new save can cancel a run; `ErrTimeout` with a "didn't finish, possible infinite loop" message; the whole process tree is killed (test: no leftover process). Record whether the build step failed on `Result` instead of parsing `[build failed]`.
- [ ] Watcher: extract `WatchEvents` (reads the working dir, `log.Fatal`s) into a testable function with a root folder, a context and an events channel; debounce a burst of events; treat create/rename as changes so atomic saves (vim, JetBrains) are seen. Fix the goroutine leak in `watch` (one more `for range update` goroutine is started per typed command).
The TUI (Bubble Tea, Lip Gloss, Bubbles; check each module's current major version and `go` directive first):
- [ ] Minimal TUI: header with overall and per-topic progress, output pane for the current exercise, auto re-run on save, quit. `Update` unit tested without a terminal.
- [ ] Keys: hint, list/jump, reset (with confirmation), reveal solution, check all, next.
- [ ] `--plain` line mode kept for pipes and CI; `watch` falls back to it when stdin is not a terminal.

## Phase 3 — Missing core topics (highest value)
Add in learning order; 3–6 exercises each, hints in `info.toml`.
1. **strings** — runes vs bytes, `strings.Builder`, `strings.Fields`, `SplitSeq`/`Lines` (1.24).
2. **pointers** — address/deref, pointer vs value params, nil pointers, `new`.
3. **methods** — value vs pointer receivers, method sets.
4. **interfaces** — implicit satisfaction, empty interface/`any`, type assertions and switches, the nil-interface gotcha, embedding.
5. **errors** — sentinel errors, wrapping with `%w`, `errors.Is/As/Join`, custom error types.
6. **closures** — captured variables, per-iteration loop vars (Go 1.22 semantics).
7. **defer / panic / recover** — ordering, argument evaluation time, recovering.
8. **packages and modules** — exported names, `internal`, imports, `init`.
9. **testing** — table tests, subtests, `t.Helper`, benchmarks with `b.Loop` (1.24).

## Phase 4 — Modern Go features
- [ ] Go 1.21: `min`/`max`/`clear` builtins; `slices` (Sort, SortFunc, BinarySearch, Contains, Index) and `maps` (Keys, Values, Clone) packages.
- [ ] Go 1.22: `for i := range 10`; per-iteration loop variable fix (pair with the closures exercises).
- [ ] Go 1.23: iterators — range-over-func, `iter.Seq`/`Seq2`, `slices.Collect`, `maps.Keys` as iterators.
- [ ] Go 1.24: generic type aliases; `omitzero` JSON tag; `strings.SplitSeq`; `b.Loop`.
- [ ] Go 1.25 (tagged min 1.25): `sync.WaitGroup.Go`; `testing/synctest`.
- [ ] Go 1.26 (tagged min 1.26): `new(expr)`; `errors.AsType`; self-referential generic constraints.
- [ ] `log/slog` structured logging basics.
Expand **generics** beyond 2 exercises: constraints, `cmp.Ordered`, generic structs, generic methods limits.

## Phase 5 — Concurrency and pitfalls
- [ ] Concurrency, from 3 to ~10: goroutines, `sync.WaitGroup`, `sync.Mutex`/`RWMutex`, channels (buffered/unbuffered), `select` with timeouts, `context` cancellation and deadlines, worker pool, fan-in/fan-out, `errgroup`-style patterns using stdlib only, races and `go test -race`.
- [ ] "Pitfalls" track: nil-interface != nil, slice aliasing and append capacity, map iteration order, shadowing with `:=`, defer in loops, closing channels twice, copying a mutex.
- [x] Solutions tree: a verified solution for every exercise, used by CI to prove each exercise is solvable.

## Phase 6 — Quality and release
- [x] CI job that, for each exercise, checks the unsolved version fails and the solution passes. (`TestSolutions`, part of the normal test run; a new exercise without a solution fails the build.)
- [ ] Contributor docs: how to add an exercise (update CONTRIBUTING.md).
- [ ] Tag a release; publish binaries.
- [ ] Decide: open PRs back to upstream (small, per-phase) or keep the fork independent. Upstream is dormant, so expect to maintain the fork.

## Open questions
- Exercise file format and `info.toml` schema details: confirm by reading `golings/exercises/` before Phase 2.
- Where the solutions live so learners don't stumble on them (separate `solutions/` dir vs build-tagged files).
- ~~Rename the fork's module path?~~ Done: module is now `github.com/cleonte/golings`.
