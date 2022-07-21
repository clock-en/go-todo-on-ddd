package vo

import (
	"fmt"

	"github.com/google/uuid"
)

type ID struct {
	value string
}

func (i ID) Value() string {
	return i.value
}

func NewID(value string) (*ID, error) {
	if isInvalidUUID(value) {
		return nil, fmt.Errorf("wrong ID")
	}
	return &ID{value: value}, nil
}

func isInvalidUUID(value string) bool {
	valueID, err := uuid.Parse(value)
	if err != nil || isInvalidUUIDVersion(valueID) {
		return true
	}
	return false
}

func isInvalidUUIDVersion(valueID uuid.UUID) bool {
	id, err := uuid.NewRandom()
	if err != nil {
		return true
	}
	return valueID.Version() != id.Version()
}
