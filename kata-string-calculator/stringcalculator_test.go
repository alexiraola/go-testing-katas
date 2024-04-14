package main

import "testing"

func TestEmptyStringShouldReturnZero(t *testing.T) {
	result := stringCalculator("")

	if result != 0 {
		t.Fatalf(`expected %d, got %d`, 0, result)
	}
}
