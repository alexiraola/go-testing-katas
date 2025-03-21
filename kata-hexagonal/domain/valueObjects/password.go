package valueObjects

import (
	"errors"
	"hexagonal/domain/common"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Password struct {
	password string
}

func (password Password) String() string {
	return password.password
}

func CreatePassword(password string) (*Password, error) {
	err := ensureIsValidPassword(password)
	if err != nil {
		return nil, err
	}
	return &Password{password: common.Hash(password)}, nil
}

func ensureIsValidPassword(password string) error {
	var errorMessages []string

	if !hasSixCharactersOrMore(password) {
		errorMessages = append(errorMessages, "is too short")
	}
	if !containsNumber(password) {
		errorMessages = append(errorMessages, "must contain a number")
	}
	if !containsLowercase(password) {
		errorMessages = append(errorMessages, "must contain a lowercase")
	}
	if !containsUppercase(password) {
		errorMessages = append(errorMessages, "must contain an uppercase")
	}
	if !containsUnderscore(password) {
		errorMessages = append(errorMessages, "must contain an underscore")
	}

	if len(errorMessages) == 0 {
		return nil
	}

	var stringBuilder strings.Builder
	stringBuilder.WriteString("Password ")
	stringBuilder.WriteString(strings.Join(errorMessages, ", "))

	return errors.New(stringBuilder.String())
}

func hasSixCharactersOrMore(password string) bool {
	return utf8.RuneCountInString(password) >= 6
}

func containsNumber(password string) bool {
	return regexp.MustCompile(`\d`).MatchString(password)
}

func containsLowercase(password string) bool {
	return regexp.MustCompile(`[a-z]`).MatchString(password)
}

func containsUppercase(password string) bool {
	return regexp.MustCompile(`[A-Z]`).MatchString(password)
}

func containsUnderscore(password string) bool {
	return strings.Contains(password, "_")
}
