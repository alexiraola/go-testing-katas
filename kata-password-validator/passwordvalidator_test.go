package main

import "testing"

func TestShouldReturnTrueWhenAllRequirementsAreMet(t *testing.T) {
	result := passwordValidator("1234abcdABCD_")

	if result != true {
		t.Fatalf(`expected %t, got %t`, true, result)
	}
}

func TestShouldReturnFalseWhenLengthIsLessThan6Chars(t *testing.T) {
	result := passwordValidator("1aA_")

	if result != false {
		t.Fatalf(`expected %t, got %t`, false, result)
	}
}

func TestShouldReturnFalseWhenItContainsNoDigit(t *testing.T) {
	result := passwordValidator("abcdABCD_")

	if result != false {
		t.Fatalf(`expected %t, got %t`, false, result)
	}
}

func TestShouldReturnFalseWhenItContainsNoUppercase(t *testing.T) {
	result := passwordValidator("abcd1234_")

	if result != false {
		t.Fatalf(`expected %t, got %t`, false, result)
	}
}

func TestShouldReturnFalseWhenItContainsNoLowercase(t *testing.T) {
	result := passwordValidator("ABCD1234_")

	if result != false {
		t.Fatalf(`expected %t, got %t`, false, result)
	}
}

func TestShouldReturnFalseWhenItContainsNoUnderscore(t *testing.T) {
	result := passwordValidator("abcdABCD1234")

	if result != false {
		t.Fatalf(`expected %t, got %t`, false, result)
	}
}
