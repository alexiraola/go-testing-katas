package main

import (
	"strconv"
	"strings"
)

func stringCalculator(text string) int {
	if text == "" {
		return 0
	}
	separator, textWithoutPrefix := findSeparator(text)
	numbers := strings.Split(textWithoutPrefix, separator)
	sum := 0

	for _, n := range numbers {
		number, error := strconv.Atoi(n)

		if error == nil {
			sum += number
		}
	}

	return sum
}

func findSeparator(text string) (string, string) {
	if strings.HasPrefix(text, "//") {
		separator := text[2:3]
		numbers := text[4:]

		return separator, numbers
	}
	return ",", text
}
