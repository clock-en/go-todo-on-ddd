package vo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"unicode"
)

// Password represents value obejct password
type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	if isInvalidPatternPassword(value) {
		return nil, fmt.Errorf("password must be alphanumeric and symbols (-^$*.@)")
	}
	if isNotIncludeRequiredStringInPassword(value) {
		return nil, fmt.Errorf("password must contain at least one number, one lowercase letter, one uppercase letter, and one symbol (-^$*.@)")
	}
	if isInvalidLengthPassword(value) {
		return nil, fmt.Errorf("password must be at least 8 and no more than 100 characters")
	}
	return &Password{value: value}, nil
}

func (p Password) Value() string {
	return p.value
}

func (p Password) Hash() string {
	b := sha256.Sum256([]byte(p.value))
	return hex.EncodeToString(b[:])
}

func isNotIncludeRequiredStringInPassword(value string) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	for _, s := range value {
		switch {
		case unicode.IsUpper(s):
			hasUpper = true
		case unicode.IsLower(s):
			hasLower = true
		case unicode.IsNumber(s):
			hasNumber = true
		case unicode.IsSymbol(s):
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func isInvalidPatternPassword(value string) bool {
	return !regexp.MustCompile(`^[0-9a-zA-Z\-^$*.@]+$`).MatchString(value)
}

func isInvalidLengthPassword(value string) bool {
	return 254 < len(value)
}
