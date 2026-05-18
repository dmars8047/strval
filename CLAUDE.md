# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test ./...

# Run a single test
go test -run TestMustNotBeEmpty

# Build
go build ./...
```

## Architecture

`strval` is a single-package Go library (`github.com/dmars8047/strval`) for composable string validation. There are no external dependencies.

**Core type:** `StringValidationOption` is `func(str, strName string) error`. Every validation rule is a factory function (e.g. `MustNotBeEmpty()`) that returns one of these.

**Entry points:**
- `ValidateStringWithName(str, strName string, options ...StringValidationOption) StringValidationResult` — runs all options and aggregates failures. `strName` appears in every error message.
- `ValidateString(str string, options ...StringValidationOption) StringValidationResult` — same, but uses `"String"` as the default name.

**Pattern:** public factory functions call private helpers (`isEmpty`, `containsNonPrintableCharacters`, etc.) that do the actual character-level work. New options should follow this same split.

**`internal/test/`** contains a placeholder file and is not used anywhere meaningful.

## Conventions

- Error messages must include `strName` so callers get field-specific feedback.
- Passing an empty `[]rune{}` to `MustContainAtLeastOne` or `MustNotContainAnyOf` is a no-op (always passes) — preserve this behavior.
- `MustNotBeEmpty` treats whitespace-only strings (spaces, tabs, newlines) as empty.
- Length checks (`MustHaveMinLengthOf` / `MustHaveMaxLengthOf`) operate on byte length via `len()`, not rune count.
