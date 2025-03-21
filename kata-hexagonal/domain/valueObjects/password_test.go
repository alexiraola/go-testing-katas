package valueObjects_test

import (
	"hexagonal/domain/valueObjects"
	"regexp"
	"testing"
	"unicode/utf8"
)

func TestCreateAPasswordFromAValidString(t *testing.T) {
	validPassword := "SecurePass123_"
	_, err := valueObjects.CreatePassword(validPassword)
	if err != nil {
		t.Fatalf("Unexpected error for valid password %s", validPassword)
	}
}

func TestFailsWhenThePasswordIsTooShort(t *testing.T) {
	shortPassword := "1aaA_"
	_, err := valueObjects.CreatePassword(shortPassword)
	if err == nil || err.Error() != "Password is too short" {
		t.Fatalf("Expected error for short password %s", shortPassword)
	}
}

func TestFailsWhenThePasswordIsMissingANumber(t *testing.T) {
	password := "aaaaaA_"
	_, err := valueObjects.CreatePassword(password)
	if err == nil || err.Error() != "Password must contain a number" {
		t.Fatalf("Expected error for missing number password %s", password)
	}
}

func TestFailsWhenThePasswordIsMissingALowercase(t *testing.T) {
	password := "A1234_"
	_, err := valueObjects.CreatePassword(password)
	if err == nil || err.Error() != "Password must contain a lowercase" {
		t.Fatalf("Expected error for missing lowercase password %s", password)
	}
}

func TestFailsWhenThePasswordIsMissingAnUppercase(t *testing.T) {
	password := "a1234_"
	_, err := valueObjects.CreatePassword(password)
	if err == nil || err.Error() != "Password must contain an uppercase" {
		t.Fatalf("Expected error for missing uppercase password %s", password)
	}
}

func TestFailsWhenThePasswordIsMissingAnUnderscore(t *testing.T) {
	password := "aA12345"
	_, err := valueObjects.CreatePassword(password)
	if err == nil || err.Error() != "Password must contain an underscore" {
		t.Fatalf("Expected error for missing underscore password %s", password)
	}
}

func TestFailsWhenThePasswordIsMissingSeveralRequirements(t *testing.T) {
	password := "abc"
	_, err := valueObjects.CreatePassword(password)
	if err == nil || err.Error() != "Password is too short, must contain a number, must contain an uppercase, must contain an underscore" {
		t.Fatalf("Expected error for missing several requirements password %s", password)
	}
}

func TestEnsurePasswordIsHashed(t *testing.T) {
	plaintext := "SecurePass123_"
	password, _ := valueObjects.CreatePassword(plaintext)
	hashedValue := password.String()

	if hashedValue == plaintext {
		t.Fatalf("Expected password to be hashed")
	}
	if utf8.RuneCountInString(hashedValue) != 64 {
		t.Fatalf("Expected hash to be 64 chars length")
	}
	if !isHashed(hashedValue) {
		t.Fatalf("Expected hash to be a valid hash")
	}
}

func TestMatchesWhenSomeGivenPasswordsAreTheSame(t *testing.T) {
	plaintext := "SecurePass123_"
	aPassword, _ := valueObjects.CreatePassword(plaintext)
	anotherPassword, _ := valueObjects.CreatePassword(plaintext)

	if *aPassword != *anotherPassword {
		t.Fatalf("Expected %s to be equal to %s", aPassword, anotherPassword)
	}
}

func TestDoesNotMatchWhenSomeGivenPasswordAreDifferent(t *testing.T) {
	aPassword, _ := valueObjects.CreatePassword("SecurePass123_")
	anotherPassword, _ := valueObjects.CreatePassword("DifferentPass123_")

	if *aPassword == *anotherPassword {
		t.Fatalf("Expected %s to be different to %s", aPassword, anotherPassword)
	}
}

func isHashed(hash string) bool {
	return regexp.MustCompile(`[a-f-F0-9]{64}`).MatchString(hash)
}
