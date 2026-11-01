// Package golings holds the exercises that ship inside the golings binary.
//
// The CLI itself lives in the golings/ folder. This package sits at the root
// of the module because go:embed can't reach into parent folders.
package golings

import "embed"

// FS contains the exercise manifest and every exercise, as they are in the
// repository. `golings init` writes it to disk and `golings reset` restores
// single exercises from it. Note that building from a working copy embeds
// whatever the exercise files look like at that moment.
//
//go:embed info.toml exercises
var FS embed.FS

// Solutions contains a working version of each exercise, laid out like
// exercises/ (solutions/<topic>/<name>/<file>). It is only ever read by
// `golings solution`; `golings init` must not write it, which is why it is
// separate from FS.
//
//go:embed solutions
var Solutions embed.FS
