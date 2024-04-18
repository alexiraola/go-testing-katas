package main

import (
	"regexp"
	"strings"
)

func passwordValidator(password string) bool {
	return hasMoreThanSixCharacters(password) &&
		containsDigit(password) &&
		containsUppercase(password) &&
		containsLowercase(password) &&
		containsUnderscore(password)
}

func hasMoreThanSixCharacters(s string) bool {
	return len(s) >= 6
}

func containsDigit(s string) bool {
	digitRegex := regexp.MustCompile(`\d`)
	return digitRegex.MatchString(s)
}

func containsUppercase(s string) bool {
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	return uppercaseRegex.MatchString(s)
}

func containsLowercase(s string) bool {
	uppercaseRegex := regexp.MustCompile(`[a-z]`)
	return uppercaseRegex.MatchString(s)
}

func containsUnderscore(s string) bool {
	return strings.Contains(s, "_")
}
