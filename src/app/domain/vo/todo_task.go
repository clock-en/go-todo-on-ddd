package vo

import (
	"fmt"
)

// TodoTask represents value obejct user name
type TodoTask struct {
	value string
}

func NewTodoTask(value string) (*TodoTask, error) {
	if isInvalidLengthTodoTask(value) {
		return nil, fmt.Errorf("todo task must be no longer than 140 characters")
	}
	return &TodoTask{value: value}, nil
}

func (t TodoTask) Value() string {
	return t.value
}

func isInvalidLengthTodoTask(value string) bool {
	return 140 < len(value)
}
