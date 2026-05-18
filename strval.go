package strval

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var reEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type StringValidationOption func(string, string) error

func MustHaveMinLengthOf(minLength int) StringValidationOption {
	return func(str, strName string) error {
		if utf8.RuneCountInString(str) < minLength {
			return fmt.Errorf("%s must have a minimum length of %d", strName, minLength)
		}
		return nil
	}
}

func MustHaveMaxLengthOf(maxLength int) StringValidationOption {
	return func(str, strName string) error {
		if utf8.RuneCountInString(str) > maxLength {
			return fmt.Errorf("%s must have a maximum length of %d", strName, maxLength)
		}
		return nil
	}
}

func MustNotBeEmpty() StringValidationOption {
	return func(str, strName string) error {
		if strings.TrimSpace(str) == "" {
			return fmt.Errorf("%s must not be empty", strName)
		}
		return nil
	}
}

func MustBeAlphaNumeric() StringValidationOption {
	return func(str, strName string) error {
		if !isAlphaNumeric(str) {
			return fmt.Errorf("%s must be alphanumeric", strName)
		}
		return nil
	}
}

func MustContainNumbers() StringValidationOption {
	return func(str, strName string) error {
		for _, r := range str {
			if unicode.IsDigit(r) {
				return nil
			}
		}
		return fmt.Errorf("%s must contain numbers", strName)
	}
}

func MustContainAtLeastOne(characters []rune) StringValidationOption {
	set := make(map[rune]struct{}, len(characters))
	for _, r := range characters {
		set[r] = struct{}{}
	}
	return func(str, strName string) error {
		if len(set) == 0 {
			return nil
		}
		for _, r := range str {
			if _, ok := set[r]; ok {
				return nil
			}
		}
		return fmt.Errorf("%s must contain at least one of the following characters: %s", strName, string(characters))
	}
}

func MustNotContainAnyOf(disallowedCharacters []rune) StringValidationOption {
	set := make(map[rune]struct{}, len(disallowedCharacters))
	for _, r := range disallowedCharacters {
		set[r] = struct{}{}
	}
	return func(str, strName string) error {
		for _, r := range str {
			if _, ok := set[r]; ok {
				return fmt.Errorf("%s must not contain any of the following characters: %s", strName, string(disallowedCharacters))
			}
		}
		return nil
	}
}

func MustContainUppercaseLetter() StringValidationOption {
	return func(str, strName string) error {
		for _, r := range str {
			if unicode.IsUpper(r) {
				return nil
			}
		}
		return fmt.Errorf("%s must contain at least one uppercase letter", strName)
	}
}

func MustContainLowercaseLetter() StringValidationOption {
	return func(str, strName string) error {
		for _, r := range str {
			if unicode.IsLower(r) {
				return nil
			}
		}
		return fmt.Errorf("%s must contain at least one lowercase letter", strName)
	}
}

func MustOnlyContainPrintableCharacters() StringValidationOption {
	return func(str, strName string) error {
		for _, r := range str {
			if !unicode.IsPrint(r) {
				return fmt.Errorf("%s must only contain printable characters", strName)
			}
		}
		return nil
	}
}

func MustOnlyContainASCIICharacters() StringValidationOption {
	return func(str, strName string) error {
		for _, r := range str {
			if r > unicode.MaxASCII {
				return fmt.Errorf("%s must only contain ASCII characters", strName)
			}
		}
		return nil
	}
}

func MustBeValidEmailFormat() StringValidationOption {
	return func(str, strName string) error {
		if !reEmail.MatchString(str) {
			return fmt.Errorf("%s must be a valid email format", strName)
		}
		return nil
	}
}

type StringValidationResult struct {
	Valid    bool
	Messages []string
}

func ValidateStringWithName(str, strName string, options ...StringValidationOption) StringValidationResult {
	var messages []string
	isValid := true
	for _, option := range options {
		if err := option(str, strName); err != nil {
			messages = append(messages, err.Error())
			isValid = false
		}
	}
	return StringValidationResult{Valid: isValid, Messages: messages}
}

func ValidateString(str string, options ...StringValidationOption) StringValidationResult {
	return ValidateStringWithName(str, "String", options...)
}

func isAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}
