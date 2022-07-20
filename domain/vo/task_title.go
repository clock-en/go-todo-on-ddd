package vo

import (
	"fmt"
)

type ITaskTitle interface {
	Value()
}

type TaskTitle struct {
	ITaskTitle
	value string
}

func NewTaskTitle(value string) (*TaskTitle, error) {
	if isInvalidLengthTaskTitle(value) {
		return nil, fmt.Errorf("todo title must be no longer than 50 characters")
	}
	return &TaskTitle{value: value}, nil
}

func (t TaskTitle) Value() string {
	return t.value
}

func isInvalidLengthTaskTitle(value string) bool {
	return 50 < len(value)
}
