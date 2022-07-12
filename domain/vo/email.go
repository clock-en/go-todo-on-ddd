package vo

import (
	"fmt"
	"regexp"
)

// Email represents value obejct e-mail address
type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if isInvalidPatternEmail(value) {
		return nil, fmt.Errorf("email is not in the correct format")
	}
	if isInvalidLengthEmail(value) {
		return nil, fmt.Errorf("email must be no longer than 254 characters")
	}
	return &Email{value: value}, nil
}

func (e Email) Value() string {
	return e.value
}

func isInvalidPatternEmail(value string) bool {
	return !regexp.MustCompile(`^[0-9a-zA-Z_.+-]+@[0-9a-zA-Z-]+\.[0-9a-zA-Z-.]+$`).MatchString(value)
}

func isInvalidLengthEmail(value string) bool {
	return 254 < len(value)
}
