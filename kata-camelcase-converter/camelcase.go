package main

import (
	"strings"
	"unicode"
)

func camelcase(text string) string {
	var converted []string
	for _, word := range getWords(text) {
		converted = append(converted, convertWord(word))
	}

	return strings.Join(converted, "")
}

func getWords(text string) []string {
	return strings.FieldsFunc(text, func(c rune) bool {
		return unicode.IsSpace(c) || c == '-' || c == '_'
	})
}

func convertWord(word string) string {
	if unicode.IsLower(rune(word[0])) {
		return string(unicode.ToUpper(rune(word[0]))) + word[1:]
	}
	return word
}
