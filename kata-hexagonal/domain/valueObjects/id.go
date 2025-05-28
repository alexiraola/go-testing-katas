package valueObjects

import (
	"errors"
	"hexagonal/domain/common"
	"regexp"
)

type Id struct {
	id string
}

func (id *Id) String() string {
	return id.id
}

func (id *Id) Equals(other *Id) bool {
	if other == nil {
		return false
	}
	return id.id == other.id
}

func GenerateUniqueIdentifier() *Id {
	// create uuid
	return &Id{id: common.GenerateUuid()}
}

func CreateId(id string) (*Id, error) {
	err := ensureIsValidIdentifier(id)
	if err != nil {
		return nil, err
	}
	return &Id{id}, nil
}

func ensureIsValidIdentifier(id string) error {
	const validUuidRegex = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`

	if !regexp.MustCompile(validUuidRegex).MatchString(id) {
		return errors.New("invalid Id format")
	}
	return nil
}
