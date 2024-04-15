package main

import "testing"

func TestEmptyStringShouldReturnZero(t *testing.T) {
	result := stringCalculator("")

	if result != 0 {
		t.Fatalf(`expected %d, got %d`, 0, result)
	}
}

func TestNumericStringShouldReturnNumber(t *testing.T) {
	result := stringCalculator("3")

	if result != 3 {
		t.Fatalf(`expected %d, got %d`, 3, result)
	}
}

func TestNumbersSeperatedByCommaShouldReturnSum(t *testing.T) {
	result := stringCalculator("3,2,1")

	if result != 6 {
		t.Fatalf(`expected %d, got %d`, 6, result)
	}
}

func TestNaNValuesShouldNotCount(t *testing.T) {
	result := stringCalculator("3,a,2a,1")

	if result != 4 {
		t.Fatalf(`expected %d, got %d`, 4, result)
	}
}

func TestCustomSeparator(t *testing.T) {
	result := stringCalculator("//#/3#2")

	if result != 5 {
		t.Fatalf(`expected %d, got %d`, 5, result)
	}
}

func TestCustomSeparatorInvalidNumbers(t *testing.T) {
	result := stringCalculator("//#/3,2")

	if result != 0 {
		t.Fatalf(`expected %d, got %d`, 0, result)
	}
}
