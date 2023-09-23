package strval

import (
	"fmt"
	"regexp"
	"unicode"
)

// Strings
type StringValidationOption func(string, string) error

// This option will validate that the string is at least minLength characters long
func MustHaveMinLengthOf(minLength int) StringValidationOption {
	return func(str, strName string) error {
		if !isWithMinLength(str, minLength) {
			return fmt.Errorf("%s must have a minimum length of %d", strName, minLength)
		}

		return nil
	}
}

// This option will validate that the string is at most maxLength characters long
func MustHaveMaxLengthOf(maxLength int) StringValidationOption {
	return func(str, strName string) error {
		if !isWithMaxLength(str, maxLength) {
			return fmt.Errorf("%s must have a maximum length of %d", strName, maxLength)
		}

		return nil
	}
}

// This option will validate that the string is not empty or all whitespace
func MustNotBeEmpty() StringValidationOption {
	return func(str, strName string) error {
		if isEmpty(str) {
			return fmt.Errorf("%s must not be empty", strName)
		}

		return nil
	}
}

// This option will validate that the string is alphanumeric
func MustBeAlphaNumeric() StringValidationOption {
	return func(str, strName string) error {
		if !isAlphaNumeric(str) {
			return fmt.Errorf("%s must be alphanumeric", strName)
		}

		return nil
	}
}

// This option will validate that the string contains at least one number
func MustContainNumbers() StringValidationOption {
	return func(str, strName string) error {
		if !containsNumbers(str) {
			return fmt.Errorf("%s must contain numbers", strName)
		}

		return nil
	}
}

// This option will validate that the string contains at least one of the allowable characters
func MustContainAtLeastOne(characters []rune) StringValidationOption {
	return func(str, strName string) error {
		if !containsAtLeastOneOf(str, characters) {
			return fmt.Errorf("%s must contain at least one of the following characters: %s", strName, string(characters))
		}

		return nil
	}
}

// This option will validate that the string does not contain any of the disallowed characters
func MustNotContainAnyOf(disallowedCharacters []rune) StringValidationOption {
	return func(str, strName string) error {
		if containsAny(str, disallowedCharacters) {
			return fmt.Errorf("%s must not contain any of the following characters: %s", strName, string(disallowedCharacters))
		}

		return nil
	}
}

// This option will validate that the string contains at least one uppercase letter
func MustContainUppercaseLetter() StringValidationOption {
	return func(str, strName string) error {
		if !containsUppercaseLetter(str) {
			return fmt.Errorf("%s must contain at least one uppercase letter", strName)
		}

		return nil
	}
}

// This option will validate that the string contains at least one lowercase letter
func MustContainLowercaseLetter() StringValidationOption {
	return func(str, strName string) error {
		if !containsLowercaseLetter(str) {
			return fmt.Errorf("%s must contain at least one lowercase letter", strName)
		}

		return nil
	}
}

// This option will validate that the string only contains printable characters
func MustOnlyContainPrintableCharacters() StringValidationOption {
	return func(str, strName string) error {
		if containsNonPrintableCharacters(str) {
			return fmt.Errorf("%s must only contain printable characters", strName)
		}

		return nil
	}
}

// This option will validate that the string only contains ASCII characters
func MustOnlyContainASCIICharacters() StringValidationOption {
	return func(str, strName string) error {
		if containsNonASCIICharacters(str) {
			return fmt.Errorf("%s must only contain ASCII characters", strName)
		}

		return nil
	}
}

// This option will validate that the string is a valid email format
func MustBeValidEmailFormat() StringValidationOption {
	return func(str, strName string) error {
		if !isValidEmailFormat(str) {
			return fmt.Errorf("%s must be a valid email format", strName)
		}

		return nil
	}
}

// This represents the result of a string validation operation
type StringValidationResult struct {
	Valid    bool
	Messages []string
}

// ValidateStringWithName validates a string against the provided options
// str: The string to validate
// strName: The name of the string to validate (used in error messages)
// options: The options to validate the string against
// Returns a StringValidationResult
func ValidateStringWithName(str, strName string, options ...StringValidationOption) StringValidationResult {
	var messages []string
	var isValid = true

	for _, option := range options {
		if err := option(str, strName); err != nil {
			messages = append(messages, err.Error())
			isValid = false
		}
	}

	return StringValidationResult{
		Valid:    isValid,
		Messages: messages,
	}
}

// ValidateString validates a string against the provided options
// str: The string to validate
// options: The options to validate the string against
// Returns a StringValidationResult
func ValidateString(str string, options ...StringValidationOption) StringValidationResult {
	return ValidateStringWithName(str, "String", options...)
}

// containsNumbers checks if a string contains a number
func containsNumbers(password string) bool {
	if matched, _ := regexp.MatchString(`[0-9]`, password); matched {
		return true
	}
	return false
}

// containsAtLeastOneOf checks if a string at least on of the characters in the provided string
func containsAtLeastOneOf(value string, query []rune) bool {
	if len(query) == 0 {
		return true
	}

	for _, char := range value {
		for _, queryChar := range query {
			if char == queryChar {
				return true
			}
		}
	}

	return false
}

// containsAny checks if a string containsAny none of the characters in the provided string
func containsAny(value string, query []rune) bool {
	if len(query) == 0 {
		return false
	}

	for _, char := range value {
		for _, queryChar := range query {
			if char == queryChar {
				return true
			}
		}
	}

	return false
}

// containsUppercaseLetter checks if a string contains an uppercase letter
func containsUppercaseLetter(password string) bool {
	if matched, _ := regexp.MatchString(`[A-Z]`, password); matched {
		return true
	}
	return false
}

// containsLowercaseLetter checks if a string contains a lowercase letter
func containsLowercaseLetter(password string) bool {
	if matched, _ := regexp.MatchString(`[a-z]`, password); matched {
		return true
	}
	return false
}

// containsNonPrintableCharacters checks if a string contains a non-printable character
func containsNonPrintableCharacters(password string) bool {
	for _, char := range password {
		if !unicode.IsPrint(char) {
			return true
		}
	}
	return false
}

// containsNonASCIICharacters checks if a string contains a non-ASCII character
func containsNonASCIICharacters(password string) bool {
	for _, char := range password {
		if char > unicode.MaxASCII {
			return true
		}
	}
	return false
}

// IsAlphaNumeric checks if a string contains only alphanumeric characters
func isAlphaNumeric(str string) bool {
	// Regular expression pattern for alphanumeric validation
	pattern := "^[a-zA-Z0-9]+$"

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Check if the string matches the pattern
	return regex.MatchString(str)
}

// isValidEmailFormat checks if a string contains only alphanumeric characters and spaces
func isValidEmailFormat(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	return regex.MatchString(email)
}

// isEmpty checks to see if a string is empty or all whitespace
func isEmpty(str string) bool {
	// Regular expression pattern for email validation
	pattern := `^\s*$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	return regex.MatchString(str)
}

// isWithMinLength checks to see if a string is at least minLength characters long
func isWithMinLength(str string, minLength int) bool {
	return len(str) >= minLength
}

// isWithMaxLength checks to see if a string is at most maxLength characters long
func isWithMaxLength(str string, maxLength int) bool {
	return len(str) <= maxLength
}
