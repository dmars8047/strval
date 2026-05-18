package strval

import (
	"strings"
	"testing"
)

func assertOption(t *testing.T, option StringValidationOption, str, strName string, errExpected bool) {
	t.Helper()
	err := option(str, strName)
	if (err != nil) != errExpected {
		t.Errorf("got error %v, wantErr %v", err, errExpected)
	}
	if err != nil && errExpected && strName != "" && !strings.Contains(err.Error(), strName) {
		t.Errorf("error %q does not contain strName %q", err.Error(), strName)
	}
}

func TestMustBeAlphaNumeric(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", true},
		{"only spaces", "   ", "str", true},
		{"only special characters", "!@#$%^&*()_+", "str", true},
		{"only numbers", "1234567890", "str", false},
		{"only uppercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "str", false},
		{"only lowercase", "abcdefghijklmnopqrstuvwxyz", "str", false},
		{"numbers and uppercase", "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ", "str", false},
		{"numbers and lowercase", "1234567890abcdefghijklmnopqrstuvwxyz", "str", false},
		{"uppercase and lowercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "str", false},
		{"numbers, uppercase, and lowercase", "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "str", false},
		{"alphanumeric with special characters", "abc123!@#", "str", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustBeAlphaNumeric(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustContainNumbers(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", true},
		{"only spaces", "   ", "str", true},
		{"only special characters", "!@#$%^&*()_+", "str", true},
		{"only numbers", "1234567890", "str", false},
		{"only uppercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "str", true},
		{"only lowercase", "abcdefghijklmnopqrstuvwxyz", "str", true},
		{"numbers and uppercase", "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ", "str", false},
		{"numbers and lowercase", "1234567890abcdefghijklmnopqrstuvwxyz", "str", false},
		{"uppercase and lowercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "str", true},
		{"all character types", "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "str", false},
		{"alphanumeric with special characters", "abc123!@#", "str", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustContainNumbers(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustContainAtLeastOne(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		query       []rune
		errExpected bool
	}{
		{"empty string with empty query", "", "str", []rune{}, false},
		{"string with empty query", "abc", "str", []rune{}, false},
		{"no matching characters", "abc", "str", []rune{'d', 'e', 'f'}, true},
		{"matching characters", "abc", "str", []rune{'c'}, false},
		{"empty string with query", "", "str", []rune{'c'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustContainAtLeastOne(tt.query), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustContainLowercaseLetter(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", true},
		{"no lowercase letters", "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+", "str", true},
		{"with lowercase letters", "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+", "str", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustContainLowercaseLetter(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustContainUppercaseLetter(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", true},
		{"no uppercase letters", "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+", "str", true},
		{"with uppercase letters", "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+", "str", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustContainUppercaseLetter(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustOnlyContainPrintableCharacters(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", false},
		{"only printable characters", "abc", "str", false},
		{"with non-printable characters", "abc\t\n", "str", true},
		{"emojis", "😀️", "str", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustOnlyContainPrintableCharacters(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustOnlyContainASCIICharacters(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", false},
		{"only ASCII characters", "abcDEF\t\n!@#123", "str", false},
		{"with emojis", "😀️", "str", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustOnlyContainASCIICharacters(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustHaveMinLengthOf(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		minLength   int
		errExpected bool
	}{
		{"meets min length", "abc", "str", 3, false},
		{"below min length", "ab", "str", 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustHaveMinLengthOf(tt.minLength), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustHaveMaxLengthOf(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		maxLength   int
		errExpected bool
	}{
		{"meets max length", "abc", "str", 3, false},
		{"exceeds max length", "abcd", "str", 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustHaveMaxLengthOf(tt.maxLength), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustNotBeEmpty(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", true},
		{"only spaces", "   ", "str", true},
		{"not empty", "abc", "str", false},
		{"only tab", "\t", "str", true},
		{"only newline", "\n", "str", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustNotBeEmpty(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustNotContainAnyOf(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		query       []rune
		errExpected bool
	}{
		{"empty string with empty query", "", "str", []rune{}, false},
		{"string with empty query", "abc", "str", []rune{}, false},
		{"no matching characters", "abc", "str", []rune{'d', 'e', 'f'}, false},
		{"matching characters", "abc", "str", []rune{'c'}, true},
		{"empty string with query", "", "str", []rune{'c'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustNotContainAnyOf(tt.query), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestMustBeValidEmailFormat(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{"empty string", "", "str", true},
		{"no @ symbol", "abc", "str", true},
		{"no username", "@abc.com", "str", true},
		{"no domain", "abc@.com", "str", true},
		{"no domain or extension", "abc@", "str", true},
		{"no domain extension", "abc@abc", "str", true},
		{"valid email", "test@email.com", "str", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOption(t, MustBeValidEmailFormat(), tt.str, tt.strName, tt.errExpected)
		})
	}
}

func TestValidateStringWithName(t *testing.T) {
	result := ValidateStringWithName("", "str", MustNotBeEmpty())
	if result.Valid {
		t.Error("expected invalid result for empty string")
	}
	if len(result.Messages) > 0 && !strings.Contains(result.Messages[0], "str") {
		t.Errorf("message %q does not contain strName", result.Messages[0])
	}
}

func TestValidateString(t *testing.T) {
	result := ValidateString("", MustNotBeEmpty())
	if result.Valid {
		t.Error("expected invalid result for empty string")
	}
	if len(result.Messages) > 0 && !strings.HasPrefix(result.Messages[0], "String") {
		t.Errorf("message %q does not start with 'String'", result.Messages[0])
	}
}
