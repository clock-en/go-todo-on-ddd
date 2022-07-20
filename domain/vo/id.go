package vo

import (
	"fmt"
)

type IId interface {
	Value() int
}

type Id struct {
	IId
	value int
}

func (i Id) Value() int {
	return i.value
}

func NewId(value int) (*Id, error) {
	if isInvalidId(value) {
		return nil, fmt.Errorf("ID must be greater than or equal to 1 character")
	}
	return &Id{value: value}, nil
}

func isInvalidId(value int) bool {
	return 0 < value
}
