# strval

A Go package for composable string validation. Pass one or more validation options to `ValidateString` or `ValidateStringWithName` and get back a result with all failure messages collected.

## Installation

```bash
go get github.com/dmars8047/strval
```

## Usage

```go
import "github.com/dmars8047/strval"

// Basic validation — field name defaults to "String"
result := strval.ValidateString(input,
    strval.MustNotBeEmpty(),
    strval.MustHaveMinLengthOf(8),
    strval.MustContainUppercaseLetter(),
    strval.MustContainNumbers(),
)

if !result.Valid {
    for _, msg := range result.Messages {
        fmt.Println(msg)
    }
}
```

Use `ValidateStringWithName` to include a field name in error messages:

```go
result := strval.ValidateStringWithName(req.Email, "Email",
    strval.MustNotBeEmpty(),
    strval.MustBeValidEmailFormat(),
)
// on failure: "Email must be a valid email format"
```

All failing options are collected — validation does not short-circuit on the first error.

## Available Options

| Option | Description |
|---|---|
| `MustNotBeEmpty()` | Fails if the string is empty or contains only whitespace |
| `MustHaveMinLengthOf(n)` | Fails if the string is shorter than `n` bytes |
| `MustHaveMaxLengthOf(n)` | Fails if the string is longer than `n` bytes |
| `MustBeAlphaNumeric()` | Fails if the string contains any non-alphanumeric characters |
| `MustContainNumbers()` | Fails if the string contains no digits |
| `MustContainUppercaseLetter()` | Fails if the string contains no uppercase letters |
| `MustContainLowercaseLetter()` | Fails if the string contains no lowercase letters |
| `MustContainAtLeastOne(chars []rune)` | Fails if none of the provided runes appear in the string |
| `MustNotContainAnyOf(chars []rune)` | Fails if any of the provided runes appear in the string |
| `MustOnlyContainPrintableCharacters()` | Fails if the string contains non-printable characters (tabs, newlines, null bytes, etc.) |
| `MustOnlyContainASCIICharacters()` | Fails if the string contains any non-ASCII characters |
| `MustBeValidEmailFormat()` | Fails if the string does not match a standard email format |
