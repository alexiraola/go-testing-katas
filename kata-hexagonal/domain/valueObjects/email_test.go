package valueObjects

import "testing"

func TestCreatesAnEmailWithValidString(t *testing.T) {
	validEmail := "test@example.com"
	_, err := CreateEmail(validEmail)
	if err != nil {
		t.Fatalf("Unexpected error for valid email %s", validEmail)
	}
}

func TestDoesNotAllowCreatingAnEmailForAGivenIncorrectlyFormattedAddress(t *testing.T) {
	invalidEmail := "invalid"
	_, err := CreateEmail(invalidEmail)

	if err == nil {
		t.Fatalf("Expected error for invalid email %s", invalidEmail)
	}
}

func TestConsidersTwoEmailsWithTheSameAddressAsEqual(t *testing.T) {
	aEmail, _ := CreateEmail("example@example.com")
	otherEmail, _ := CreateEmail("example@example.com")

	if !aEmail.Equals(otherEmail) {
		t.Errorf("Expected %s and %s to be equal", aEmail.String(), otherEmail.String())
	}
}

func TestDifferentiatesBetweenTwoEmailsWithDifferentAddress(t *testing.T) {
	aEmail, _ := CreateEmail("example@example.com")
	otherEmail, _ := CreateEmail("anotherexample@example.com")

	if aEmail.Equals(otherEmail) {
		t.Errorf("Expected %s and %s not to be equal", aEmail.String(), otherEmail.String())
	}
}
