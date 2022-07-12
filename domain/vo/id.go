package vo

import (
	"fmt"
)

// Password represents value obejct id (Table PK)
type Id struct {
	value int
}

func NewId(value int) (*Id, error) {
	if isInvalidId(value) {
		return nil, fmt.Errorf("ID must be greater than or equal to 1 character")
	}
	return &Id{value: value}, nil
}

func (i Id) Value() int {
	return i.value
}

func isInvalidId(value int) bool {
	return 0 < value
}
