package vo

import (
	"fmt"
)

// TodoTitle represents value obejct user name
type TodoTitle struct {
	value string
}

func NewTodoTitle(value string) (*TodoTitle, error) {
	if isInvalidLengthTodoTitle(value) {
		return nil, fmt.Errorf("todo title must be no longer than 50 characters")
	}
	return &TodoTitle{value: value}, nil
}

func (t TodoTitle) Value() string {
	return t.value
}

func isInvalidLengthTodoTitle(value string) bool {
	return 50 < len(value)
}
