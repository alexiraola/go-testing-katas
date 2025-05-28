package entities_test

import (
	"hexagonal/domain/common"
	"hexagonal/domain/entities"
	"hexagonal/domain/valueObjects"
	"testing"
)

func TestChangesThePasswordWhenADifferentOneIsProvided(t *testing.T) {
	initialPassword, _ := valueObjects.CreatePassword("Safepass123_")
	user := createUser(initialPassword)
	newPassword, _ := valueObjects.CreatePassword("AnotherSafepass123_")

	err := user.ChangePassword(newPassword)
	if err != nil {
		t.Fatalf("Expected password to be %s", newPassword)
	}
}

func TestDoesNotAllowToChangeThePasswordWhenTheGivenOneIsTheSame(t *testing.T) {
	initialPassword, _ := valueObjects.CreatePassword("Safepass123_")
	user := createUser(initialPassword)
	newPassword, _ := valueObjects.CreatePassword("Safepass123_")

	err := user.ChangePassword(newPassword)

	if err == nil || err.Error() != "new password must be different" {
		t.Fatalf("Expected password to be different")
	}
}

func createUser(password *valueObjects.Password) *entities.User {
	id, _ := valueObjects.CreateId(common.GenerateUuid())
	email, _ := valueObjects.CreateEmail("test@example.com")
	return entities.CreateUser(password, id, email)
}
