package valueObjects_test

import (
	"hexagonal/domain/common"
	"hexagonal/domain/valueObjects"
	"regexp"
	"testing"
)

func TestGeneratesAValidIdentifier(t *testing.T) {
	id := valueObjects.GenerateUniqueIdentifier()

	if !isValidUuid(id.String()) {
		t.Fatalf("Expected %s to be a valid UUID", id)
	}
}

func TestCreatesAnIDFromAValidIdentifier(t *testing.T) {
	validId := common.GenerateUuid()
	_, err := valueObjects.CreateId(validId)
	if err != nil {
		t.Fatalf("Unexpected error for valid id %s", validId)
	}
}

func TestDoesNotAllowToCreateAnIDFromAnInvalidIdentifier(t *testing.T) {
	invalidId := "invalid-id"

	_, err := valueObjects.CreateId(invalidId)
	if err == nil || err.Error() != "invalid Id format" {
		t.Fatalf("Expected error for invalid id %s", invalidId)
	}
}

func TestIdentifisTwoIdenticalIdentifiersAsEqual(t *testing.T) {
	validId := common.GenerateUuid()
	id1, _ := valueObjects.CreateId(validId)
	id2, _ := valueObjects.CreateId(validId)

	if !id1.Equals(id2) {
		t.Fatalf("Expected %s to be equal to %s", id1, id2)
	}
}

func TestIdentifisTwoDifferentIdentifiersAsNotEqual(t *testing.T) {
	id1, _ := valueObjects.CreateId(common.GenerateUuid())
	id2, _ := valueObjects.CreateId(common.GenerateUuid())

	if id1.Equals(id2) {
		t.Fatalf("Expected %s to be different to %s", id1, id2)
	}
}

func isValidUuid(id string) bool {
	return regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`).MatchString(id)
}
