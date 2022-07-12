package vo

import (
	"fmt"
	"regexp"
)

// UserName represents value obejct user name
type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if isInvalidPatternUserName(value) {
		return nil, fmt.Errorf("user name must be a alphanumeric")
	}
	if isInvalidLengthUserName(value) {
		return nil, fmt.Errorf("user name must be no longer than 50 characters")
	}
	return &UserName{value: value}, nil
}

func (n UserName) Value() string {
	return n.value
}

func isInvalidPatternUserName(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(value)
}

func isInvalidLengthUserName(value string) bool {
	return 50 < len(value)
}
