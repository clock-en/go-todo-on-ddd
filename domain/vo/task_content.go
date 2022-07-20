package vo

import (
	"fmt"
)

type ITaskContent interface {
	Value() string
}

type TaskContent struct {
	ITaskContent
	value string
}

func NewTaskContent(value string) (*TaskContent, error) {
	if isInvalidLengthTaskContent(value) {
		return nil, fmt.Errorf("todo task must be no longer than 140 characters")
	}
	return &TaskContent{value: value}, nil
}

func (t TaskContent) Value() string {
	return t.value
}

func isInvalidLengthTaskContent(value string) bool {
	return 140 < len(value)
}
