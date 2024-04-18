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
	beginOfConfig := "//"
	endOfConfig := "/"
	if strings.HasPrefix(text, beginOfConfig) {
		startIndex := len(beginOfConfig)
		endIndex := strings.LastIndex(text, endOfConfig)
		separator := text[startIndex:endIndex]
		numbers := text[endIndex+1:]

		return separator, numbers
	}
	return ",", text
}
