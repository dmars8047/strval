package strval

import (
	"strings"
	"testing"
)

// Tests StringValidationOption MustBeAlphaNumeric()
func TestMustBeAlphaNumeric(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only spaces",
			str:         "   ",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only special characters",
			str:         "!@#$%^&*()_+",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only numbers",
			str:         "1234567890",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only uppercase letters",
			str:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only lowercase letters",
			str:         "abcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only numbers and uppercase letters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only numbers and lowercase letters",
			str:         "1234567890abcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only uppercase and lowercase letters",
			str:         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only numbers, uppercase and lowercase letters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with numbers, uppercase and lowercase letters and special characters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with numbers, uppercase and lowercase letters and special characters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+",
			strName:     "str",
			errExpected: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustBeAlphaNumeric()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustBeAlphaNumeric() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustBeAlphaNumeric() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustContainNumbers()
func TestMustContainNumbers(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only spaces",
			str:         "   ",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only special characters",
			str:         "!@#$%^&*()_+",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only numbers",
			str:         "1234567890",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only uppercase letters",
			str:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only lowercase letters",
			str:         "abcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only numbers and uppercase letters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only numbers and lowercase letters",
			str:         "1234567890abcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only uppercase and lowercase letters",
			str:         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only numbers, uppercase and lowercase letters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with numbers, uppercase and lowercase letters and special characters",
			str:         "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+",
			strName:     "str",
			errExpected: false,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustContainNumbers()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustContainNumbers() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustContainNumbers() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustContainAtLeastOne()
func TestMustContainAtLeastOne(t *testing.T) {
	//  Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		query       []rune
		errExpected bool
	}{
		{
			name:        "empty string with empty query",
			str:         "",
			strName:     "str",
			query:       []rune{},
			errExpected: false,
		},
		{
			name:        "string with empty query",
			str:         "abc",
			query:       []rune{},
			errExpected: false,
		},
		{
			name:        "string without matching query characters",
			str:         "abc",
			query:       []rune{'d', 'e', 'f'},
			errExpected: true,
		},
		{
			name:        "string with matching query characters",
			str:         "abc",
			query:       []rune{'c'},
			errExpected: false,
		},
		{
			name: "empty string with query characters",
			str:  "",
			query: []rune{
				'c',
			},
			errExpected: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustContainAtLeastOne(tt.query)(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustContainAtLeastOne() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustContainAtLeastOne() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustContainLowercaseLetter()
func TestMustContainLowercaseLetter(t *testing.T) {
	//  Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with no lowercase letters",
			str:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with lowercase letters",
			str:         "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+",
			strName:     "str",
			errExpected: false,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustContainLowercaseLetter()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustContainLowercaseLetter() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustContainLowercaseLetter() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustContainUppercaseLetter()
func TestMustContainUppercaseLetter(t *testing.T) {
	//  Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with no uppercase letters",
			str:         "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with uppercase letters",
			str:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+",
			strName:     "str",
			errExpected: false,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustContainUppercaseLetter()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustContainUppercaseLetter() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustContainUppercaseLetter() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustMustOnlyContainPrintableCharacters()
func TestMustOnlyContainPrintableCharacters(t *testing.T) {
	//  Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only printable characters",
			str:         "abc",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with non printable characters",
			str:         "abc\t\n",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with emojis",
			str:         "😀️",
			strName:     "str",
			errExpected: false,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustOnlyContainPrintableCharacters()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustNotContainNonPrintableCharacters() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustNotContainNonPrintableCharacters() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Test StringValidationOption MustOnlyContainASCIICharacters()
func TestMustOnlyContainASCIICharacters(t *testing.T) {
	//  Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with only ASCII characters",
			str:         "abcDEF\t\n!@#123",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with emojis",
			str:         "😀️",
			strName:     "str",
			errExpected: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustOnlyContainASCIICharacters()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustOnlyContainASCIICharacters() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustOnlyContainASCIICharacters() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustHaveMinLengthOf()
func TestMustHaveMinLengthOf(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		minLength   int
		errExpected bool
	}{
		{
			name:        "string that meets min length",
			str:         "abc",
			strName:     "str",
			minLength:   3,
			errExpected: false,
		},
		{
			name:        "string that does not meet min length",
			str:         "ab",
			strName:     "str",
			minLength:   3,
			errExpected: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustHaveMinLengthOf(tt.minLength)(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustHaveMinLengthOf() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustHaveMinLengthOf() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustHaveMaxLengthOf()
func TestMustHaveMaxLengthOf(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		maxLength   int
		errExpected bool
	}{
		{
			name:        "string that meets max length",
			str:         "abc",
			strName:     "str",
			maxLength:   3,
			errExpected: false,
		},
		{
			name:        "string that does not meet max length",
			str:         "abcd",
			strName:     "str",
			maxLength:   3,
			errExpected: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustHaveMaxLengthOf(tt.maxLength)(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustHaveMaxLengthOf() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustHaveMaxLengthOf() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustNotBeEmpty()
func TestMustNotBeEmpty(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with only spaces",
			str:         "   ",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string that is not empty",
			str:         "abc",
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with tab",
			str:         "\t",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with newline",
			str:         "\n",
			strName:     "str",
			errExpected: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustNotBeEmpty()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustNotBeEmpty() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustNotBeEmpty() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustNotContain()
func TestMustNotContain(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		query       []rune
		errExpected bool
	}{
		{
			name:        "empty string with empty query",
			str:         "",
			strName:     "str",
			query:       []rune{},
			errExpected: false,
		},
		{
			name:        "string with empty query",
			str:         "abc",
			query:       []rune{},
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string without matching query characters",
			str:         "abc",
			query:       []rune{'d', 'e', 'f'},
			strName:     "str",
			errExpected: false,
		},
		{
			name:        "string with matching query characters",
			str:         "abc",
			query:       []rune{'c'},
			strName:     "str",
			errExpected: true,
		},
		{
			name:    "empty string with query characters",
			str:     "",
			strName: "str",
			query: []rune{
				'c',
			},
			errExpected: false,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustNotContainAnyOf(tt.query)(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustNotContain() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustNotContain() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Tests StringValidationOption MustBeValidEmailFormat()
func TestMustBeValidEmailFormat(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		errExpected bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string without @",
			str:         "abc",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string without username",
			str:         "@abc.com",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string without domain",
			str:         "abc@.com",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string without domain or extension",
			str:         "abc@",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string without domain extension",
			str:         "abc@abc",
			strName:     "str",
			errExpected: true,
		},
		{
			name:        "string with valid email format",
			str:         "test@email.com",
			strName:     "str",
			errExpected: false,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustBeValidEmailFormat()(tt.str, tt.strName)

			errFound := err != nil

			if (errFound && !tt.errExpected) || (!errFound && tt.errExpected) {
				t.Errorf("MustBeValidEmailFormat() error = %v, wantErr %v", err, tt.errExpected)
			}

			// Make sure the strName is in the error message
			if errFound && tt.errExpected && !strings.Contains(err.Error(), tt.strName) {
				t.Errorf("MustBeValidEmailFormat() strName error = %v, expected to contain strName %v", err, tt.strName)
			}
		})
	}
}

// Test ValidateStringWithName(str, strName string, options ...StringValidationOption) StringValidationResult
func TestValidateStringWithName(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		expectValid bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			expectValid: false,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateStringWithName(tt.str, tt.strName, MustNotBeEmpty())

			if result.Valid != tt.expectValid {
				t.Errorf("ValidateStringWithName() IsValid = %v, want %v", result.Valid, tt.expectValid)
			}

			// Make sure the strName is in the first message
			if len(result.Messages) > 0 && !strings.Contains(result.Messages[0], tt.strName) {
				t.Errorf("ValidateStringWithName() strName error = %v, expected to contain strName %v", result.Messages[0], tt.strName)
			}
		})
	}
}

// Test ValidateString(str, strName string, options ...StringValidationOption) StringValidationResult
func TestValidateString(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		str         string
		strName     string
		expectValid bool
	}{
		{
			name:        "empty string",
			str:         "",
			strName:     "str",
			expectValid: false,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateString(tt.str, MustNotBeEmpty())

			if result.Valid != tt.expectValid {
				t.Errorf("ValidateString() IsValid = %v, want %v", result.Valid, tt.expectValid)
			}

			// Make sure the first message starts with 'String'
			if len(result.Messages) > 0 && !strings.HasPrefix(result.Messages[0], "String") {
				t.Errorf("ValidateString() strName error = %v, expected to start with 'String'", result.Messages[0])
			}
		})
	}
}
